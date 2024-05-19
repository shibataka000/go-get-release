package platform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGuess(t *testing.T) {
	tests := []struct {
		name string
		os   OS
		arch Arch
	}{
		{
			name: "linux_amd64",
			os:   "linux",
			arch: "amd64",
		},
		{
			name: "dragonfly_js_mips64le_arm",
			os:   "dragonfly",
			arch: "mips64le",
		},
		{
			name: "",
			os:   "unknown",
			arch: "amd64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			os, arch := Guess(tt.name)
			require.Equal(tt.os, os)
			require.Equal(tt.arch, arch)
		})
	}
}
