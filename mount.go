package flexvol

import (
	"encoding/json"

	"github.com/urfave/cli"
)

func MountCommand() cli.Command {
	return cli.Command{
		Name:   "mount",
		Usage:  "Mount flex volume",
		Action: handleErr(MountVol),
	}
}

func MountVol(c *cli.Context) error {
	params := map[string]interface{}{}

	if len(c.Args()) > 0 {
		if err := json.Unmarshal([]byte(c.Args()[2]), &params); err != nil {
			return err
		}
		if err := volumeDriver.Mount(c.Args()[0], c.Args()[1], params); err != nil {
			return err
		}
		Success().Print()
		return nil
	}

	return ErrIncorrectArgNumber
}
