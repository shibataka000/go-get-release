package pkg

import "fmt"

// Factory.
type Factory struct {
	index Index
}

// NewFactory return new factory instance.
func NewFactory(index Index) *Factory {
	return &Factory{
		index: index,
	}
}

// NewAssetMeta return new metadata instance about GitHub release asset.
// If metadata is defined in index, it is returned. Otherwise metadata got from GitHub is returned.
func (f *Factory) NewAssetMeta(repo GitHubRepository, release GitHubRelease, assets []GitHubAsset, platform Platform) (AssetMeta, error) {
	if f.index.HasAsset(repo, platform) {
		asset, err := f.index.FindAsset(repo, platform)
		if err != nil {
			return AssetMeta{}, err
		}
		return NewAssetMetaFromIndex(asset, release)
	}

	asset, err := findAssetByPlatform(assets, platform)
	if err != nil {
		return AssetMeta{}, err
	}
	return AssetMeta(asset), err
}

// NewExecBinaryMeta return new metadata instance about executable binary.
// If metadata is defined in index, it is returned. Otherwise metadata got from GitHub is returned.
// In latter case, binary name is same to GitHub repository name.
func (f *Factory) NewExecBinaryMeta(repo GitHubRepository, platform Platform) (ExecBinaryMeta, error) {
	if f.index.HasExecBinary(repo) {
		execBinary, err := f.index.FindExecBinary(repo)
		if err != nil {
			return ExecBinaryMeta{}, err
		}
		return NewExecBinaryMetaWithPlatform(execBinary.BaseName, platform), nil
	}

	return NewExecBinaryMetaWithPlatform(NewFileName(repo.Name), platform), nil
}

// findAssetByPlatform find asset which has executable binary and whose platform is same to one in arguments from assets in arguments.
func findAssetByPlatform(assets []GitHubAsset, platform Platform) (GitHubAsset, error) {
	for _, asset := range assets {
		p, err := asset.Platform()
		if err != nil {
			continue
		}
		if asset.HasExecBinary() && platform.Equals(p) {
			return asset, nil
		}
	}
	return GitHubAsset{}, fmt.Errorf("asset for %v was not found in %v", platform, assets)
}
