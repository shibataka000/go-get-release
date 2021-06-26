package github

// Asset in GitHub repository
type Asset interface {
	Name() string
	DownloadURL() string
}

type asset struct {
	name        string
	downloadURL string
}

func (a *asset) Name() string {
	return a.name
}

func (a *asset) DownloadURL() string {
	return a.downloadURL
}
