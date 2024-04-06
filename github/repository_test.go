package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepositoryEqual(t *testing.T) {
	tests := []struct {
		name  string
		r1    Repository
		r2    Repository
		equal bool
	}{
		{
			name:  "Equal",
			r1:    NewRepository("shibataka000", "go-get-release"),
			r2:    NewRepository("shibataka000", "go-get-release"),
			equal: true,
		},
		{
			name:  "NotEqual",
			r1:    NewRepository("shibataka000", "go-get-release"),
			r2:    NewRepository("shibataka000", "go-get-release-2"),
			equal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			equal := tt.r1.Equal(tt.r2)
			require.Equal(tt.equal, equal)
		})
	}
}
