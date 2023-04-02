package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewExecBinaryMetaWithPlatform(t *testing.T) {
	tests := []struct {
		name   string
		expect ExecBinaryMeta
		actual ExecBinaryMeta
	}{
		{
			name:   "terraform",
			expect: NewExecBinaryMetaWithPlatform("terraform", NewPlatform("linux", "amd64")),
			actual: NewExecBinaryMeta("terraform"),
		},
		{
			name:   "terraform.exe",
			expect: NewExecBinaryMetaWithPlatform("terraform", NewPlatform("windows", "amd64")),
			actual: NewExecBinaryMeta("terraform.exe"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.actual, tt.expect)
		})
	}
}
