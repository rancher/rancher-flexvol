package flexvol

import (
	"errors"

	"github.com/rancher/rancher-flexvol/flexvol/drivers/secrets"
	"github.com/rancher/rancher-flexvol/flexvol/responsetypes"
)

type FlexDriver interface {
	Init() (*responsetypes.DriverOutput, error)
	Attach(map[string]interface{}) (*responsetypes.DriverOutput, error)
	Detach(string) (*responsetypes.DriverOutput, error)
	Mount(string, string, string) (*responsetypes.DriverOutput, error)
	Unmount(string) (*responsetypes.DriverOutput, error)
}

func NewFlexVol(volType string) (FlexDriver, error) {
	switch volType {
	case "rancher-secrets":
		return &secrets.FlexVolume{}, nil
	default:
		return nil, errors.New("Invalid Volume type")
	}
}
