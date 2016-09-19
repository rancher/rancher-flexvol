package flexvol

import (
	"github.com/urfave/cli"
)

var volumeDriver FlexDriver
var rancherVolumeDriver RancherFlexDriver

func NewApp(backend FlexDriver) *cli.App {
	volumeDriver = backend

	app := cli.NewApp()
	app.Usage = "Flex volume for Secrets"
	app.Commands = []cli.Command{
		InitCommand(),
		AttachCommand(),
		DetachCommand(),
		MountCommand(),
		UnmountCommand(),
	}

	if rfd, ok := volumeDriver.(RancherFlexDriver); ok {
		rancherVolumeDriver = rfd
		app.Commands = append(app.Commands, CreateCommand(), DeleteCommand())
	}

	return app
}
