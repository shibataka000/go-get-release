package github

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReleaseSemVer(t *testing.T) {
	factory := NewReleaseFactory()

	tests := []struct {
		name    string
		release Release
		semver  string
	}{
		{
			name:    "v1.2.3",
			release: factory.new("v1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "1.2.3",
			release: factory.new("1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "x.y.z",
			release: factory.new("x.y.z"),
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
	repositoryFactory := NewRepositoryFactory()
	releaseFactory := NewReleaseFactory()

	tests := []struct {
		name    string
		repo    Repository
		release Release
	}{
		{
			name:    "shibataka000/go-get-release-test",
			repo:    repositoryFactory.new("shibataka000", "go-get-release-test"),
			release: releaseFactory.new("v0.0.2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewReleaseRepository(ctx, os.Getenv("GITHUB_TOKEN"), releaseFactory)
			release, err := repository.latest(ctx, tt.repo)
			require.NoError(err)
			require.Equal(tt.release, release)
		})
	}
}
