package flexvol

type FlexDriver interface {
	Init() error
	Attach(options map[string]interface{}) (device string, err error)
	Detach(device string) error
	Mount(mountDir string, device string, options map[string]interface{}) error
	Unmount(mountDir string) error
}

type RancherFlexDriver interface {
	Create(options map[string]interface{}) error
	Delete(options map[string]interface{}) error
}
