package crontab

import (
	"encoding/json"
	"errors"
	// "fmt"
	"github.com/astaxie/beego/logs"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"html"
	"strconv"
	"strings"
	"time"
)

type BookLinkInfo struct {
	Name string
	Url  string
}

//小说内容更新
func CronNovelContentList(v *models.AppItem, n *models.AppNovel) {
	var bli []BookLinkInfo
	err := json.Unmarshal([]byte(n.List), &bli)

	key := "lock_" + strconv.Itoa(n.Id)
	lock_val := "1"
	//加锁
	lockData, _ := libs.GetCache(key)
	if strings.EqualFold(lockData.(string), lock_val) {
		// logs.Warn("小说内容更新,正在执行中...!")
		return
	}

	if err == nil {
		for _, data := range bli {
			var url = data.Url
			go CronNovelContent(url, v.ContentRule, v.PageCharset)
		}
	}
	libs.SetCache(key, lock_val, 60*60*time.Second)
}

//小说内容更新
func CronNovelContent(url, rule string, charset string) (string, error) {
	key := "novel_" + url
	if libs.CacheIsExist(key) {
		content, err := libs.GetCache(key)
		return content.(string), err
	}
	logs.Info("请求缓存(%s)不存在,向目标地址发送请求", url)

	urlData, err := libs.GetHttpData(url)
	if err != nil {
		return "", err
	}

	content, err := RegNovelSigleInfo(urlData, rule)
	if err != nil {
		return "", err
	}

	if strings.EqualFold(charset, "gbk") {
		content = libs.ConvertToString(content, "gbk", "utf8")
	}

	content = html.UnescapeString(content)

	libs.SetCache(key, content, 3*24*60*60*time.Second)
	return content, nil
}

//获取小说path页数据
func CronNovelUpdate(v *models.AppItem, n *models.AppNovel, url, name string) {

	logs.Info("小说(%s)自动更新开始:fromid:%d,novel:%d,%s", name, v.Id, n.Id, url)

	if content, err := getHttpData2Code(url, v.PageCharset); err == nil {

		var (
			status          = ""
			path            = ""
			desc            = ""
			chapter_content = ""
			err             = errors.New("new")
			book_status     = 0
		)

		status, err = RegNovelSigleInfo(content, v.StatusRule)
		if err != nil {
			logs.Error("状态获取(失败)url(%s):%s", url, err)
			return
		}
		logs.Info("状态:%s", status)

		desc, err = RegNovelSigleInfo(content, v.DescRule)
		if err != nil {
			logs.Error("描述获取(失败[%s]):%s", url, err)
			return
		}
		desc = libs.TrimHtml(desc)
		logs.Info("描述:%s", desc)

		//判断是否已经结束
		if strings.EqualFold(status, v.StatusEndMark) {
			book_status = 1
		}

		var list = ""
		var last_chapter = ""
		var last_chapter_url = ""
		var chapter_num = 0
		var ab_path = ""

		//如果获取章节目录页规则为空，则认为目录首页里,直接匹配
		if strings.EqualFold(v.ChapterPathRule, "") {
			chapter_content = content
			ab_path = url
		} else {
			path, err = RegNovelSigleInfo(content, v.ChapterPathRule)
			if err != nil {
				logs.Error("目录获取(失败):%s", err)
				return
			}

			ab_path = GetAbsoluteAddr(url, path)
			logs.Info("绝对路径:", ab_path)

			chapter_content, err = getHttpData2Code(ab_path, v.PageCharset)
			if err != nil {
				logs.Error("资源获取失败:%s", err, ab_path)
				return
			}
		}

		chapter_list, chapter_list_err := RegNovelListAutoPath(chapter_content, v.ChapterListRule, ab_path)
		if chapter_list_err != nil {
			logs.Error("小说列表获取错误:%s", err, ab_path)
			return
		}

		chapter_num = len(chapter_list)
		if chapter_num > 0 {
			last_chapter = chapter_list[chapter_num-1]["name"].(string)
			last_chapter_url = chapter_list[chapter_num-1]["url"].(string)
		}

		tmp_list, tmp_list_err := json.Marshal(chapter_list)
		if tmp_list_err == nil {
			list = string(tmp_list)
		}

		if strings.EqualFold(n.List, list) {
			logs.Warn("更新内容相同(end):fromid:%d,novel:%d,%s:%s", v.Id, n.Id, name, url)
			n.Update("UpdateTime")
			return
		}
		n.ChapterNum = chapter_num
		n.LastChapter = last_chapter
		n.LastChapterUrl = last_chapter_url
		n.BookStatus = book_status
		n.Desc = desc
		n.List = list
		err = n.Update("Desc", "ChapterNum", "LastChapter", "LastChapterUrl", "BookStatus", "UpdateTime")
		if err == nil {
			logs.Info("更新结束:fromid:%d,novel:%d,%s:%s", v.Id, n.Id, name, url)
		} else {
			logs.Info("更新错误:fromid:%d,novel:%d,%s:%s", v.Id, n.Id, name, url)
		}
	}
}

//爬取数据
func NovelIndexSpider() error {

	logs.Info("小说开始更新---start!")

	list := models.CronNovelGetByStatus("0", 60*60*3, 10)
	if len(list) == 0 {
		logs.Info("小说开始更新(无更新数据)---end!")
		return nil
	}

	for i := 0; i < len(list); i++ {

		item, err := models.ItemGetById(list[i].FromId)
		if err == nil {
			go CronNovelUpdate(item, list[i], list[i].Url, list[i].Name)
		}
	}

	logs.Info("小说开始更新---end!")
	return nil
}
