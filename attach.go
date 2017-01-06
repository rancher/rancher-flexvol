package flexvol

import (
	"encoding/json"

	"github.com/urfave/cli"
)

func AttachCommand() cli.Command {
	return cli.Command{
		Name:   "attach",
		Usage:  "Attach flex volume",
		Action: handleErr(AttachVol),
	}
}

func AttachVol(c *cli.Context) error {
	params := map[string]interface{}{}

	if len(c.Args()) > 0 {
		if err := json.Unmarshal([]byte(c.Args()[0]), &params); err != nil {
			return err
		}

		device, err := volumeDriver.Attach(params)
		if err != nil {
			return err
		}
		Device(device).Print()
		return nil
	}

	return ErrIncorrectArgNumber
}
