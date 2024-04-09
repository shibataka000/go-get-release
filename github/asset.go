package github

import "github.com/shibataka000/go-get-release/slices"

// Asset in GitHub release.
type Asset struct {
	Meta    AssetMeta
	Content AssetContent
}

// AssetMeta is meta of GitHub release asset.
type AssetMeta struct {
	DownloadURL URL
	Platform    Platform
}

// AssetMetaTemplate is template of AssetMeta.
type AssetMetaTemplate struct {
	DownloadURL URLTemplate
	Platform    Platform
}

// AssetContent is content of GitHub release asset.
type AssetContent []byte

// AssetMetaSet is set of AssetMeta.
type AssetMetaSet []AssetMeta

// NewAssetMeta return new asset meta instance.
func NewAssetMeta(url URL) AssetMeta {
	return NewAssetMetaWithPlatform(url, url.FileName().Platform())
}

func NewAssetMetaWithPlatform(url URL, platform Platform) AssetMeta {
	return AssetMeta{
		DownloadURL: url,
		Platform:    platform,
	}
}

func (s AssetMetaSet) FindByPlatform(platform Platform) (AssetMeta, error) {
	return slices.Find(s, func(asset AssetMeta) bool {
		return asset.Platform.Equal(platform)
	})
}
