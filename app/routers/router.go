package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/backends"
	"github.com/midoks/novelsearch/app/fontends"
	"github.com/midoks/novelsearch/app/models"
)

func init() {

	models.Init()

	//前台
	beego.Router("/", &fontends.IndexController{}, "*:Index")

	//前台接口
	ns := beego.NewNamespace("/v1", beego.NSAutoRouter(&backends.ItemController{}))
	beego.AddNamespace(ns)

	//后台
	admin_name := beego.AppConfig.String("admin_path")
	admin_path := fmt.Sprintf("/%s", admin_name)

	beego.Router(admin_path, &backends.IndexController{}, "*:Index")
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.LoginController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.IndexController{})))

	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysUserController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysFuncController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysRoleController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.SysLogController{})))

	//后台功能开发
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.AppItemController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.AppDebugController{})))
	beego.AddNamespace(beego.NewNamespace(admin_path, beego.NSAutoRouter(&backends.AppNovelController{})))
}
