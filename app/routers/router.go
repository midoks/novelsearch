package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/backends"
	"github.com/midoks/novelsearch/app/fontends"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"html/template"
	"net/http"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	// this.Redirect(url, 302)
	// r.Write([]byte("Location: /"))
	t.Execute(rw, data)
}

func init() {
	//错误页面设置
	beego.ErrorHandler("404", page_not_found)

	//前台
	beego.Router("/", &fontends.IndexController{}, "*:Index")
	beego.Router("/baidutop.html", &fontends.IndexController{}, "*:Baidutop")
	beego.Router("/s", &fontends.IndexController{}, "*:Soso")
	beego.Router("/b/:unique_id(.*).(html|htm|shtml)", &fontends.IndexController{}, "*:Details")

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
