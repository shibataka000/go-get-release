package embed

import (
	"fmt"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
	"github.com/stretchr/testify/require"
)

func TestFindExecBinaryMetadata(t *testing.T) {
	tests := []struct {
		repo     repository.Repository
		platform platform.Platform
		meta     execbinary.Metadata
	}{
		{
			repo:     repository.New("cli", "cli"),
			platform: platform.New("linux", "amd64"),
			meta:     execbinary.NewMetadata("gh", platform.New("linux", "amd64")),
		},
	}

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			repository, err := New()
			assert.NoError(err)
			meta, err := repository.FindExecBinaryMetadata(tt.repo, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.meta, meta)
		})
	}
}

func TestHasExecBinaryMetadata(t *testing.T) {
	tests := []struct {
		repo        repository.Repository
		platform    platform.Platform
		hasMetadata bool
	}{
		{
			repo:        repository.New("hashicorp", "terraform"),
			platform:    platform.New("linux", "amd64"),
			hasMetadata: false,
		},
		{
			repo:        repository.New("cli", "cli"),
			platform:    platform.New("linux", "amd64"),
			hasMetadata: true,
		},
	}

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			repository, err := New()
			assert.NoError(err)
			hasAssetMetadata := repository.HasExecBinaryMetadata(tt.repo, tt.platform)
			assert.Equal(tt.hasMetadata, hasAssetMetadata)
		})
	}
}
