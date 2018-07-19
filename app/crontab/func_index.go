package crontab

import (
	// "errors"
	// "github.com/astaxie/beego"
	// "encoding/json"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "regexp"
	// "strconv"
	"strings"
	// "time"
)

//首页爬取数据
func PageIndexSpider() error {
	logs.Info("首页开始采集---start!")
	filters := make([]interface{}, 0)
	filters = append(filters, "status", "1")
	list, _ := models.ItemGetList(1, 10, filters...)

	for _, v := range list {

		if content, err := getHttpData(v.PageIndex); err == nil {

			var tmpRule = strings.TrimSpace(v.PageIndexRule)
			tmpRule = strings.Replace(tmpRule, "\n", "|", -1)
			tmpRule = strings.Replace(tmpRule, "\r\n", "|", -1)
			tmpRuleList := strings.Split(tmpRule, "|")

			for i := 0; i < len(tmpRuleList); i++ {
				logs.Info(tmpRuleList[i])

				pathList, err := RegPathInfo(content, tmpRuleList[i])
				if err == nil {
					// logs.Info(pathList)
					for _, val := range pathList {
						if libs.IsUrlRe(val[1]) {
							CronPathInfo(v, val[1], val[2])
						} else {
							CronPathInfo(v, val[2], val[1])
						}
					}
				}
			}
		}
	}
	logs.Info("首页开始采集---end!")
	return nil
}
