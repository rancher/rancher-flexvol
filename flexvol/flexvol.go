package flexvol

type FlexDriver interface {
	Init() (*DriverOutput, error)
	Attach(string) (*DriverOutput, error)
	Detach(string) (*DriverOutput, error)
	Mount(string, string, string) (*DriverOutput, error)
	Unmount(string) (*DriverOutput, error)
}
