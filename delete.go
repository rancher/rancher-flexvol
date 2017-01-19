package flexvol

import (
	"encoding/json"

	"github.com/urfave/cli"
)

func DeleteCommand() cli.Command {
	return cli.Command{
		Name:   "delete",
		Usage:  "Delete a flex volume",
		Action: handleErr(DeleteVol),
	}
}

func DeleteVol(c *cli.Context) error {
	params := map[string]interface{}{}

	if len(c.Args()) > 0 {
		json.Unmarshal([]byte(c.Args()[0]), &params)
		if err := rancherVolumeDriver.Delete(params); err != nil {
			return err
		}
		Success().Print()
		return nil
	}

	return ErrIncorrectArgNumber
}
