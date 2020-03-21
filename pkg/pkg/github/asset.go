package github

// Asset in GitHub repository
type Asset interface {
	Name() string
	DownloadURL() string
	BinaryName() string
}

type asset struct {
	name        string
	downloadURL string
	binaryName  string
}

func (a *asset) Name() string {
	return a.name
}

func (a *asset) DownloadURL() string {
	return a.downloadURL
}

func (a *asset) BinaryName() string {
	return a.binaryName
}
