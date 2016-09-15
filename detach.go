package flexvol

import "github.com/urfave/cli"

func DetachCommand() cli.Command {
	return cli.Command{
		Name:   "detach",
		Usage:  "Detach flex volume",
		Action: DetachVol,
	}
}

func DetachVol(c *cli.Context) error {
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
