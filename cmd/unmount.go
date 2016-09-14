package cmd

import (
	"os"
	"path"

	"github.com/rancher/rancher-flexvol/flexvol"
	"github.com/urfave/cli"
)

func UnmountCommand() cli.Command {
	return cli.Command{
		Name:   "mount",
		Usage:  "Unmount flex volume",
		Action: UnmountVol,
	}
}

func UnmountVol(c *cli.Context) error {
	volType := path.Base(os.Args[0])

	volumeDriver, err := flexvol.NewFlexVol(volType)
	if err != nil {
		return err
	}

	if len(c.Args()) > 2 {
		output, err := volumeDriver.Unmount(c.Args()[0])
		if err != nil {
			output.Print()
			return err
		}
		output.Print()
	}

	return nil
}
