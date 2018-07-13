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
