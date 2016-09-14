package main

import (
	"os"
	"path"

	"github.com/rancher/rancher-flexvol/cmd"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.1-dev"

func main() {
	appName := path.Base(os.Args[0])

	app := cli.NewApp()
	app.Name = appName
	app.Version = VERSION
	app.Usage = "Flex volume for Secrets"
	app.Commands = []cli.Command{
		cmd.InitCommand(),
		cmd.AttachCommand(),
		cmd.DetachCommand(),
		cmd.MountCommand(),
		cmd.UnmountCommand(),
	}

	app.Run(os.Args)
}
