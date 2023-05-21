package pkg

// GitHubRepository is repository in GitHub.
type GitHubRepository struct {
	Owner string
	Name  string
}

// GitHubRelease is release in GitHub.
type GitHubRelease struct {
	ID  int64
	Tag string
}

// GitHubAsset is release asset in GitHub.
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

// NewGitHubRelease return new GitHub release instance.
func NewGitHubRelease(id int64, tag string) GitHubRelease {
	return GitHubRelease{
		ID:  id,
		Tag: tag,
	}
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

// Platform return platform guessed by asset file name.
func (a GitHubAsset) Platform() (Platform, error) {
	filename := a.DownloadURL.FileName()
	return filename.Platform()
}

// FilterGitHubAssetByPlatform filter assets which has executable binary for specified platform.
func FilterGitHubAssetByPlatform(assets []GitHubAsset, platform Platform) []GitHubAsset {
	result := []GitHubAsset{}
	for _, asset := range assets {
		p, err := asset.Platform()
		if err != nil {
			continue
		}
		if asset.HasExecBinary() && platform.Equals(p) {
			result = append(result, asset)
		}
	}
	return result
}
