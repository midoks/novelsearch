package crontab

import (
	"errors"
	// "fmt"
	// "github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getHttpData(url string) (string, error) {
	req := httplib.Get(url)

	str, err := req.String()
	if err != nil {
		return "", errors.New("资源获取错误!")
	}

	return str, nil
}

//匹配路径
func RegPathInfo(content, reg string) []string {
	match_exp := regexp.MustCompile(reg)
	list := libs.RemoveDuplicatesAndEmpty(match_exp.FindAllString(content, -1))
	return list
}

//匹配当个信息
func RegNovelSigleInfo(content, reg string) (string, error) {

	match_exp := regexp.MustCompile(reg)
	name := match_exp.FindAllStringSubmatch(content, -1)
	if len(name) == 0 {
		return "", errors.New("没有匹配到!")
	}
	return strings.TrimSpace(name[0][1]), nil
}

//匹配当个信息
func RegNovelList(content, reg string) ([]map[string]interface{}, error) {

	match_exp := regexp.MustCompile(reg)
	name := match_exp.FindAllStringSubmatch(content, -1)

	list := make([]map[string]interface{}, len(name))

	if len(name) == 0 {
		return list, errors.New("没有匹配到!")
	}

	for k, v := range name {
		tmp := make(map[string]interface{})
		tmp["name"] = v[2]
		tmp["url"] = v[1]
		list[k] = tmp
		// fmt.Println(list[k], v[1], v[2])
	}
	return list, nil
}

//
func CronPathInfoAndName(v *models.AppItem, name string, url string) {

}

//获取小说path页数据
func CronPathInfo(v *models.AppItem, url string) {

	logs.Info("采集开始:fromid:%d,%s", v.Id, url)
	if content, err := getHttpData(url); err == nil {

		var (
			name            = ""
			author          = ""
			status          = ""
			category        = ""
			desc            = ""
			path            = ""
			chapter_content = ""
			err             = errors.New("new")
		)

		name, err = RegNovelSigleInfo(content, v.NameRule)
		if err != nil {
			return
		}
		logs.Info("小说名:%s", name)
		vId := strconv.Itoa(v.Id)
		_, err = models.NovelGetByNameAndFromId(name, vId)
		if err == nil {
			logs.Info("已经采集了")
			return
		}

		author, err = RegNovelSigleInfo(content, v.AuthorRule)
		if err != nil {
			return
		}
		logs.Info("作者:%s", author)

		desc, err = RegNovelSigleInfo(content, v.DescRule)
		if err != nil {
			return
		}
		logs.Info("描述:%s", desc)

		status, err = RegNovelSigleInfo(content, v.StatusRule)
		if err != nil {
			return
		}
		logs.Info("状态:%s", status)

		category, err = RegNovelSigleInfo(content, v.CategoryRule)
		if err != nil {
			return
		}
		logs.Info("分类:%s", category)

		path, err = RegNovelSigleInfo(content, v.ChapterPathRule)
		if err != nil {
			return
		}
		logs.Info("小说目录页:%s", path)

		var list = ""
		var last_chapter = ""
		var last_chapter_url = ""
		var chapter_num = 0
		if chapter_content, err = getHttpData(path); err == nil {
			//fmt.Println(chapter_content)
			chapter_list, chapter_list_err := RegNovelList(chapter_content, v.ChapterListRule)
			if chapter_list_err != nil {
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
		}

		data := new(models.AppNovel)
		data.FromId = v.Id
		data.Name = name
		data.Author = author
		data.Desc = desc
		data.List = list
		data.ChapterNum = chapter_num
		data.LastChapter = last_chapter
		data.LastChapterUrl = last_chapter_url
		data.BookStatus = status
		data.UpdateTime = time.Now().Unix()
		data.CreateTime = time.Now().Unix()
		_, err = orm.NewOrm().Insert(data)
		if err == nil {
			logs.Info("采集结束:%s", url)
		} else {
			logs.Warn("采集发生错误:%s", err)
		}
	}
}

//首页爬取数据
func PageIndexSpider() error {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", "1")
	list, _ := models.ItemGetList(1, 10, filters...)

	for _, v := range list {

		if content, err := getHttpData(v.PageIndex); err == nil {
			pathList := RegPathInfo(content, v.PathRule)
			// fmt.Println(pathList, time.Now())
			for _, url := range pathList {
				CronPathInfo(v, url)
				break
			}
		}
	}

	return nil
}
