package cmd

import (
	"os"
	"path"

	"github.com/rancher/rancher-flexvol/flexvol"
	"github.com/urfave/cli"
)

func AttachCommand() cli.Command {
	return cli.Command{
		Name:   "attach",
		Usage:  "Attach flex volume",
		Action: AttachVol,
	}
}

func AttachVol(c *cli.Context) error {
	volType := path.Base(os.Args[0])
	volumeDriver, err := flexvol.NewFlexVol(volType)
	if err != nil {
		return err
	}

	if len(c.Args()) > 0 {
		output, err := volumeDriver.Attach(c.Args()[0])
		if err != nil {
			output.Print()
			return err
		}
		output.Print()
	}

	return nil
}
