package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/midoks/novelsearch/internal/cmd"
	"github.com/midoks/novelsearch/internal/conf"
)

const Version = "2.0.0"
const AppName = "novelsearch"

func init() {
	conf.App.Version = Version
	conf.App.Name = AppName
}

func main() {

	app := cli.NewApp()
	app.Name = conf.App.Name
	app.Version = conf.App.Version
	app.Usage = "A NovalSearch service"
	app.Commands = []cli.Command{
		cmd.Service,
	}

	app.Run(os.Args)

}
