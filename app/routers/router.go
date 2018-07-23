package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/backends"
	"github.com/midoks/novelsearch/app/fontends"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
)

func init() {

	//前台
	beego.Router("/", &fontends.IndexController{}, "*:Index")
	beego.Router("/top", &fontends.IndexController{}, "*:Top")
	beego.Router("/s", &fontends.IndexController{}, "*:Soso")
	beego.Router("/d", &fontends.IndexController{}, "*:Details")

	//前台接口
	ns := beego.NewNamespace("/v1", beego.NSAutoRouter(&backends.ItemController{}))
	beego.AddNamespace(ns)

	//后台安装
	var admin_path = "admin"
	ok, _ := libs.PathExists("conf/app.conf")
	if ok {
		models.Init()
		admin_name := beego.AppConfig.String("admin_path")
		admin_path = fmt.Sprintf("/%s", admin_name)
	}
	//后台
	beego.Router(admin_path, &backends.IndexController{}, "*:Index")
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.LoginController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.IndexController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.InstallController{})))

	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysUserController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysFuncController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysRoleController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysLogController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysSettingController{})))

	//后台功能开发
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.AppItemController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.AppDebugController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.AppNovelController{})))
}
