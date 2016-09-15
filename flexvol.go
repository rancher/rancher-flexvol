package flexvol

type FlexDriver interface {
	Name() string
	Init() (*responsetypes.DriverOutput, error)
	Attach(map[string]interface{}) (*responsetypes.DriverOutput, error)
	Detach(string) (*responsetypes.DriverOutput, error)
	Mount(string, string, string) (*responsetypes.DriverOutput, error)
	Unmount(string) (*responsetypes.DriverOutput, error)
}
