package crontab

import (
	// "errors"
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	// "github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"strings"
)

func CronSosoSpider(v *models.AppItem, url string,
	page_arg string,
	rule string,
	path_tpl string,
	keyword string) {

	var page_url = strings.Replace(url, "{$KEYWORD}", keyword, -1)
	var page_num = 1

	for {

		var isEmpty = true
		cur_page_url := fmt.Sprintf("%s&%s=%d", page_url, page_arg, page_num)
		logs.Info("搜索爬取开始:url:%s", cur_page_url)

		if content, errcur := getHttpData2Code(cur_page_url, v.PageCharset); errcur == nil {

			list, errlist := RegNovelList(content, rule)
			if errlist == nil {

				if len(list) > 0 {
					for j := 0; j < len(list); j++ {
						url := strings.Replace(path_tpl, "{$ID}", list[j]["url"].(string), -1)
						name := list[j]["name"].(string)
						CronPathInfo(v, url, name)
					}
					isEmpty = false
				} else {
					logs.Info("搜索爬取(没有数据):url:%s", cur_page_url)
					break
				}
			}
		}

		logs.Info("搜索爬取结束:url:%s", cur_page_url)
		if isEmpty {
			break
		}
		page_num++
	}
}

//搜索内容派爬取
func SosoSpider() error {

	logs.Info("搜索内容派爬取---start!")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", "1")
	list, _ := models.ItemGetList(1, 10, filters...)

	if len(list) == 0 {
		logs.Info("搜索内容派爬取(无更新数据)---end!")
		return nil
	}

	for i := 0; i < len(list); i++ {
		var r = list[i]
		if r.SosoExp != "" && r.SosoRule != "" &&
			r.SosoPageArgs != "" &&
			r.PathTpl != "" {
			CronSosoSpider(r, r.SosoExp, r.SosoPageArgs, r.SosoRule, r.PathTpl, "凡人")
		} else {
			logs.Info("搜索内容派爬取(条件不足)---end!")
		}
	}

	logs.Info("搜索内容派爬取---end!")
	return nil
}
