package host

import "github.com/asticode/machine/libmachine/drivers"

type V2 struct {
	ConfigVersion int
	Driver        drivers.Driver
	DriverName    string
	HostOptions   *Options
	Name          string
}
