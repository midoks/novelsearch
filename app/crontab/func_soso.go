package crontab

import (
	// "errors"
	// "fmt"
	// "github.com/astaxie/beego"
	// "encoding/json"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	// "github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "regexp"
	"strings"
	// "time"
	// "strconv"
)

func CronSosoSpider(v *models.AppItem, url string, page string, rule string, keyword string) {
	cur_page_url := strings.Replace(url, "{$KEYWORD}", keyword, -1)
	logs.Info("搜索爬取开始:url:%s", cur_page_url)
	for {
		var isEmpty = true

		if content, errcur := getHttpData(cur_page_url); errcur == nil {

			list, errlist := RegNovelList(content, rule)
			if errlist == nil {

				if len(list) > 0 {
					for j := 0; j < len(list); j++ {
						// CronPathInfo(v, list[j]["url"].(string), list[j]["name"].(string))
					}
				} else {
					logs.Info("搜索爬取(没有数据):url:%s", cur_page_url)
					// isEmpty = false
					break
				}
			}
		}

		if isEmpty {
			break
		}
	}
	logs.Info("搜索爬取结束:url:%s", cur_page_url)
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
		if r.SosoExp != "" && r.SosoRule != "" && r.SosoPageArgs != "" {
			CronSosoSpider(r, r.SosoExp, r.SosoPageArgs, r.SosoRule, "凡人")
		} else {
			logs.Info("搜索内容派爬取(条件不足)---end!")
		}
	}

	logs.Info("搜索内容派爬取---end!")
	return nil
}
