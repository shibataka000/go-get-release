package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecBinaryUnmarshalCSV(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		execBinary ExecBinary
	}{
		{
			name:       "terraform",
			value:      "terraform",
			execBinary: newExecBinary("terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			execBinary := ExecBinary{}
			execBinary.UnmarshalCSV(tt.value)
			require.Equal(tt.execBinary, execBinary)
		})
	}
}

// UnmarshalCSV converts the CSV string as executable binary.
func (b *ExecBinary) UnmarshalCSV(value string) error {
	b.name = value
	return nil
}
