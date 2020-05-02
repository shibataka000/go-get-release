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
