package cmd

import (
	"os"
	"path"

	"github.com/urfave/cli"
)

func MountCommand() cli.Command {
	return cli.Command{
		Name:   "mount",
		Usage:  "Mount flex volume",
		Action: MountVol,
	}
}

func MountVol(c *cli.Context) error {
	volType := path.Base(os.Args[0])
	volumeDriver, err := GetFlexVol(volType)
	if err != nil {
		return err
	}

	if len(c.Args()) > 2 {
		output, err := volumeDriver.Mount(c.Args()[0], c.Args()[1], c.Args()[2])
		if err != nil {
			output.Print()
			return err
		}
		output.Print()
	}

	return nil
}
