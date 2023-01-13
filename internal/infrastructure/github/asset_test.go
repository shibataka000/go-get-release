package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/stretchr/testify/require"
)

func TestListAssetMetadata(t *testing.T) {
	tests := []struct {
		owner string
		name  string
		tag   string
		meta  []asset.Metadata
	}{
		{
			owner: "cli",
			name:  "cli",
			tag:   "v2.21.1",
			meta: []asset.Metadata{
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_checksums.txt"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.deb"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.rpm"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.tar.gz"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.deb"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.rpm"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.deb"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.rpm"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.tar.gz"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.deb"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.rpm"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.tar.gz"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_macOS_amd64.tar.gz"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_386.zip"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.msi"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.zip"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_arm64.zip"),
			},
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
			meta, err := repository.ListAssetMetadata(ctx, repo, release)
			assert.NoError(err)
			assert.Equal(tt.meta, meta)
		})
	}
}
