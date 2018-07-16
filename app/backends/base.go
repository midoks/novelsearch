package backends

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/libs"
	_ "github.com/midoks/novelsearch/app/models"
	"strings"
)

type BaseController struct {
	CommonController
}

func (this *BaseController) Prepare() {

	this.initBaseData()
	if strings.EqualFold(this.controllerName, "install") {

		ok, _ := libs.PathExists("conf/app.conf")
		if ok {
			this.redirect(beego.URLFor("IndexController.Index"))
		}
	} else if strings.EqualFold(this.controllerName, "login") {

		this.initConfData()
		this.auth()
	} else {

		ok, _ := libs.PathExists("conf/app.conf")
		if !ok {
			this.redirect(beego.URLFor("InstallController.Index"))
		} else {
			this.initConfData()
			this.auth()
		}
	}
}
