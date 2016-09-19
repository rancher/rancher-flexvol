package flexvol

import "github.com/urfave/cli"

func InitCommand() cli.Command {
	return cli.Command{
		Name:   "init",
		Usage:  "Initialize flex volume",
		Action: handleErr(InitVol),
	}
}

func InitVol(c *cli.Context) error {
	if err := volumeDriver.Init(); err != nil {
		return err
	}
	Success().Print()
	return nil
}
