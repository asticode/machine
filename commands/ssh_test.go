package commands

import (
	"testing"

	"github.com/asticode/machine/commands/commandstest"
	"github.com/asticode/machine/drivers/fakedriver"
	"github.com/asticode/machine/libmachine"
	"github.com/asticode/machine/libmachine/drivers"
	"github.com/asticode/machine/libmachine/host"
	"github.com/asticode/machine/libmachine/libmachinetest"
	"github.com/asticode/machine/libmachine/ssh"
	"github.com/asticode/machine/libmachine/ssh/sshtest"
	"github.com/asticode/machine/libmachine/state"
	"github.com/stretchr/testify/assert"
)

type FakeSSHClientCreator struct {
	client ssh.Client
}

func (fsc *FakeSSHClientCreator) CreateSSHClient(d drivers.Driver) (ssh.Client, error) {
	if fsc.client == nil {
		fsc.client = &sshtest.FakeClient{}
	}
	return fsc.client, nil
}

func TestCmdSSH(t *testing.T) {
	testCases := []struct {
		commandLine   CommandLine
		api           libmachine.API
		expectedErr   error
		helpShown     bool
		clientCreator host.SSHClientCreator
		expectedShell []string
	}{
		{
			commandLine: &commandstest.FakeCommandLine{
				CliArgs: []string{"-h"},
			},
			api:         &libmachinetest.FakeAPI{},
			expectedErr: nil,
			helpShown:   true,
		},
		{
			commandLine: &commandstest.FakeCommandLine{
				CliArgs: []string{"--help"},
			},
			api:         &libmachinetest.FakeAPI{},
			expectedErr: nil,
			helpShown:   true,
		},
		{
			commandLine: &commandstest.FakeCommandLine{
				CliArgs: []string{},
			},
			api:         &libmachinetest.FakeAPI{},
			expectedErr: ErrNoDefault,
		},
		{
			commandLine: &commandstest.FakeCommandLine{
				CliArgs: []string{"default", "df", "-h"},
			},
			api: &libmachinetest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "default",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
						},
					},
				},
			},
			expectedErr:   nil,
			clientCreator: &FakeSSHClientCreator{},
			expectedShell: []string{"df", "-h"},
		},
		{
			commandLine: &commandstest.FakeCommandLine{
				CliArgs: []string{"default"},
			},
			api: &libmachinetest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "default",
						Driver: &fakedriver.Driver{
							MockState: state.Stopped,
						},
					},
				},
			},
			expectedErr: errStateInvalidForSSH{"default"},
		},
	}

	for _, tc := range testCases {
		host.SetSSHClientCreator(tc.clientCreator)

		err := cmdSSH(tc.commandLine, tc.api)
		assert.Equal(t, err, tc.expectedErr)

		if fcl, ok := tc.commandLine.(*commandstest.FakeCommandLine); ok {
			assert.Equal(t, tc.helpShown, fcl.HelpShown)
		}

		if fcc, ok := tc.clientCreator.(*FakeSSHClientCreator); ok {
			assert.Equal(t, tc.expectedShell, fcc.client.(*sshtest.FakeClient).ActivatedShell)
		}
	}
}
