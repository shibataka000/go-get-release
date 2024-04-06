package github

type AssetMeta struct {
	DownloadURL URL
}

type AssetMetaTemplate struct {
	DownloadURL URLTemplate
}

type AssetContent []byte

type Asset struct {
	Meta    AssetMeta
	Content AssetContent
}

func NewAssetMeta(url URL) AssetMeta {
	return AssetMeta{
		DownloadURL: url,
	}
}

func NewAssetMetaTemplate(url URLTemplate) AssetMetaTemplate {
	return AssetMetaTemplate{
		DownloadURL: url,
	}
}

func NewAsset(meta AssetMeta, content AssetContent) Asset {
	return Asset{
		Meta:    meta,
		Content: content,
	}
}
