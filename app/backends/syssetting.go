package backends

import (
	"github.com/midoks/novelsearch/app/models"
)

type SysSettingController struct {
	BaseController
}

func (this *SysSettingController) Index() {

	if this.isPost() {
		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		models.OptionSet(models.WEB_NAME, vars[models.WEB_NAME])
		models.OptionSet(models.WEB_KEYWORD, vars[models.WEB_KEYWORD])
		models.OptionSet(models.WEB_DESC, vars[models.WEB_DESC])
		models.OptionSet(models.WEB_STAT, vars[models.WEB_STAT])
		models.OptionSet(models.WEB_NOTICE, vars[models.WEB_NOTICE])
		models.OptionSet(models.WEB_NOTICE_MAIL, vars[models.WEB_NOTICE_MAIL])
	}

	webName := models.OptionGet(models.WEB_NAME, "小说搜索")
	this.Data[models.WEB_NAME] = webName

	webKeyword := models.OptionGet(models.WEB_KEYWORD, "小说关键字")
	this.Data[models.WEB_KEYWORD] = webKeyword

	webDesc := models.OptionGet(models.WEB_DESC, "小说描述")
	this.Data[models.WEB_DESC] = webDesc

	webStat := models.OptionGet(models.WEB_STAT, "小说统计代码")
	this.Data[models.WEB_STAT] = webStat

	webNotice := models.OptionGet(models.WEB_NOTICE, "网站通知")
	this.Data[models.WEB_NOTICE] = webNotice

	webNoticeMail := models.OptionGet(models.WEB_NOTICE_MAIL, "")
	this.Data[models.WEB_NOTICE_MAIL] = webNoticeMail

	this.display()
}

func (this *SysSettingController) Mail() {

	if this.isPost() {
		vars := make(map[string]string)
		this.Ctx.Input.Bind(&vars, "vars")

		models.OptionSet(models.MAIL_HOST, vars[models.MAIL_HOST])
		models.OptionSet(models.MAIL_PORT, vars[models.MAIL_PORT])
		models.OptionSet(models.MAIL_USER, vars[models.MAIL_USER])
		models.OptionSet(models.MAIL_PWD, vars[models.MAIL_PWD])
		models.OptionSet(models.WEB_NOTICE, vars[models.WEB_NOTICE])
	}

	mailHost := models.OptionGet(models.MAIL_HOST, "smtp.163.com")
	this.Data[models.MAIL_HOST] = mailHost

	mailPort := models.OptionGet(models.MAIL_PORT, "25")
	this.Data[models.MAIL_PORT] = mailPort

	mailUser := models.OptionGet(models.MAIL_USER, "")
	this.Data[models.MAIL_USER] = mailUser

	mailPwd := models.OptionGet(models.MAIL_PWD, "")
	this.Data[models.MAIL_PWD] = mailPwd

	mailQueueSize := models.OptionGet(models.MAIL_QUEUE_SIZE, "100")
	this.Data[models.MAIL_QUEUE_SIZE] = mailQueueSize

	this.display()
}
