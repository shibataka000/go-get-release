package github

import "path"

// Asset in GitHub repository
type Asset interface {
	Name() string
	DownloadURL() string
	BinaryName() string
}

type asset struct {
	downloadURL string
	binaryName  string
}

func getAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	if isSpecialAsset(repo.owner, repo.name) {
		return getSpecialAsset(c, repo, release, goos, goarch)
	}
	return getGeneralAsset(c, repo, release, goos, goarch)
}

func (a *asset) Name() string {
	_, file := path.Split(a.downloadURL)
	return file
}

func (a *asset) DownloadURL() string {
	return a.downloadURL
}

func (a *asset) BinaryName() string {
	return a.binaryName
}
