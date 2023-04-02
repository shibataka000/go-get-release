package pkg

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

// GitHubRepository is repository in GitHub.
type GitHubRepository struct {
	Owner string
	Name  string
}

// GitHubRelease is release in GitHub repository.
type GitHubRelease struct {
	ID  int64
	Tag string
}

// GitHubAsset is asset in GitHub release.
type GitHubAsset struct {
	DownloadURL URL
}

// NewGitHubRepository return new GitHub repository instance.
func NewGitHubRepository(owner string, name string) GitHubRepository {
	return GitHubRepository{
		Owner: owner,
		Name:  name,
	}
}

// FullName return GitHub repository full name.
func (r GitHubRepository) FullName() string {
	return fmt.Sprintf("%s/%s", r.Owner, r.Name)
}

// NewGitHubRelease return new GitHub release instance.
func NewGitHubRelease(id int64, tag string) GitHubRelease {
	return GitHubRelease{
		ID:  id,
		Tag: tag,
	}
}

// SemVer return semver formatted release tag.
// For example, if release tag is "v1.2.3", this return "1.2.3".
func (r GitHubRelease) SemVer() (string, error) {
	if !semver.IsValid(r.Tag) && !semver.IsValid(fmt.Sprintf("v%s", r.Tag)) {
		return "", fmt.Errorf("%s is not valid semver", r.Tag)
	}
	return strings.TrimLeft(r.Tag, "v"), nil
}

// NewGitHubAsset return new GitHub release asset instance.
func NewGitHubAsset(downloadURL URL) GitHubAsset {
	return GitHubAsset{
		DownloadURL: downloadURL,
	}
}

// HasExecBinary return true if asset has exec binary.
func (a GitHubAsset) HasExecBinary() bool {
	filename := a.DownloadURL.FileName()
	return filename.IsExecBinary() || filename.IsArchived() || filename.IsCompressed()
}

// Platform return platform guessed by asset name.
func (a GitHubAsset) Platform() (Platform, error) {
	filename := a.DownloadURL.FileName()
	return filename.Platform()
}
