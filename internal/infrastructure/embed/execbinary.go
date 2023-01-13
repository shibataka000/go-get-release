package embed

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// FindExecBinaryMetadata find executable binary metadata from mebed data.
func (r *Repository) FindExecBinaryMetadata(repo repository.Repository, platform platform.Platform) (execbinary.Metadata, error) {
	repoNode, err := r.findRepositoryNode(repo)
	if err != nil {
		return execbinary.Metadata{}, err
	}
	if repoNode.ExecBinary.Name == "" {
		return execbinary.Metadata{}, fmt.Errorf("no executable binary name was found in embed data")
	}
	return execbinary.NewMetadata(repoNode.ExecBinary.Name, platform), nil
}

// HasExecBinaryMetadata return true if executable binary metadata is defined in embed data.
func (r *Repository) HasExecBinaryMetadata(repo repository.Repository, platform platform.Platform) bool {
	_, err := r.FindExecBinaryMetadata(repo, platform)
	return err == nil
}
