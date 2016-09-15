package flexvol

type FlexDriver interface {
	Name() string
	Init() (*DriverOutput, error)
	Attach(map[string]interface{}) (*DriverOutput, error)
	Detach(string) (*DriverOutput, error)
	Mount(string, string, string) (*DriverOutput, error)
	Unmount(string) (*DriverOutput, error)
}
