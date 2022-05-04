package cmd

import (
	"github.com/urfave/cli"
	// "github.com/midoks/novelsearch/internal/app"
	// "github.com/midoks/novelsearch/internal/app/router"
	// "github.com/midoks/novelsearch/internal/conf"

	"github.com/astaxie/beego"

	"github.com/midoks/novelsearch/app/crontab"
	"github.com/midoks/novelsearch/app/libs"
	_ "github.com/midoks/novelsearch/app/routers"
)

var Service = cli.Command{
	Name:        "service",
	Usage:       "This command starts all services",
	Description: `Start DHT services`,
	Action:      runAllService,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runAllService(c *cli.Context) error {

	// err := router.Init("")
	// if err != nil {
	// 	return err
	// }

	// app.Start(conf.Web.HttpPort)

	libs.Init()
	crontab.Init()
	beego.Run()

	return nil
}
