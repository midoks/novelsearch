package backends

import (
	"github.com/midoks/novelsearch/app/models"
)

type SysSettingController struct {
	BaseController
}

const (
	WEB_NAME        = "web_name"
	WEB_KEYWORD     = "web_keyword"
	WEB_DESC        = "web_desc"
	WEB_STAT        = "web_stat"
	WEB_NOTICE      = "web_notice"
	WEB_NOTICE_MAIL = "web_notice_mail"

	MAIL_HOST       = "mail_host"
	MAIL_PORT       = "mail_port"
	MAIL_USER       = "mail_user"
	MAIL_PWD        = "mail_pwd"
	MAIL_QUEUE_SIZE = "mail_queue_size"
)

func (this *SysSettingController) Index() {

	if this.isPost() {
		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		models.OptionSet(WEB_NAME, vars[WEB_NAME])
		models.OptionSet(WEB_KEYWORD, vars[WEB_KEYWORD])
		models.OptionSet(WEB_DESC, vars[WEB_DESC])
		models.OptionSet(WEB_STAT, vars[WEB_STAT])
		models.OptionSet(WEB_NOTICE, vars[WEB_NOTICE])
		models.OptionSet(WEB_NOTICE_MAIL, vars[WEB_NOTICE_MAIL])
	}

	webName := models.OptionGet(WEB_NAME, "小说搜索")
	this.Data[WEB_NAME] = webName

	webKeyword := models.OptionGet(WEB_KEYWORD, "小说关键字")
	this.Data[WEB_KEYWORD] = webKeyword

	webDesc := models.OptionGet(WEB_DESC, "小说描述")
	this.Data[WEB_DESC] = webDesc

	webStat := models.OptionGet(WEB_STAT, "小说统计代码")
	this.Data[WEB_STAT] = webStat

	webNotice := models.OptionGet(WEB_NOTICE, "网站通知")
	this.Data[WEB_NOTICE] = webNotice

	webNoticeMail := models.OptionGet(WEB_NOTICE_MAIL, "")
	this.Data[WEB_NOTICE_MAIL] = webNoticeMail

	this.display()
}

func (this *SysSettingController) Mail() {

	if this.isPost() {
		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		models.OptionSet(MAIL_HOST, vars[MAIL_HOST])
		models.OptionSet(MAIL_PORT, vars[MAIL_PORT])
		models.OptionSet(MAIL_USER, vars[MAIL_USER])
		models.OptionSet(MAIL_PWD, vars[MAIL_PWD])
		models.OptionSet(WEB_NOTICE, vars[WEB_NOTICE])
	}

	mailHost := models.OptionGet(MAIL_HOST, "smtp.163.com")
	this.Data[MAIL_HOST] = mailHost

	mailPort := models.OptionGet(MAIL_PORT, "25")
	this.Data[MAIL_PORT] = mailPort

	mailUser := models.OptionGet(MAIL_USER, "")
	this.Data[MAIL_USER] = mailUser

	mailPwd := models.OptionGet(MAIL_PWD, "")
	this.Data[MAIL_PWD] = mailPwd

	mailQueueSize := models.OptionGet(MAIL_QUEUE_SIZE, "100")
	this.Data[MAIL_QUEUE_SIZE] = mailQueueSize

	this.display()
}
