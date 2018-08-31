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

	//邮件发送
	// CronSendMail()
	sendMail := beego.AppConfig.String("cron.sendmail")
	if !strings.EqualFold(sendMail, "") {
		tk6 := toolbox.NewTask("sendMail", sendMail, CronSendMail)
		toolbox.AddTask("sendMail", tk6)
	}
}
