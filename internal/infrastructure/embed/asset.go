package embed

import (
	"fmt"

	assetmodel "github.com/shibataka000/go-get-release/internal/domain/model/asset"
	platformmodel "github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// FindAssetMetadata find GitHub asset metadata from embed data.
func (r *Repository) FindAssetMetadata(repo repository.Repository, release release.Release, platform platformmodel.Platform) (assetmodel.Metadata, error) {
	repoNode, err := r.findRepositoryNode(repo)
	if err != nil {
		return assetmodel.Metadata{}, err
	}
	for _, asset := range repoNode.Assets {
		p := platformmodel.New(asset.OS, asset.Arch)
		if p.Equal(platform) {
			return assetmodel.NewMetadataFromTemplate(asset.DownloadURL, release)
		}
	}
	return assetmodel.Metadata{}, fmt.Errorf("no asset was found in embed data")
}

// HasAssetMetadata return true if GitHub asset metadata is defined in embed data.
func (r *Repository) HasAssetMetadata(repo repository.Repository, release release.Release, platform platformmodel.Platform) bool {
	_, err := r.FindAssetMetadata(repo, release, platform)
	return err == nil
}
