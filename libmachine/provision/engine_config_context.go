package provision

import (
	"github.com/asticode/machine/libmachine/auth"
	"github.com/asticode/machine/libmachine/engine"
)

type EngineConfigContext struct {
	DockerPort       int
	AuthOptions      auth.Options
	EngineOptions    engine.Options
	DockerOptionsDir string
}
