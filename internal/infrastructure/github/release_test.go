package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/stretchr/testify/require"
)

func TestLatestRelease(t *testing.T) {
	tests := []struct {
		owner   string
		name    string
		release release.Release
	}{
		{
			owner:   "shibataka000",
			name:    "go-get-release-test",
			release: release.New(23712476, "v0.0.2"),
		},
	}

	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	repository := New(ctx, token)

	for _, tt := range tests {
		name := fmt.Sprintf("%s/%s", tt.owner, tt.name)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			repo, err := repository.FindRepository(ctx, tt.owner, tt.name)
			assert.NoError(err)
			release, err := repository.LatestRelease(ctx, repo)
			assert.NoError(err)
			assert.Equal(tt.release, release)
		})
	}
}

func TestFindReleaseByTag(t *testing.T) {
	tests := []struct {
		owner   string
		name    string
		tag     string
		release release.Release
	}{
		{
			owner:   "shibataka000",
			name:    "go-get-release-test",
			tag:     "v0.0.1",
			release: release.New(23712441, "v0.0.1"),
		},
	}

	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	repository := New(ctx, token)

	for _, tt := range tests {
		name := fmt.Sprintf("%s/%s", tt.owner, tt.name)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			repo, err := repository.FindRepository(ctx, tt.owner, tt.name)
			assert.NoError(err)
			release, err := repository.FindReleaseByTag(ctx, repo, tt.tag)
			assert.NoError(err)
			assert.Equal(tt.release, release)
		})
	}
}
