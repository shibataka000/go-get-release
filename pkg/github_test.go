package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGitHubRepositoryFullName(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		fullName   string
	}{
		{
			name:       "hashicorp/terraform",
			repository: NewGitHubRepository("hashicorp", "terraform"),
			fullName:   "hashicorp/terraform",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.fullName, tt.repository.FullName())
		})
	}
}

func TestGitHubReleaseSemVer(t *testing.T) {
	tests := []struct {
		name    string
		release GitHubRelease
		semver  string
		err     error
	}{
		{
			name:    "v1.2.3",
			release: NewGitHubRelease(0, "v1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "1.2.3",
			release: NewGitHubRelease(0, "1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "x.y.z",
			release: NewGitHubRelease(0, "x.y.z"),
			semver:  "",
			err:     fmt.Errorf("x.y.z is not valid semver"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			semver, err := tt.release.SemVer()
			if tt.err == nil {
				assert.NoError(err)
				assert.Equal(tt.semver, semver)
			} else {
				assert.EqualError(err, tt.err.Error())
			}
		})
	}
}

func TestGitHubAssetHasExecBinary(t *testing.T) {
	tests := []struct {
		name          string
		asset         GitHubAsset
		hasExecBinary bool
	}{
		{
			name:          "https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz",
			asset:         NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			hasExecBinary: true,
		},
		{
			name:          "https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_checksums.txt",
			asset:         NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_checksums.txt"),
			hasExecBinary: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.hasExecBinary, tt.asset.HasExecBinary())
		})
	}
}

func TestGitHubAssetPlatform(t *testing.T) {
	tests := []struct {
		name     string
		asset    GitHubAsset
		platform Platform
	}{
		{
			name:     "https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz",
			asset:    NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			platform: NewPlatform("linux", "amd64"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			platform, err := tt.asset.Platform()
			assert.NoError(err)
			assert.Equal(tt.platform, platform)
		})
	}
}
