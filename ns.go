package main

import (
	"embed"
	"os"

	"github.com/urfave/cli"

	"github.com/midoks/novelsearch/internal/cmd"
	"github.com/midoks/novelsearch/internal/conf"
)

const Version = "2.0.0"
const AppName = "novelsearch"

//go:embed templates/*
var viewsfs embed.FS

func init() {
	conf.App.Version = Version
	conf.App.Name = AppName
	conf.App.TemplateFs = viewsfs
}

func main() {

	app := cli.NewApp()
	app.Name = conf.App.Name
	app.Version = conf.App.Version
	app.Usage = "A NovalSearch service"
	app.Commands = []cli.Command{
		cmd.Service,
		cmd.Robot,
	}

	app.Run(os.Args)

}
