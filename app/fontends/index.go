package fontends

import (
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/libs"
	"github.com/midoks/novelsearch/app/models"
	"strings"
)

type IndexController struct {
	CommonController
}

func (this *IndexController) Index() {
	this.display()
}

func (this *IndexController) Top() {
	this.display()
}

func (this *IndexController) Baidutop() {
	list, err := libs.GetAllBaiduTop()

	if err == nil {
		this.Data["list"] = list
	} else {
		// this.Data["list"] = make(map[string]interface{})
	}

	this.display()
}

func (this *IndexController) Soso() {
	kw := this.GetString("wd")
	if strings.EqualFold(kw, "") {
		this.redirect("/")
	}
	fmt.Println(kw)
	list := models.SosoNovelByKw(kw)
	fmt.Println(list)
	this.display()
}

func (this *IndexController) Page() {
	this.display()
}

func (this *IndexController) Details() {
	unique_id := this.Ctx.Input.Param(":unique_id")
	fmt.Println(unique_id)
	this.display()
}
