package backends

import (
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	_ "runtime"
	"strconv"
	"strings"
	_ "time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Index() {

	errMsg := ""
	if this.isPost() {

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")

		if username != "" && password != "" {
			user, err := models.UserGetByName(username)
			// this.D(user.Password)

			if err != nil || user.Password != libs.Md5([]byte(password)) {
				errMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errMsg = "该帐号已禁用"
			} else {

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password))
				if remember == "on" {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
				}

				this.user = user
				this.uLog("登录成功!")
				this.redirect(beego.URLFor("IndexController.Index"))
			}
		}
		//this.D(username, password, remember)
	}

	this.Data["errMsg"] = errMsg
	this.TplName = "backends/" + this.controllerName + "/" + this.actionName + ".html"
}

// 退出登录
func (this *LoginController) Out() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("LoginController.Index"))
}
