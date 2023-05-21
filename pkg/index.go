package pkg

import (
	"fmt"
)

// Index.
type Index struct {
	Repositories []RepositoryInIndex
}

// RepositoryInIndex is repository metadata in index.
type RepositoryInIndex struct {
	Owner      string            `yaml:"owner"`
	Name       string            `yaml:"repo"`
	Assets     []AssetInIndex    `yaml:"assets"`
	ExecBinary ExecBinaryInIndex `yaml:"execBinary"`
}

// AssetInIndex is asset metadata in index.
type AssetInIndex struct {
	DownloadURL URLTemplate `yaml:"downloadURL"`
	OS          string      `yaml:"os"`
	Arch        string      `yaml:"arch"`
}

// ExecBinaryInIndex is executable binary metadata in index.
type ExecBinaryInIndex struct {
	BaseName FileName `yaml:"name"`
}

// NewIndex return new index instance.
func NewIndex(repos []RepositoryInIndex) Index {
	return Index{
		Repositories: repos,
	}
}

// NewRepositoryInIndex return new repository metadata instance in index.
func NewRepositoryInIndex(owner string, name string, assets []AssetInIndex, execBinary ExecBinaryInIndex) RepositoryInIndex {
	return RepositoryInIndex{
		Owner:      owner,
		Name:       name,
		Assets:     assets,
		ExecBinary: execBinary,
	}
}

// NewAssetInIndex return new asset metadata instance in index.
func NewAssetInIndex(downloadURL URLTemplate, os string, arch string) AssetInIndex {
	return AssetInIndex{
		DownloadURL: downloadURL,
		OS:          os,
		Arch:        arch,
	}
}

// NewExecBinaryInIndex return new executable binary metadata instance in index.
func NewExecBinaryInIndex(baseName FileName) ExecBinaryInIndex {
	return ExecBinaryInIndex{
		BaseName: baseName,
	}
}

// FindRepository find repository metadata from index.
func (i Index) FindRepository(repo Repository) (RepositoryInIndex, error) {
	for _, r := range i.Repositories {
		if r.Equals(repo) {
			return r, nil
		}
	}
	return RepositoryInIndex{}, fmt.Errorf("repository %v was not found in index", repo)
}

// FindAsset find asset metadata from index.
func (i Index) FindAsset(repo Repository, platform Platform) (AssetInIndex, error) {
	r, err := i.FindRepository(repo)
	if err != nil {
		return AssetInIndex{}, err
	}
	return r.FindAsset(platform)
}

// HasAsset return true if index has asset metadata about specified repository and platform.
func (i Index) HasAsset(repo Repository, platform Platform) bool {
	_, err := i.FindAsset(repo, platform)
	return err == nil
}

// FindExecBianry find executable binary metadata from index.
func (i Index) FindExecBinary(repo Repository) (ExecBinaryInIndex, error) {
	r, err := i.FindRepository(repo)
	if err != nil {
		return ExecBinaryInIndex{}, err
	}
	return r.ExecBinary, nil
}

// HasExecBinary return true if index has executable binary metadata about speficied repository.
func (i Index) HasExecBinary(repo Repository) bool {
	execBinary, err := i.FindExecBinary(repo)
	if err != nil {
		return false
	}
	return !execBinary.IsEmpty()
}

// Equals return true if RepositoryInIndex and Repository specify same repository.
func (r RepositoryInIndex) Equals(repo Repository) bool {
	return r.Owner == repo.Owner && r.Name == repo.Name
}

// FindAsset find asset metadata from index.
func (r RepositoryInIndex) FindAsset(platform Platform) (AssetInIndex, error) {
	for _, asset := range r.Assets {
		p := NewPlatform(asset.OS, asset.Arch)
		if platform.Equals(p) {
			return asset, nil
		}
	}
	return AssetInIndex{}, fmt.Errorf("asset for platform %v was not found in index", platform)
}

// IsEmpty return true if executable binary metadata is not defined.
func (b ExecBinaryInIndex) IsEmpty() bool {
	return b.BaseName == ""
}
