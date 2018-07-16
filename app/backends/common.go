package backends

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
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
	user           *models.SysUser

	// xsrf data
	_xsrfToken string
	XSRFExpire int
	EnableXSRF bool
}

func (this *CommonController) uLog(behavior string) {
	models.LogAdd(this.user.Id, 1, behavior)
}

func (this *CommonController) dLog(behavior string) {
	models.DebugAdd(1, behavior)
}

func (this *CommonController) D(args ...string) {
	if beego.AppConfig.String("runmode") == "dev" {
		for i := 0; i < len(args); i++ {
			this.Ctx.WriteString(args[i])
		}
		//this.StopRun()
	}
}

func (this *CommonController) initXSRF() {
	this.EnableXSRF = true
	this._xsrfToken = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	this.XSRFExpire = 3600 //过期时间，默认1小时
}

func (this *CommonController) initBaseData() {
	this.Data["pageStartTime"] = time.Now()
	this.pageSize = 10
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
}

func (this *CommonController) initConfData() {

	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
}

//登录状态验证
func (this *CommonController) auth() {

	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	//fmt.Println(arr)

	if len(arr) == 2 {

		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)

		if userId > 0 {
			user, err := models.UserGetById(userId)
			if err == nil && password == libs.Md5([]byte(this.getClientIp()+"|"+user.Password)) {

				this.user = user
				this.Data["user"] = user
				role, _ := models.RoleGetById(user.Roleid)
				//fmt.Println(role)

				menuNav, curMenuName, curMenuFuncName, isAuth := models.FuncGetNav(this.controllerName, this.actionName)

				this.Data["curMenuName"] = curMenuName
				this.Data["curMenuFuncName"] = curMenuFuncName

				newMenuNav := make([]models.SysFuncNav, 0)

				var tmpMenuNav models.SysFuncNav
				var menuList []models.SysFunc

				for i := 0; i < len(menuNav); i++ {

					menuList = make([]models.SysFunc, 0)

					for mi := 0; mi < menuNav[i].ListCount; mi++ {
						if this.isIntInList(menuNav[i].List[mi].Id, role.List) {
							menuList = append(menuList, menuNav[i].List[mi])
						}
					}

					if len(menuList) > 0 {

						tmpMenuNav.Info = menuNav[i].Info
						tmpMenuNav.List = menuList
						tmpMenuNav.MenuOpen = menuNav[i].MenuOpen
						tmpMenuNav.ListCount = menuNav[i].ListCount
						newMenuNav = append(newMenuNav, tmpMenuNav)
					}
				}

				this.Data["menuNav"] = newMenuNav

				if !isAuth && this.controllerName != "login" && this.controllerName != "index" {
					xrw := this.Ctx.Input.Header("X-Requested-With")
					if strings.EqualFold(xrw, "XMLHttpRequest") {
						this.retFail("无权访问")
					}

					this.Ctx.WriteString("无权访问无权访问")
					this.StopRun()
					return
				}
			}
		}
	}

	//跳到登录页
	if (this.user == nil || this.user.Id == 0) && this.controllerName != "login" && (this.actionName != "out") {
		this.redirect(beego.URLFor("LoginController.Index"))
	}

	//跳到首页
	if (this.user != nil) && (this.controllerName == "login" && this.actionName == "index") {
		this.redirect(beego.URLFor("IndexController.Index"))
	}

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
		tplname = "backends/" + tpl[0] + ".html"
	} else if len(tpl) == 2 {
		tplname = "backends/" + tpl[0] + "/" + tpl[1] + ".html"
	} else {
		tplname = "backends/" + this.controllerName + "/" + this.actionName + ".html"
	}

	this.Layout = "backends/layout/index.html"
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
