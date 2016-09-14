package secrets

import (
	"github.com/Sirupsen/logrus"
	"github.com/rancher/rancher-flexvol/flexvol/responsetypes"
)

type FlexVolume struct{}

func (sv *FlexVolume) Init() (*responsetypes.DriverOutput, error) {
	return &responsetypes.DriverOutput{Status: "Success"}, nil
}

func (sv *FlexVolume) Attach(params string) (*responsetypes.DriverOutput, error) {
	status := "Failure"
	output := &responsetypes.DriverOutput{}

	logrus.Errorf("PARAM: %s", params)

	output.Status = status

	return output, nil
}

func (sv *FlexVolume) Detach(device string) (*responsetypes.DriverOutput, error) {
	status := "Failure"
	output := &responsetypes.DriverOutput{}

	logrus.Errorf("DEVICE: %s", device)

	output.Status = status

	return output, nil
}

func (sv *FlexVolume) Mount(dir, device, params string) (*responsetypes.DriverOutput, error) {
	status := "Failure"
	output := &responsetypes.DriverOutput{}

	output.Status = status
	logrus.Errorf("Dir: %s, DEV: %s, PARAMS: %s", dir, device, params)

	return output, nil
}

func (sv *FlexVolume) Unmount(dir string) (*responsetypes.DriverOutput, error) {
	status := "Failure"
	output := &responsetypes.DriverOutput{}

	logrus.Errorf("Dir: %s", dir)
	output.Status = status

	return output, nil
}
