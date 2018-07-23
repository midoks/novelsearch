package fontends

import (
	// "fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	// "github.com/midoks/novelsearch/app/libs"
	// "github.com/midoks/novelsearch/app/models"
	"strconv"
	"strings"
	"time"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type CommonController struct {
	beego.Controller
	controllerName string
	actionName     string
	pageSize       int

	// xsrf data
	_xsrfToken string
	XSRFExpire int
	EnableXSRF bool
}

func (this *CommonController) Prepare() {
	this.initData()
}

func (this *CommonController) D(args ...string) {
	if beego.AppConfig.String("runmode") == "dev" {
		for i := 0; i < len(args); i++ {
			this.Ctx.WriteString(args[i])
		}
		//this.StopRun()
	}
}

func (this *CommonController) initData() {

	this.Data["pageStartTime"] = time.Now()
	this.pageSize = 10
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	//println(this.controllerName, this.actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
}

func (this *CommonController) isIntInList(check int, list string) (out bool) {
	out = false
	numList := strings.Split(list, ",")
	for i := 0; i < len(numList); i++ {
		if numList[i] == strconv.Itoa(check) {
			out = true
		}
	}
	return out
}

// 是否POST提交
func (this *CommonController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

// 重定向
func (this *CommonController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//获取用户IP地址
func (this *CommonController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//渲染模版
func (this *CommonController) display(tpl ...string) {
	var tplname string

	if len(tpl) == 1 {
		tplname = "fontends/" + tpl[0] + ".html"
	} else if len(tpl) == 2 {
		tplname = "fontends/" + tpl[0] + "/" + tpl[1] + ".html"
	} else {
		tplname = "fontends/" + this.controllerName + "/" + this.actionName + ".html"
	}

	this.Layout = "fontends/layout/index.html"
	this.TplName = tplname
}

// 输出json
func (this *CommonController) retJson(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *CommonController) retResult(code int, msg interface{}, data ...interface{}) {
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg

	if len(data) > 0 {
		out["data"] = data
	}

	this.retJson(out)
}

func (this *CommonController) retOk(msg interface{}, data ...interface{}) {
	this.retResult(MSG_OK, msg, data...)
}

func (this *CommonController) retFail(msg interface{}, data ...interface{}) {
	this.retResult(MSG_ERR, msg, data...)
}
