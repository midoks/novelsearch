package cmd

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/urfave/cli"

	"github.com/midoks/novelsearch/internal/conf"
	"github.com/midoks/novelsearch/internal/router"
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

	// libs.Init()
	// crontab.Init()
	// beego.Run()

	// engine := html.New("templates", ".html")
	engine := html.NewFileSystem(http.FS(conf.App.TemplateFs), ".html")
	engine.AddFunc(
		// add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "public")

	app.Get("/", router.Home)
	app.Get("/hl", router.Hello)

	app.Listen(fmt.Sprintf(":%s", conf.Web.HttpPort))

	return nil
}
