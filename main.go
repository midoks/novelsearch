package main

import (
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/crontab"
	"github.com/midoks/novelsearch/app/libs"
	_ "github.com/midoks/novelsearch/app/routers"
)

func main() {
	libs.Init()
	crontab.Init()
	beego.Run()
}
