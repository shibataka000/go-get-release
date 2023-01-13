package application

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCommandFromQuery(t *testing.T) {
	tests := []struct {
		query   string
		command Command
	}{
		{
			query:   "terraform",
			command: NewCommand("", "terraform", "", "linux", "amd64", "", false),
		},
		{
			query:   "hashicorp/terraform",
			command: NewCommand("hashicorp", "terraform", "", "linux", "amd64", "", false),
		},
		{
			query:   "terraform=v1.0.0",
			command: NewCommand("", "terraform", "v1.0.0", "linux", "amd64", "", false),
		},
		{
			query:   "hashicorp/terraform=v1.0.0",
			command: NewCommand("hashicorp", "terraform", "v1.0.0", "linux", "amd64", "", false),
		},
	}

	for _, tt := range tests {
		name := tt.query
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			command, err := NewCommandFromQuery(tt.query, tt.command.platform().OS(), tt.command.platform().Arch(), tt.command.InstallDir(), tt.command.isInteractive())
			assert.NoError(err)
			assert.Equal(tt.command, command)
		})
	}
}
