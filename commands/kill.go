package commands

import "github.com/asticode/machine/libmachine"

func cmdKill(c CommandLine, api libmachine.API) error {
	return runAction("kill", c, api)
}
