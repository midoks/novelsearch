package crontab

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"strings"
)

//	 second minute hour day month week   command
//顺序：秒      分    时   日   月    周      命令
func Init() {
	fmt.Println("crontab init")

	//测试使用
	// tk0 := toolbox.NewTask("test", "0 * * * * *", func() error { fmt.Println("begin--cron"); return nil })
	// toolbox.AddTask("test", tk0)

	//设置beego日志地址
	setLog()
	tkLog := toolbox.NewTask("tkLog", "59 * * * * *", setLog)
	toolbox.AddTask("tkLog", tkLog)

	checkFile := beego.AppConfig.String("cron.check_file")
	if !strings.EqualFold(checkFile, "") {
		checkFileFunc()
		tkCheckFile := toolbox.NewTask("tkCheckFile", checkFile, checkFileFunc)
		toolbox.AddTask("tkCheckFile", tkCheckFile)
	}

	//首页更新
	// PageIndexSpider()
	indexSpider := beego.AppConfig.String("cron.index_spider")
	if !strings.EqualFold(indexSpider, "") {
		tk1 := toolbox.NewTask("indexSpider", indexSpider, PageIndexSpider)
		toolbox.AddTask("indexSpider", tk1)
	}

	//连载更新
	// NovelIndexSpider()
	novelIndexSpider := beego.AppConfig.String("cron.novel_index_spider")
	if !strings.EqualFold(novelIndexSpider, "") {
		tk2 := toolbox.NewTask("novelIndexSpider", novelIndexSpider, NovelIndexSpider)
		toolbox.AddTask("novelIndexSpider", tk2)
	}

	//全站更新
	// WebRuleSpider()
	allSpider := beego.AppConfig.String("cron.all_spider")
	if !strings.EqualFold(allSpider, "") {
		tk3 := toolbox.NewTask("allSpider", allSpider, WebRuleSpider)
		toolbox.AddTask("allSpider", tk3)
	}
	//搜索更新
	// SosoSpider()
	sosoSpider := beego.AppConfig.String("cron.soso_spider")
	if !strings.EqualFold(sosoSpider, "") {
		tk4 := toolbox.NewTask("SosoSpider", sosoSpider, SosoSpider)
		toolbox.AddTask("SosoSpider", tk4)
	}

	//百度榜单
	baidubangSpider := beego.AppConfig.String("cron.baidubang_spider")
	if !strings.EqualFold(baidubangSpider, "") {
		BaiduTopAll()
		tk5 := toolbox.NewTask("baidubangSpider", baidubangSpider, BaiduTopAll)
		toolbox.AddTask("baidubangSpider", tk5)
	}

	//邮件发送
	// CronSendMail()
	sendMail := beego.AppConfig.String("cron.sendmail")
	if !strings.EqualFold(sendMail, "") {
		tk6 := toolbox.NewTask("sendMail", sendMail, CronSendMail)
		toolbox.AddTask("sendMail", tk6)
	}
}
