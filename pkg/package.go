package pkg

import "fmt"

// Package.
type Package struct {
	Repository GitHubRepository
	Release    GitHubRelease
	Asset      AssetMeta
	ExecBinary ExecBinaryMeta
}

// New package instance.
func New(repo GitHubRepository, release GitHubRelease, asset AssetMeta, execBinary ExecBinaryMeta) Package {
	return Package{
		Repository: repo,
		Release:    release,
		Asset:      asset,
		ExecBinary: execBinary,
	}
}

// StringToPrompt return string to prompt.
func (p Package) StringToPrompt() string {
	return fmt.Sprintf("Repo:\t%s/%s\nTag:\t%s\nAsset:\t%s\nBinary:\t%s", p.Repository.Owner, p.Repository.Name, p.Release.Tag, p.Asset.DownloadURL.FileName().String(), p.ExecBinary.Name)
}
