// +build !darwin

package vmwarefusion

import "github.com/asticode/machine/libmachine/drivers"

func NewDriver(hostName, storePath string) drivers.Driver {
	return drivers.NewDriverNotSupported("vmwarefusion", hostName, storePath)
}
