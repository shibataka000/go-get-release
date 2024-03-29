package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlatformEquals(t *testing.T) {
	tests := []struct {
		name   string
		p1     Platform
		p2     Platform
		equals bool
	}{
		{
			name:   "equals",
			p1:     NewPlatform("linux", "amd64"),
			p2:     NewPlatform("linux", "amd64"),
			equals: true,
		},
		{
			name:   "not equals",
			p1:     NewPlatform("linux", "amd64"),
			p2:     NewPlatform("windows", "amd64"),
			equals: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			equals := tt.p1.Equals(tt.p2)
			assert.Equal(tt.equals, equals)
		})
	}
}
