package cmd

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/urfave/cli"

	"github.com/midoks/novelsearch/internal/conf"
	"github.com/midoks/novelsearch/internal/router"
	"github.com/midoks/novelsearch/internal/router/admin"
	"github.com/midoks/novelsearch/internal/tmpl"
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

	if conf.App.RunMode != "prod" {
		go func() {
			http.ListenAndServe(":"+conf.Debug.Port, nil)
		}()
	}

	// engine := html.New("templates", ".html")
	// fmt.Println("tmp:", conf.App.TemplateFs)
	engine := html.NewFileSystem(http.FS(conf.App.TemplateFs), ".html")

	engine.AddFunc("unescape", tmpl.Unescape)
	engine.AddFunc("adminPath", tmpl.AdminPath)
	engine.AddFunc("join", strings.Join)
	engine.AddFunc("str2html", tmpl.Str2HTML)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "public")

	app.Get("/", router.Home)
	app.Get("/s", router.Soso)
	app.Get("/hl", router.Hello)

	//admin
	adminPath := conf.Admin.AdminPath
	app.Get("/"+adminPath, admin.Admin)
	app.Get("/"+adminPath+"/spider/index", admin.SpiderList)
	app.Get("/"+adminPath+"/spider/add", admin.SpiderAdd)

	app.Listen(fmt.Sprintf(":%s", conf.Web.HttpPort))

	return nil
}
