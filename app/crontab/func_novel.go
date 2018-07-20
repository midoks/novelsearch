package crontab

import (
	"errors"
	// "fmt"
	// "github.com/astaxie/beego"
	"encoding/json"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	// "github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "regexp"
	"strings"
	"time"
)

//获取小说path页数据
func CronNovelUpdate(v *models.AppItem, n *models.AppNovel, url, name string) {

	logs.Info("更新开始:fromid:%d,novel:%d,%s:%s", v.Id, n.Id, name, url)

	if content, err := getHttpData(url); err == nil {

		var (
			status          = ""
			path            = ""
			chapter_content = ""
			err             = errors.New("new")
			book_status     = 0
		)

		status, err = RegNovelSigleInfo(content, v.StatusRule)
		if err != nil {
			return
		}
		logs.Info("状态:%s", status)

		//判断是否已经结束
		if strings.EqualFold(status, v.StatusEndMark) {
			book_status = 1
		}

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

		if strings.EqualFold(n.List, list) {
			logs.Info("更新内容相同:fromid:%d,novel:%d,%s:%s", v.Id, n.Id, name, url)
		}
		n.ChapterNum = chapter_num
		n.LastChapter = last_chapter
		n.LastChapterUrl = last_chapter_url
		n.BookStatus = book_status
		n.List = list
		n.UpdateTime = time.Now().Unix()
		err = n.Update()
		if err == nil {
			logs.Info("更新结束:fromid:%d,novel:%d,%s:%s", v.Id, n.Id, name, url)
		} else {
			logs.Info("更新错误:fromid:%d,novel:%d,%s:%s", v.Id, n.Id, name, url)
		}
	}
}

//首页爬取数据
func NovelIndexSpider() error {

	logs.Info("小说开始更新---start!")

	list := models.CronNovelGetByStatus("0", 6, 1)
	if len(list) == 0 {
		logs.Info("小说开始更新(无更新数据)---end!")
		return nil
	}

	for i := 0; i < len(list); i++ {

		item, err := models.ItemGetById(list[i].FromId)
		if err == nil {
			CronNovelUpdate(item, list[i], list[i].Url, list[i].Name)
		}
	}

	logs.Info("小说开始更新---end!")
	return nil
}
