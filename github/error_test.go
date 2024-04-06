package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		errStr string
	}{
		{
			name:   "NotFoundError",
			err:    NewNotFoundError(""),
			errStr: "",
		},
		{
			name:   "NotFoundError",
			err:    NewNotFoundError("value '%s' was not found in %v", "x", []string{"a", "b", "c"}),
			errStr: "value 'x' was not found in [a b c]",
		},
		{
			name:   "InvalidSemVerError",
			err:    NewInvalidSemVerError("x.y.z"),
			errStr: "invalid semver: x.y.z",
		},
		{
			name:   "UnsupportedFileFormatError",
			err:    NewUnsupportedFileFormatError(".tar.gz"),
			errStr: "unsupported file format: .tar.gz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.EqualError(tt.err, tt.errStr)
		})
	}
}
