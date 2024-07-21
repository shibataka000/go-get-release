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
	}{
		{
			name:    "v1.2.3",
			release: newRelease("v1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "1.2.3",
			release: newRelease("1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "x.y.z",
			release: newRelease("x.y.z"),
			semver:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			semver := tt.release.semver()
			require.Equal(tt.semver, semver)
		})
	}
}

func TestReleaseUnmarshalCSV(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		release Release
	}{
		{
			name:    "v1.2.3",
			value:   "v1.2.3",
			release: newRelease("v1.2.3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			release := Release{}
			release.UnmarshalCSV(tt.value)
			require.Equal(tt.release, release)
		})
	}
}

// UnmarshalCSV converts the CSV string as GitHub release.
func (r *Release) UnmarshalCSV(value string) error {
	release := newRelease(value)
	r.tag = release.tag
	return nil
}
