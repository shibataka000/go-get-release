package pkg

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/shibataka000/go-get-release/pkg/pkg/github"
)

// InfoInput is input to Info function
type InfoInput struct {
	Name        string
	GithubToken string
	Goos        string
	Goarch      string
}

// InfoOutput is output by Info function. This have information about golang release binary
type InfoOutput struct {
	Owner       string
	Repo        string
	Tag         string
	Asset       string
	DownloadURL string
	BinaryName  string
	IsArchived  bool
}

// Info return information about golang release binary
func Info(input *InfoInput) (*InfoOutput, error) {
	owner, repoStr, tag, err := parse(input.Name)
	if err != nil {
		return nil, err
	}

	client, err := github.NewClient(input.GithubToken)
	if err != nil {
		return nil, err
	}

	var repo github.Repository
	if owner != "" {
		repo, err = client.GetRepository(owner, repoStr)
	} else {
		repo, err = client.FindRepository(repoStr)
	}
	if err != nil {
		return nil, err
	}

	var release github.Release
	if tag != "" {
		release, err = repo.GetRelease(tag)
	} else {
		release, err = repo.GetLatestRelease()
	}
	if err != nil {
		return nil, err
	}

	asset, err := release.GetAsset(input.Goos, input.Goarch)
	if err != nil {
		return nil, err
	}

	return &InfoOutput{
		Owner:       repo.Owner(),
		Repo:        repo.Name(),
		Tag:         release.Tag(),
		Asset:       asset.Name(),
		DownloadURL: asset.DownloadURL(),
		BinaryName:  asset.BinaryName(),
		IsArchived:  isArchived(asset.Name()),
	}, nil
}

func parse(name string) (string, string, string, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^=]+))?`)
	if re.MatchString(name) {
		match := re.FindStringSubmatch(name)
		return match[2], match[3], match[5], nil
	}
	return "", "", "", fmt.Errorf("Parsing package name failed: %s\npackage name should be \"owner/repo=tag\" format", name)
}

func isArchived(asset string) bool {
	archivedExt := []string{".tar", ".gz", ".tgz", ".bz2", ".tbz", ".Z", ".zip", ".bz2", ".lzh", ".7z", ".gz", ".rar", ".cab", ".afz"}
	for _, ext := range archivedExt {
		if filepath.Ext(asset) == ext {
			return true
		}
	}
	return false
}
