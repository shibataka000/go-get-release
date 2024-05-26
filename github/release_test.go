package github

import (
	"context"
	"os"
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

func TestLatestRelease(t *testing.T) {
	tests := []struct {
		name    string
		repo    Repository
		release Release
	}{
		{
			name:    "shibataka000/go-get-release-test",
			repo:    newRepository("shibataka000", "go-get-release-test"),
			release: newRelease("v0.0.2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewReleaseRepository(ctx, os.Getenv("GITHUB_TOKEN"))
			release, err := repository.latest(ctx, tt.repo)
			require.NoError(err)
			require.Equal(tt.release, release)
		})
	}
}
