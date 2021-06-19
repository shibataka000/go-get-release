package pkg

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/shibataka000/go-get-release/internal/github"
)

// FindInput is input to Find function
type FindInput struct {
	Name        string
	GithubToken string
	Goos        string
	Goarch      string
}

// FindOutput is output by Find function. This have information about golang release binary
type FindOutput struct {
	Owner       string
	Repo        string
	Tag         string
	Asset       string
	DownloadURL string
	BinaryName  string
	IsArchived  bool
}

// Find golang release binary and return its information
func Find(input *FindInput) (*FindOutput, error) {
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
		repo, err = client.Repository(owner, repoStr)
	} else {
		repo, err = client.FindRepository(repoStr)
	}
	if err != nil {
		return nil, err
	}

	var release github.Release
	if tag != "" {
		release, err = repo.Release(tag)
	} else {
		release, err = repo.LatestRelease()
	}
	if err != nil {
		return nil, err
	}

	asset, err := release.Asset(input.Goos, input.Goarch)
	if err != nil {
		return nil, err
	}

	return &FindOutput{
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
	return "", "", "", fmt.Errorf("parsing package name failed: %s\npackage name should be \"owner/repo=tag\" format", name)
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
