package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPatternMatch(t *testing.T) {
	tests := []struct {
		name    string
		pattern Pattern
		asset   Asset
		match   bool
	}{
		{
			name:    "gh_2.52.0_linux_amd64.tar.gz",
			pattern: mustNewPatternFromString("gh_.*_linux_amd64.tar.gz", "gh"),
			asset:   newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
			match:   true,
		},
		{
			name:    "gh_2.52.0_linux_arm64.tar.gz",
			pattern: mustNewPatternFromString("gh_.*_linux_amd64.tar.gz", "gh"),
			asset:   newAsset(0, "gh_2.52.0_linux_arm64.tar.gz"),
			match:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.match, tt.pattern.match(tt.asset))
		})
	}
}

func TestPatternExecute(t *testing.T) {
	tests := []struct {
		name       string
		pattern    Pattern
		asset      Asset
		execBinary ExecBinary
	}{
		{
			name:       "gh_2.52.0_linux_amd64.tar.gz",
			pattern:    mustNewPatternFromString(`(?P<name>\w+)_[\d\.]+_linux_amd64.tar.gz`, "{{.name}}"),
			asset:      newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
			execBinary: newExecBinary("gh"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			execBinary, err := tt.pattern.execute(tt.asset)
			require.NoError(err)
			require.Equal(tt.execBinary, execBinary)
		})
	}
}
