package embed

import (
	"fmt"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
	"github.com/stretchr/testify/require"
)

func TestFindAssetMetadata(t *testing.T) {
	tests := []struct {
		repo     repository.Repository
		release  release.Release
		platform platform.Platform
		meta     asset.Metadata
	}{
		{
			repo:     repository.New("hashicorp", "terraform"),
			release:  release.New(0, "v1.3.7"),
			platform: platform.New("linux", "amd64"),
			meta:     asset.NewMetadata("https://releases.hashicorp.com/terraform/1.3.7/terraform_1.3.7_linux_amd64.zip"),
		},
	}

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			repository, err := New()
			assert.NoError(err)
			meta, err := repository.FindAssetMetadata(tt.repo, tt.release, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.meta, meta)
		})
	}
}
