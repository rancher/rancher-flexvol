package flexvol

import (
	"encoding/json"

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
	params := map[string]interface{}{}

	if len(c.Args()) > 0 {
		json.Unmarshal([]byte(c.Args()[0]), params)
		output, err := volumeDriver.Attach(params)
		if err != nil {
			output.Print()
			return err
		}
		output.Print()
	}

	return nil
}
