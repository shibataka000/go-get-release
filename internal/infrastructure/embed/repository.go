package embed

import (
	"gopkg.in/yaml.v3"
)

// Repository for embed data.
type Repository struct {
	repos []RepositoryNode
}

// New repository for embed data.
func New() (*Repository, error) {
	repos := []RepositoryNode{}
	err := yaml.Unmarshal(embedAssetData, &repos)
	return &Repository{
		repos: repos,
	}, err
}
