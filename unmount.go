package flexvol

import (
	"github.com/urfave/cli"
)

func UnmountCommand() cli.Command {
	return cli.Command{
		Name:   "unmount",
		Usage:  "Unmount flex volume",
		Action: handleErr(UnmountVol),
	}
}

func UnmountVol(c *cli.Context) error {
	if len(c.Args()) > 0 {
		err := volumeDriver.Unmount(c.Args()[0])
		if err != nil {
			return err
		}
		Success().Print()
		return nil
	}

	return ErrIncorrectArgNumber
}
