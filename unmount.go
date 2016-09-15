package flexvol

import "github.com/urfave/cli"

func UnmountCommand() cli.Command {
	return cli.Command{
		Name:   "mount",
		Usage:  "Unmount flex volume",
		Action: UnmountVol,
	}
}

func UnmountVol(c *cli.Context) error {
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
