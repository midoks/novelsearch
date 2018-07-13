package backends

import (
	_ "github.com/midoks/novelsearch/app/models"
)

type BaseController struct {
	CommonController
}

func (this *BaseController) Prepare() {

	this.initData()
	this.auth()
}
