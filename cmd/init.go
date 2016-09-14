package cmd

import (
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func InitCommand() cli.Command {
	return cli.Command{
		Name:   "init",
		Usage:  "Initialize flex volume",
		Action: InitVol,
	}
}

func InitVol(c *cli.Context) error {
	volType := path.Base(os.Args[0])
	volumeDriver, err := GetFlexVol(volType)
	if err != nil {
		logrus.Error(err)
		return err
	}

	output, err := volumeDriver.Init()
	if err != nil {
		return err
	}

	output.Print()

	return nil
}
