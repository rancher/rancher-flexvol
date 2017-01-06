package flexvol

import (
	"github.com/urfave/cli"
)

func DetachCommand() cli.Command {
	return cli.Command{
		Name:   "detach",
		Usage:  "Detach flex volume",
		Action: handleErr(DetachVol),
	}
}

func DetachVol(c *cli.Context) error {
	if len(c.Args()) > 0 {
		err := volumeDriver.Detach(c.Args()[0])
		if err != nil {
			return err
		}
		Success().Print()
		return nil
	}

	return ErrIncorrectArgNumber
}
