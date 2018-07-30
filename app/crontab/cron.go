package crontab

import (
	"fmt"
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
	// tk1 := toolbox.NewTask("PageIndexSpider", "0/10 * * * * *", PageIndexSpider)
	// toolbox.AddTask("PageIndexSpider", tk1)

	//连载更新
	// NovelIndexSpider()
	// tk2 := toolbox.NewTask("NovelIndexSpider", "0/30 * * * * *", NovelIndexSpider)
	// toolbox.AddTask("NovelIndexSpider", tk2)

	//全站更新
	// WebRuleSpider()
	// tk3 := toolbox.NewTask("WebRuleSpider", "0 */2 * * * *", WebRuleSpider)
	// toolbox.AddTask("WebRuleSpider", tk3)

	//搜索更新
	// SosoSpider()
	// tk4 := toolbox.NewTask("SosoSpider", "0 */2 * * * *", SosoSpider)
	// toolbox.AddTask("SosoSpider", tk4)

	//百度榜单
	BaiduTopAll()
	tk5 := toolbox.NewTask("BaiduBang", "59 23 * * *", BaiduTopAll)
	toolbox.AddTask("BaiduBang", tk5)

	//邮件发送
	// CronSendMail()
	tk6 := toolbox.NewTask("CronSendMail", "59 59 23 * * *", CronSendMail)
	toolbox.AddTask("CronSendMail", tk6)
}
