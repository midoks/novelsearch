package fontends

import (
	// "fmt"
	_ "github.com/astaxie/beego"
	// "github.com/midoks/novelsearch/app/models"
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

func (this *IndexController) Soso() {
	this.display()
}

func (this *IndexController) Page() {
	this.display()
}

func (this *IndexController) Details() {
	this.display()
}
