package cmd

import (
	"fmt"
	"html/template"
	"net/http"
	_ "net/http/pprof"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/urfave/cli"

	"github.com/midoks/novelsearch/internal/conf"
	"github.com/midoks/novelsearch/internal/router"
	"github.com/midoks/novelsearch/internal/router/admin"
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
	app.Get("/s", router.Soso)
	app.Get("/hl", router.Hello)

	//admin
	adminPath := conf.Admin.AdminPath

	fmt.Println(adminPath)
	app.Get("/"+adminPath, admin.Admin)

	app.Listen(fmt.Sprintf(":%s", conf.Web.HttpPort))

	return nil
}
