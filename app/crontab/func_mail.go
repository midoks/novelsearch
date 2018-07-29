package crontab

import (
	"bytes"
	"fmt"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"html/template"
	"strings"
)

type DataInfo struct {
	AllCollection int64
	DayCollection int64
}

func getMailInfo() string {
	const MAIL_TEMPLATE = `
<div>
    <table style="border-collapse:collapse;border: 1px solid gray;">
        <tr style="border-collapse:collapse;text-align: center;">
            <th style="border-collapse:collapse;border: 1px solid gray;">总收录</th>
            <th style="border-collapse:collapse;border: 1px solid gray;">日收录</th>
        </tr>

        <tr style="border-collapse:collapse;text-align: center;">
            <td style="border-collapse:collapse;border: 1px solid gray;">{{.AllCollection}}</td>
            <td style="border-collapse:collapse;border: 1px solid gray;">{{.DayCollection}}</td>
        </tr>
    </table>
</div>
<hr />
<div>欢迎使用novelsearch | 源码地址:<a href="https://github.com/midoks/novelsearch">https://github.com/midoks/novelsearch</a></div>
`

	data := DataInfo{}
	data.AllCollection = models.NovelCount("")
	data.DayCollection = models.NovelTodayCount("")

	buffer := new(bytes.Buffer)
	t, _ := template.New("day_template").Parse(MAIL_TEMPLATE)
	t.Execute(buffer, data)

	return buffer.String()
}

func CronSendMail() error {
	info := getMailInfo()

	webNoticeMail := models.OptionGet(models.WEB_NOTICE_MAIL, "")
	if strings.EqualFold(webNoticeMail, "") {
		fmt.Println("邮件通知地址未设置")
		return nil
	}
	webName := models.OptionGet(models.WEB_NAME, "小说搜索")
	// fmt.Println(info)
	libs.SendMail(webNoticeMail, webName+":每日汇总信息", info)
	return nil
}
