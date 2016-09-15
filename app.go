package flexvol

import (
	"github.com/urfave/cli"
)

var volumeDriver FlexDriver

func NewApp(backend FlexDriver) *cli.App {
	volumeDriver = backend

	app := cli.NewApp()
	app.Name = backend.Name()
	app.Usage = "Flex volume for Secrets"
	app.Commands = []cli.Command{
		InitCommand(),
		AttachCommand(),
		DetachCommand(),
		MountCommand(),
		UnmountCommand(),
	}

	return app
}
