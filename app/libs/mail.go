package libs

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/models"
	"gopkg.in/gomail.v2"
	"strconv"
	"strings"
	"time"
)

func SendMail(tomail string, subject string, conent string) {

	mailHost := models.OptionGet(models.MAIL_HOST, "")
	mailPort := models.OptionGet(models.MAIL_PORT, "25")
	mailPortInt, _ := strconv.Atoi(mailPort)
	mailUser := models.OptionGet(models.MAIL_USER, "")
	mailPwd := models.OptionGet(models.MAIL_PWD, "")

	if strings.EqualFold(mailHost, "") ||
		strings.EqualFold(mailUser, "") ||
		strings.EqualFold(mailPwd, "") {

		fmt.Println("邮件未设置")
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mailUser)
	m.SetHeader("To", tomail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", conent)

	d := gomail.Dialer{Host: mailHost, Port: mailPortInt, Username: mailUser, Password: mailPwd}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(beego.Date(time.Now(), "Y-m-d H:i:s"), err)
		return
	}
	return
}
