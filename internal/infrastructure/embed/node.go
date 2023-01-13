package embed

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// RepositoryNode is node about GitHub repository in embed data.
type RepositoryNode struct {
	Owner      string         `yaml:"owner"`
	Repo       string         `yaml:"repo"`
	Assets     []AssetNode    `yaml:"assets"`
	ExecBinary ExecBinaryNode `yaml:"execBinary"`
}

// AssetNode is node about GitHub asset in embed data.
type AssetNode struct {
	OS          string `yaml:"os"`
	Arch        string `yaml:"arch"`
	DownloadURL string `yaml:"downloadURL"`
}

// ExecBinaryNode is node about executable binary in embed data.
type ExecBinaryNode struct {
	Name string `yaml:"name"`
}

// NewRepositoryNode return RepositoryNode instance.
func NewRepositoryNode(owner string, repo string, assets []AssetNode, execBinary ExecBinaryNode) RepositoryNode {
	return RepositoryNode{
		Owner:      owner,
		Repo:       repo,
		Assets:     assets,
		ExecBinary: execBinary,
	}
}

// NewAssetNode return AssetNode instance.
func NewAssetNode(os string, arch string, downloadURL string) AssetNode {
	return AssetNode{
		OS:          os,
		Arch:        arch,
		DownloadURL: downloadURL,
	}
}

// NewExecBinaryNode return ExecBinaryNode instance.
func NewExecBinaryNode(name string) ExecBinaryNode {
	return ExecBinaryNode{
		Name: name,
	}
}

// findRepositoryNode return node about GitHub repository from embed data.
func (r *Repository) findRepositoryNode(repo repository.Repository) (RepositoryNode, error) {
	for _, repoNode := range r.repos {
		if repoNode.Owner == repo.Owner() && repoNode.Repo == repo.Name() {
			return repoNode, nil
		}
	}
	return RepositoryNode{}, fmt.Errorf("no repository was found in embed data")
}
