package provision

import (
	"testing"

	"github.com/asticode/machine/drivers/fakedriver"
	"github.com/asticode/machine/libmachine/auth"
	"github.com/asticode/machine/libmachine/engine"
	"github.com/asticode/machine/libmachine/provision/provisiontest"
	"github.com/asticode/machine/libmachine/swarm"
)

func TestDebianDefaultStorageDriver(t *testing.T) {
	p := NewDebianProvisioner(&fakedriver.Driver{}).(*DebianProvisioner)
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "overlay2" {
		t.Fatal("Default storage driver should be overlay2")
	}
}
