package main

import (
	"github.com/astaxie/beego"
	"github.com/midoks/novelsearch/app/libs"
	_ "github.com/midoks/novelsearch/app/routers"
)

func main() {
	libs.Init()
	beego.Run()
}
