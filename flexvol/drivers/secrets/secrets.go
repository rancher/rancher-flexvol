package secrets

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/mount"
	"github.com/rancher/rancher-flexvol/flexvol/responsetypes"
)

const (
	volRoot = "/var/lib/rancher/volumes"
)

type FlexVolume struct{}

func (sv *FlexVolume) Init() (*responsetypes.DriverOutput, error) {
	return &responsetypes.DriverOutput{Status: "Success"}, nil
}

func (sv *FlexVolume) Attach(params map[string]interface{}) (*responsetypes.DriverOutput, error) {
	var volName string
	var ok bool

	//Default volume mode
	mode := int(0700)
	status := "Failure"
	mountOpts := "size=10m"
	output := &responsetypes.DriverOutput{}

	if volName, ok = params["volumeName"].(string); !ok {
		output.Message = "No volumeName set"
	}

	if uMode, ok := params["mode"].(int); ok {
		mode = int(uMode)
	}

	if mOpts, ok := params["mountOpts"].(string); ok {
		mountOpts = mOpts
	}

	path := fmt.Sprintf("%s/%s", volRoot, volName)

	os.Mkdir(path, os.FileMode(mode))

	if err := mount.Mount("tmpfs", path, "tmpfs", mountOpts); err == nil {
		status = "Success"
	}

	output.Status = status
	output.Device = path

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
