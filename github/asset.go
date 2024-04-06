package github

// Asset in GitHub release.
type Asset struct {
	Meta    AssetMeta
	Content AssetContent
}

// AssetMeta is meta of GitHub release asset.
type AssetMeta struct {
	DownloadURL URL
}

// AssetMetaTemplate is template of AssetMeta.
type AssetMetaTemplate struct {
	DownloadURL URLTemplate
}

// AssetContent is content of GitHub release asset.
type AssetContent []byte

// NewAssetMeta return new asset meta instance.
func NewAssetMeta(url URL) AssetMeta {
	return AssetMeta{
		DownloadURL: url,
	}
}
