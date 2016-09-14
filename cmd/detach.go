package cmd

import (
	"os"
	"path"

	"github.com/rancher/rancher-flexvol/flexvol"
	"github.com/urfave/cli"
)

func DetachCommand() cli.Command {
	return cli.Command{
		Name:   "detach",
		Usage:  "Detach flex volume",
		Action: DetachVol,
	}
}

func DetachVol(c *cli.Context) error {
	volType := path.Base(os.Args[0])
	volumeDriver, err := flexvol.NewFlexVol(volType)
	if err != nil {
		return err
	}

	if len(c.Args()) > 0 {
		output, err := volumeDriver.Detach(c.Args()[0])
		if err != nil {
			output.Print()
			return err
		}
		output.Print()
	}

	return nil
}
