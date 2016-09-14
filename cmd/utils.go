package cmd

import (
	"errors"

	"github.com/rancher/rancher-flexvol/flexvol"
	"github.com/rancher/rancher-flexvol/secrets"
)

func GetFlexVol(volType string) (flexvol.FlexDriver, error) {
	if volType == "secrets" {
		return &secrets.FlexVolume{}, nil
	}
	return nil, errors.New("Invalid Volume Type")
}
