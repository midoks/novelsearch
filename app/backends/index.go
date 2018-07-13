package backends

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {

	o, err := models.FuncGetById(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(o)
	}

	this.display()
}
