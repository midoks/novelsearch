package crontab

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

//	 minute hour day month week   command
//顺序：分    时   日   月    周
func Init() {
	fmt.Println("crontab init")

	tk0 := toolbox.NewTask("test", "0 * * * * *", func() error { fmt.Println("begin--cron"); return nil })
	toolbox.AddTask("test", tk0)

	//首页更新
	// PageIndexSpider() //test
	indexSpider := beego.AppConfig.String("cron.index_spider")
	tk1 := toolbox.NewTask("indexSpider", indexSpider, PageIndexSpider)
	toolbox.AddTask("indexSpider", tk1)

	//连载更新
	// NovelIndexSpider()
	novelIndexSpider := beego.AppConfig.String("cron.novel_index_spider")
	tk2 := toolbox.NewTask("novelIndexSpider", novelIndexSpider, NovelIndexSpider)
	toolbox.AddTask("novelIndexSpider", tk2)

	//全站更新
	// WebRuleSpider()
	allSpider := beego.AppConfig.String("cron.all_spider")
	tk3 := toolbox.NewTask("allSpider", allSpider, WebRuleSpider)
	toolbox.AddTask("allSpider", tk3)

	//搜索更新
	// SosoSpider()
	sosoSpider := beego.AppConfig.String("cron.soso_spider")
	tk4 := toolbox.NewTask("SosoSpider", sosoSpider, SosoSpider)
	toolbox.AddTask("SosoSpider", tk4)

	//百度榜单
	BaiduTopAll()
	baidubangSpider := beego.AppConfig.String("cron.baidubang_spider")
	tk5 := toolbox.NewTask("baidubangSpider", baidubangSpider, BaiduTopAll)
	toolbox.AddTask("baidubangSpider", tk5)

	//邮件发送
	// CronSendMail()
	sendMail := beego.AppConfig.String("cron.sendmail")
	tk6 := toolbox.NewTask("sendMail", sendMail, CronSendMail)
	toolbox.AddTask("sendMail", tk6)
}
