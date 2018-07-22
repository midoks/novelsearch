package crontab

import (
	"errors"
	"fmt"
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
	"strconv"
)

func CronWebRuleSpider(v *models.AppItem, url string, ranges string, rule string) {

	var (
		start    = 1
		end      = 1
		err      = errors.New("nil")
		cur      = "0"
		cur_page = ""
	)

	list := strings.Split(ranges, ",")

	start, err = strconv.Atoi(list[0])
	if err != nil {
		return
	}

	end, err = strconv.Atoi(list[1])
	if err != nil {
		return
	}

	for i := start; i < end; i++ {
		cur = strconv.Itoa(i)
		cur_page = strings.Replace(url, "{$RANGE}", cur, -1)

		logs.Info("全站采集开始:url:%s", cur_page)
		fmt.Println(cur_page)

		var isEmpty = false

		if content, errcur := getHttpData(cur_page); errcur == nil {

			list, errlist := RegNovelList(content, rule)
			if errlist == nil {

				if len(list) > 0 {
					for j := 0; j < len(list); j++ {
						// fmt.Println(list[j])
						go CronPathInfo(v, list[j]["url"].(string), list[j]["name"].(string))
					}
				} else {
					logs.Info("全站采集结束(没有数据):url:%s", cur_page)
					isEmpty = true
					break
				}
			}
		}

		if !isEmpty {
			break
		}
		logs.Info("全站采集结束:url:%s", cur_page)
	}
}

//首页爬取数据
func WebRuleSpider() error {

	logs.Info("全站更新---start!")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", "1")
	list, _ := models.ItemGetList(1, 10, filters...)

	if len(list) == 0 {
		logs.Info("全站更新(无更新数据)---end!")
		return nil
	}

	for i := 0; i < len(list); i++ {
		var r = list[i]
		if r.SpiderExp != "" && r.SpiderRange != "" && r.SpiderRule != "" {
			CronWebRuleSpider(r, r.SpiderExp, r.SpiderRange, r.SpiderRule)
		} else {
			logs.Info("全站更新(条件不足)---end!")
		}
	}

	logs.Info("全站更新---end!")
	return nil
}
