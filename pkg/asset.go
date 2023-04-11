package pkg

import "fmt"

// AssetMeta is metadata about GitHub release asset.
type AssetMeta struct {
	DownloadURL URL
}

// AssetFile is file about GitHub release asset.
type AssetFile File

// NewAssetMeta return new metadata instance about GitHub release asset.
func NewAssetMeta(downloadURL URL) AssetMeta {
	return AssetMeta{
		DownloadURL: downloadURL,
	}
}

// NewAssetMetaFromIndex return new metadata instance about GitHub release asset from AssetInIndex and GitHubRelease.
func NewAssetMetaFromIndex(asset AssetInIndex, release GitHubRelease) (AssetMeta, error) {
	url, err := asset.DownloadURL.RenderWithRelease(release)
	if err != nil {
		return AssetMeta{}, err
	}
	return NewAssetMeta(url), nil
}

// NewAssetFile return new file instance about GitHub release asset.
func NewAssetFile(name FileName, body []byte) AssetFile {
	return AssetFile{
		Name: name,
		Body: body,
	}
}

// ExecBinary return executable binary file in GitHub release asset.
func (f AssetFile) ExecBinary(execBinary FileName) (ExecBinaryFile, error) {
	file := File(f)
	var err error

	if file.Name.IsCompressed() && (!file.Name.IsArchived() || file.Name.IsTarBall()) {
		file, err = file.Extract()
		if err != nil {
			return ExecBinaryFile{}, err
		}
	}

	if file.Name.IsArchived() {
		file, err = file.FindFile(execBinary)
		if err != nil {
			return ExecBinaryFile{}, err
		}
	}

	if !file.Name.IsExecBinary() {
		return ExecBinaryFile{}, fmt.Errorf("%s is not executable binary", file.Name)
	}

	return NewExecBinaryFile(execBinary, file.Body), nil
}
