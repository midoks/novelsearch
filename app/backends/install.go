package backends

import (
	"fmt"
)

type InstallController struct {
	BaseController
}

func (this *InstallController) Index() {

	fmt.Println(123)
	this.display()
}

//渲染模版
func (this *InstallController) display(tpl ...string) {
	var tplname string

	if len(tpl) == 1 {
		tplname = "backends/" + tpl[0] + ".html"
	} else if len(tpl) == 2 {
		tplname = "backends/" + tpl[0] + "/" + tpl[1] + ".html"
	} else {
		tplname = "backends/" + this.controllerName + "/" + this.actionName + ".html"
	}

	// this.Layout = "backends/layout/index.html"
	this.TplName = tplname
}
