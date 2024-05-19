package runtime

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
			name: "gh_2.21.0_linux_amd64.tar.gz",
			os:   "linux",
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
