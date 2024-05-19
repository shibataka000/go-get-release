package github

import (
	"testing"

	"github.com/shibataka000/go-get-release/file"
	"github.com/stretchr/testify/require"
)

func TestExecutableBinaryName(t *testing.T) {
	tests := []struct {
		name           string
		execBinary     ExecutableBinaryMeta
		execBinaryName file.Name
	}{
		{
			name:           "linux",
			execBinary:     newExecutableBinaryMeta("app", "linux"),
			execBinaryName: "app",
		},
		{
			name:           "linux",
			execBinary:     newExecutableBinaryMeta("app", "windows"),
			execBinaryName: "app.exe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.execBinary.Name(), tt.execBinaryName)
		})
	}
}

func TestNewExecutableBinaryMetaFromRepository(t *testing.T) {
	tests := []struct {
		name           string
		execBinary     ExecutableBinaryMeta
		execBinaryName file.Name
	}{
		{
			name:           "linux",
			execBinary:     newExecutableBinaryMeta("app", "linux"),
			execBinaryName: "app",
		},
		{
			name:           "linux",
			execBinary:     newExecutableBinaryMeta("app", "windows"),
			execBinaryName: "app.exe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.execBinary.Name(), tt.execBinaryName)
		})
	}
}
