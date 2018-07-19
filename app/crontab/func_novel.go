package crontab

import (
	// "errors"
	"fmt"
	// "github.com/astaxie/beego"
	// "encoding/json"
	// "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	// "github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	// "regexp"
	// "strconv"
	// "strings"
	// "time"
)

//首页爬取数据
func NovelIndexSpider() error {
	logs.Info("小说开始更新---start!")

	list := models.CronNovelGetByStatus("连载中", 100)
	fmt.Println(list)

	logs.Info("小说开始更新---end!")
	return nil
}
