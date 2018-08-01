package crontab

import (
	"errors"
	_ "fmt"
	// "github.com/astaxie/beego"
	"encoding/json"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "io/ioutil"
	"strconv"
	"strings"
	"time"
)

//获取小说path页数据
func CronPathInfo(v *models.AppItem, url, name string) {

	logs.Warn("目录页采集开始:fromid:%d,%s:%s", v.Id, name, url)
	vId := strconv.Itoa(v.Id)
	_, errFind := models.CronNovelGetByNameAndFromId(name, vId)
	if errFind == nil {
		logs.Info("已经采集了")
		return
	}

	// s_name := name
	if content, err := getHttpData2Code(url, v.PageCharset); err == nil {

		var (
			name            = ""
			author          = ""
			status          = ""
			category        = ""
			desc            = ""
			path            = ""
			chapter_content = ""
			err             = errors.New("new")
			book_status     = 0
		)
		//fmt.Println(content)
		// ioutil.WriteFile("t/"+s_name+".html", []byte(content), 0644)

		name, err = RegNovelSigleInfo(content, v.NameRule)
		if err != nil {
			logs.Error("小说名获取错误:", v.NameRule)
			return
		}
		logs.Info("小说名:%s", name)
		_, errFind := models.CronNovelGetByNameAndFromId(name, vId)
		if errFind == nil {
			logs.Error("已经采集了(名字不一致哟)")
			return
		}

		author, err = RegNovelSigleInfo(content, v.AuthorRule)
		if err != nil {
			logs.Error("作者获取(失败):%s", err)
			return
		}
		logs.Info("作者:%s", author)

		desc, err = RegNovelSigleInfo(content, v.DescRule)
		if err != nil {
			logs.Error("描述获取(失败[%s]):%s", name, err)
			return
		}
		logs.Info("描述:%s", desc)

		status, err = RegNovelSigleInfo(content, v.StatusRule)
		if err != nil {
			logs.Error("状态获取(失败):%s", err)
			return
		}
		logs.Info("状态:%s", status)

		//判断是否已经结束
		if strings.EqualFold(status, v.StatusEndMark) {
			book_status = 1
		}

		category, err = RegNovelSigleInfo(content, v.CategoryRule)
		if err != nil {
			logs.Error("分类获取(失败):%s", err)
			return
		}
		logs.Info("分类:%s", category)

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
			logs.Error("匹配列表失败:%s", chapter_list_err, v.ChapterListRule)
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

		data := new(models.AppNovel)
		data.FromId = v.Id
		data.Url = url
		data.List = list
		data.UniqueId = libs.Md5str(url)
		data.Name = name
		data.Category = category
		data.Author = author
		data.Desc = desc
		data.List = list
		data.ChapterNum = chapter_num
		data.LastChapter = last_chapter
		data.LastChapterUrl = last_chapter_url
		data.BookStatus = book_status
		data.UpdateTime = time.Now().Unix()
		data.CreateTime = time.Now().Unix()
		_, err = orm.NewOrm().Insert(data)
		if err == nil {
			logs.Warn("目录页采集结束:%s", url)
		} else {
			logs.Error("目录页采集发生错误:%s", err)
		}
	} else {
		logs.Warn("目录页采集结束(没有获取到资源):%s", url)
	}
}

//首页爬取数据
func PageIndexSpider() error {
	logs.Info("首页开始采集---start!")
	filters := make([]interface{}, 0)
	filters = append(filters, "status", "1")
	list, _ := models.ItemGetList(1, 10, filters...)

	if len(list) == 0 {
		logs.Info("首页开始采集(无更新数据)---end!")
		return nil
	}

	for _, v := range list {

		if content, err := getHttpData2Code(v.PageIndex, v.PageCharset); err == nil {

			var tmpRule = strings.TrimSpace(v.PageIndexRule)
			tmpRule = strings.Replace(tmpRule, "\n", "|", -1)
			tmpRule = strings.Replace(tmpRule, "\r\n", "|", -1)
			tmpRuleList := strings.Split(tmpRule, "|")

			for i := 0; i < len(tmpRuleList); i++ {

				pathList, err := RegPathInfo(content, tmpRuleList[i])
				if err == nil {
					for _, val := range pathList {
						url := strings.Replace(v.PathTpl, "{$ID}", val[1], -1)
						CronPathInfo(v, url, val[2])
					}
				}
			}
		}
	}
	logs.Info("首页开始采集---end!")
	return nil
}
