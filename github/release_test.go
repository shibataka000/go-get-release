package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReleaseSemVer(t *testing.T) {
	tests := []struct {
		name    string
		release Release
		semver  string
		err     error
	}{
		{
			name:    "v1.2.3",
			release: NewRelease("v1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "1.2.3",
			release: NewRelease("1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "x.y.z",
			release: NewRelease("x.y.z"),
			err:     NewInvalidSemVerError("x.y.z"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			semver, err := tt.release.SemVer()
			if tt.err == nil {
				require.NoError(err)
				require.Equal(tt.semver, semver)
			} else {
				require.EqualError(err, tt.err.Error())
			}
		})
	}
}
