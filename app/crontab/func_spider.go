package crontab

import (
	"errors"
	// "github.com/astaxie/beego"
	// "encoding/json"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "regexp"
	"strconv"
	"strings"
	"time"
)

func CronWebRuleSpider(v *models.AppItem, url string, ranges string, rule string, path_tpl string) {

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
		logs.Error("start:%s", list[0])
		return
	}

	end, err = strconv.Atoi(list[1])
	if err != nil {
		return
	}

	for i := start; i < end; i++ {
		cur = strconv.Itoa(i)
		cur_page = strings.Replace(url, "{$RANGE}", cur, -1)

		logs.Warn("全站采集开始:url:%s", cur_page)

		var isEmpty = true
		if content, errcur := getHttpData(cur_page); errcur == nil {

			if strings.EqualFold(v.PageCharset, "gbk") {
				content = libs.ConvertToString(content, "gbk", "utf8")
			}

			list, errlist := RegNovelList(content, rule)

			if errlist == nil {
				if len(list) > 0 {
					for j := 0; j < len(list); j++ {
						url := strings.Replace(path_tpl, "{$ID}", list[j]["url"].(string), -1)
						if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
							url = "http://" + url
						}
						// logs.Warn("采集页:url:%s", url)
						CronPathInfo(v, url, list[j]["name"].(string))
					}
					isEmpty = false
				} else {
					logs.Error("全站采集结束(没有数据)url:%s", cur_page)
					break
				}
			} else {
				logs.Error("全站采集错误:%s", errlist, rule, content)
			}
		}
		if isEmpty {
			break
		}

		logs.Warn("全站采集结束:url:%s", cur_page)
	}
}

//首页爬取数据
func WebRuleSpider() error {
	timeStart := time.Now().Unix()
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
		if r.SpiderExp != "" && r.SpiderRange != "" && r.SpiderRule != "" && r.PathTpl != "" {
			CronWebRuleSpider(r, r.SpiderExp, r.SpiderRange, r.SpiderRule, r.PathTpl)
		} else {
			logs.Info("全站更新(条件不足)---end!")
		}
	}
	timeEnd := time.Now().Unix()
	logs.Info("全站更新---end!", timeEnd-timeStart)
	return nil
}
