package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlatformEqual(t *testing.T) {
	tests := []struct {
		name  string
		p1    Platform
		p2    Platform
		equal bool
	}{
		{
			name:  "Equal",
			p1:    NewPlatform("linux", "amd64"),
			p2:    NewPlatform("linux", "amd64"),
			equal: true,
		},
		{
			name:  "NotEqual",
			p1:    NewPlatform("linux", "amd64"),
			p2:    NewPlatform("windows", "amd64"),
			equal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			equal := tt.p1.Equal(tt.p2)
			require.Equal(tt.equal, equal)
		})
	}
}
