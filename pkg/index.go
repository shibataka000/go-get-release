package pkg

import (
	"fmt"
)

// Index about GitHub release asset and executable binary.
type Index struct {
	Repositories []RepositoryInIndex
}

// RepositoryInIndex is metadta about GitHub repository in index.
type RepositoryInIndex struct {
	Owner      string            `yaml:"owner"`
	Name       string            `yaml:"repo"`
	Assets     []AssetInIndex    `yaml:"assets"`
	ExecBinary ExecBinaryInIndex `yaml:"execBinary"`
}

// AssetInIndex is metadata about GitHub release asset in index.
type AssetInIndex struct {
	DownloadURL URLTemplate `yaml:"downloadURL"`
	OS          string      `yaml:"os"`
	Arch        string      `yaml:"arch"`
}

// ExecBinaryInIndex is metadata about executable binary in index.
type ExecBinaryInIndex struct {
	BaseName FileName `yaml:"name"`
}

// NewIndex return new index instance.
func NewIndex(repos []RepositoryInIndex) Index {
	return Index{
		Repositories: repos,
	}
}

// FindRepository find repository from index.
func (i Index) FindRepository(repo GitHubRepository) (RepositoryInIndex, error) {
	for _, r := range i.Repositories {
		if r.Equals(repo) {
			return r, nil
		}
	}
	return RepositoryInIndex{}, fmt.Errorf("repository %v was not found in index", repo)
}

// FindAsset find asset from index.
func (i Index) FindAsset(repo GitHubRepository, platform Platform) (AssetInIndex, error) {
	r, err := i.FindRepository(repo)
	if err != nil {
		return AssetInIndex{}, err
	}
	return r.FindAsset(platform)
}

// HasAsset return true if metadata about GitHub release asset about specified repository and platform exists in index.
func (i Index) HasAsset(repo GitHubRepository, platform Platform) bool {
	_, err := i.FindAsset(repo, platform)
	return err == nil
}

// FindExecBianry return metadata about executable binary from index.
func (i Index) FindExecBinary(repo GitHubRepository) (ExecBinaryInIndex, error) {
	r, err := i.FindRepository(repo)
	if err != nil {
		return ExecBinaryInIndex{}, err
	}
	return r.ExecBinary, nil
}

// HasExecBinary return true if metadata about executable binary exists in index.
func (i Index) HasExecBinary(repo GitHubRepository) bool {
	execBinary, err := i.FindExecBinary(repo)
	if err != nil {
		return false
	}
	return !execBinary.IsEmpty()
}

// NewRepositoryInIndex return new metadata instance about GitHub repository in index.
func NewRepositoryInIndex(owner string, name string, assets []AssetInIndex, execBinary ExecBinaryInIndex) RepositoryInIndex {
	return RepositoryInIndex{
		Owner:      owner,
		Name:       name,
		Assets:     assets,
		ExecBinary: execBinary,
	}
}

// Equals return true if RepositoryInIndex and GitHubRepository specify same repository.
func (r RepositoryInIndex) Equals(repo GitHubRepository) bool {
	return r.Owner == repo.Owner && r.Name == repo.Name
}

// FindAsset find asset.
func (r RepositoryInIndex) FindAsset(platform Platform) (AssetInIndex, error) {
	for _, asset := range r.Assets {
		p := NewPlatform(asset.OS, asset.Arch)
		if platform.Equals(p) {
			return asset, nil
		}
	}
	return AssetInIndex{}, fmt.Errorf("asset for platform %v was not found in index", platform)
}

// NewAssetInIndex return new metadata instance about GitHub release asset in index.
func NewAssetInIndex(downloadURL URLTemplate, os string, arch string) AssetInIndex {
	return AssetInIndex{
		DownloadURL: downloadURL,
		OS:          os,
		Arch:        arch,
	}
}

// NewExecBinaryInIndex return new metadata instance about executable binary in index.
func NewExecBinaryInIndex(baseName FileName) ExecBinaryInIndex {
	return ExecBinaryInIndex{
		BaseName: baseName,
	}
}

// IsEmpty return true if metadata about executable binary is not defined in index.
func (b ExecBinaryInIndex) IsEmpty() bool {
	return b.BaseName == ""
}
