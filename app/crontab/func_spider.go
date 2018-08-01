package crontab

import (
	"errors"
	// "github.com/astaxie/beego"
	// "encoding/json"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/models"
	// "regexp"
	"strconv"
	"strings"
	"time"
)

func CronWebRuleSpider(v *models.AppItem, url string, ranges string, rule string, path_tpl string) {
	timeStart := time.Now().Unix()
	var (
		end      = 1
		err      = errors.New("nil")
		cur      = "0"
		cur_page = ""
	)

	list := strings.Split(ranges, ",")

	end, err = strconv.Atoi(list[1])
	if err != nil {
		return
	}

	v.SpiderProgress = v.SpiderProgress + 1
	logs.Info("网站(%s)采集:进度:%d, 结束在:%d", v.Name, v.SpiderProgress, end)

	cur = strconv.Itoa(v.SpiderProgress)
	cur_page = strings.Replace(url, "{$RANGE}", cur, -1)

	logs.Warn("全站采集开始:url:%s", cur_page)

	if content, errcur := getHttpData2Code(cur_page, v.PageCharset); errcur == nil {
		list, errlist := RegNovelList(content, rule)

		if errlist == nil {
			if len(list) > 0 {
				for j := 0; j < len(list); j++ {
					url := strings.Replace(path_tpl, "{$ID}", list[j]["url"].(string), -1)
					if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
						url = "http://" + url
					}
					go CronPathInfo(v, url, list[j]["name"].(string))
				}
			} else {
				v.SpiderProgress = 0
				v.Update("SpiderProgress")
				logs.Error("全站采集结束(重置)url:%s", cur_page)
				return
			}
		} else {
			logs.Error("全站采集错误:%s", cur_page, errlist, rule)
		}
	}

	timeEnd := time.Now().Unix()
	v.Update("SpiderProgress")
	logs.Warn("全站采集结束:url:%s耗时:%d", cur_page, timeEnd-timeStart)
}

//首页爬取数据
func WebRuleSpider() error {

	filters := make([]interface{}, 0)
	filters = append(filters, "status", "1")
	list, _ := models.ItemGetList(1, 10000, filters...)

	if len(list) == 0 {
		logs.Info("全站更新(无更新数据):end!")
		return nil
	}

	for i := 0; i < len(list); i++ {
		var r = list[i]
		if r.SpiderExp != "" && r.SpiderRange != "" && r.SpiderRule != "" && r.PathTpl != "" {
			go CronWebRuleSpider(r, r.SpiderExp, r.SpiderRange, r.SpiderRule, r.PathTpl)
		} else {
			logs.Info("全站更新(条件不足):end!")
		}
	}
	return nil
}
