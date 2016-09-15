package flexvol

import "github.com/urfave/cli"

func InitCommand() cli.Command {
	return cli.Command{
		Name:   "init",
		Usage:  "Initialize flex volume",
		Action: InitVol,
	}
}

func InitVol(c *cli.Context) error {
	output, err := volumeDriver.Init()
	if err != nil {
		return err
	}

	output.Print()

	return nil
}
