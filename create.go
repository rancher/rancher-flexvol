package flexvol

import (
	"encoding/json"

	"github.com/urfave/cli"
)

func CreateCommand() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "Create a flex volume",
		Action: handleErr(CreateVol),
	}
}

func CreateVol(c *cli.Context) error {
	params := map[string]interface{}{}

	if len(c.Args()) > 0 {
		if err := json.Unmarshal([]byte(c.Args()[0]), &params); err != nil {
			return err
		}

		data, err := rancherVolumeDriver.Create(params)
		if err != nil {
			return err
		}

		Options(data).Print()
		return nil
	}

	return ErrIncorrectArgNumber
}
