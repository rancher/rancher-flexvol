package secrets

import (
	"github.com/Sirupsen/logrus"
	"github.com/rancher/rancher-flexvol/flexvol"
)

type FlexVolume struct{}

func (sv *FlexVolume) Init() (*flexvol.DriverOutput, error) {
	return &flexvol.DriverOutput{Status: "Success"}, nil
}

func (sv *FlexVolume) Attach(params string) (*flexvol.DriverOutput, error) {
	status := "Failure"
	output := &flexvol.DriverOutput{}

	logrus.Errorf("PARAM: %s", params)

	output.Status = status

	return output, nil
}

func (sv *FlexVolume) Detach(device string) (*flexvol.DriverOutput, error) {
	status := "Failure"
	output := &flexvol.DriverOutput{}

	logrus.Errorf("DEVICE: %s", device)

	output.Status = status

	return output, nil
}

func (sv *FlexVolume) Mount(dir, device, params string) (*flexvol.DriverOutput, error) {
	status := "Failure"
	output := &flexvol.DriverOutput{}

	output.Status = status
	logrus.Errorf("Dir: %s, DEV: %s, PARAMS: %s", dir, device, params)

	return output, nil
}

func (sv *FlexVolume) Unmount(dir string) (*flexvol.DriverOutput, error) {
	status := "Failure"
	output := &flexvol.DriverOutput{}

	logrus.Errorf("Dir: %s", dir)
	output.Status = status

	return output, nil
}
