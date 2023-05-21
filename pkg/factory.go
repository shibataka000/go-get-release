package pkg

import "fmt"

// Factory.
type Factory struct{}

// NewFactory return new factory instance.
func NewFactory() *Factory {
	return &Factory{}
}

// NewRepository return new repository instance.
func (f *Factory) NewRepository(repo GitHubRepository) Repository {
	return Repository(repo)
}

// NewRelease return new repository instance.
func (f *Factory) NewRelease(release GitHubRelease) Release {
	return NewRelease(release.Tag)
}

// NewAssetFromIndex return new asset instance from index.
func (f *Factory) NewAssetFromIndex(asset AssetInIndex, release Release) (Asset, error) {
	downloadURL, err := asset.DownloadURL.RenderWithRelease(release)
	if err != nil {
		return Asset{}, err
	}
	return NewAsset(downloadURL), nil
}

// NewAssetFromGitHub return new asset instance from GitHub.
func (f *Factory) NewAssetFromGitHub(assets []GitHubAsset, platform Platform) (Asset, error) {
	filtered := FilterGitHubAssetByPlatform(assets, platform)
	if len(filtered) == 0 {
		return Asset{}, fmt.Errorf("asset for %v was not found", platform)
	}
	return Asset(filtered[0]), nil
}

// NewExecBinaryFromIndex return executable binary instance from index.
func (f *Factory) NewExecBinaryFromIndex(execBinary ExecBinaryInIndex, platform Platform) ExecBinary {
	return f.NewExecBinaryWithPlatform(execBinary.BaseName, platform)
}

// NewExecBinaryFromGitHub return executable binary instance from GitHub.
func (f *Factory) NewExecBinaryFromGitHub(repo GitHubRepository, platform Platform) ExecBinary {
	return f.NewExecBinaryWithPlatform(NewFileName(repo.Name), platform)
}

// NewExecBinaryWithPlatform return new executable binary instance.
// If os is windows, extension ".exe" is added.
func (f *Factory) NewExecBinaryWithPlatform(baseName FileName, platform Platform) ExecBinary {
	if platform.OS == "windows" {
		return NewExecBinary(baseName.AddExt(".exe"))
	}
	return NewExecBinary(baseName)
}
