package main

import (
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/crontab"
	"github.com/midoks/novelsearch/app/libs"
	_ "github.com/midoks/novelsearch/app/routers"
	"github.com/midoks/novelsearch/app/website"
)

func main() {

	libs.Init()
	website.Init()
	crontab.Init()
	beego.Run()
}
