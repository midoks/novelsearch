package crontab

import (
	"errors"
	"fmt"
	// "github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "log"
	"regexp"
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

//获取小说path页数据
func CronPathInfo(v *models.AppItem, url string) {

	//time.Sleep(time.Duration(1) * time.Second)
	// fmt.Println(url)

	if content, err := getHttpData(url); err == nil {

		var (
			name            = ""
			author          = ""
			status          = ""
			category        = ""
			desc            = ""
			path            = ""
			chapter_content = ""
			// list            = ""
			err = errors.New("new")
		)
		//var chapter_list = make([]map[string]interface{})

		name, err = RegNovelSigleInfo(content, v.NameRule)
		if err != nil {
			return
		}
		fmt.Println(name)

		author, err = RegNovelSigleInfo(content, v.AuthorRule)
		if err != nil {
			return
		}
		fmt.Println(author)

		desc, err = RegNovelSigleInfo(content, v.DescRule)
		if err != nil {
			return
		}
		fmt.Println(desc)

		status, err = RegNovelSigleInfo(content, v.StatusRule)
		if err != nil {
			return
		}
		fmt.Println(status)

		category, err = RegNovelSigleInfo(content, v.CategoryRule)
		if err != nil {
			return
		}
		fmt.Println(category)

		path, err = RegNovelSigleInfo(content, v.ChapterPathRule)
		if err != nil {
			return
		}
		fmt.Println(path)

		var list = ""
		if chapter_content, err = getHttpData(path); err == nil {
			//fmt.Println(chapter_content)
			chapter_list, err2 := RegNovelList(chapter_content, v.ChapterListRule)
			if err2 != nil {
				return
			}
			if tmp_list, err3 := json.Marshal(chapter_list); err3 == nil {
				list = string(tmp_list)
			}
		}

		data := new(models.AppNovel)
		data.FromId = v.Id
		data.Name = name
		data.Author = author
		data.Desc = desc
		data.List = list
		data.BookStatus = status
		data.UpdateTime = time.Now().Unix()
		data.CreateTime = time.Now().Unix()
		_, err = orm.NewOrm().Insert(data)
		if err != nil {
			fmt.Println(err)
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
			//fmt.Println(pathList, time.Now())
			for _, url := range pathList {
				CronPathInfo(v, url)
				break
			}
		}
	}

	return nil
}
