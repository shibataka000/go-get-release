package execbinary

import (
	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// Factory to make executable binary metadata instance.
type Factory struct {
	embed EmbedRepository
}

// New factory instance.
func New(embed EmbedRepository) Factory {
	return Factory{
		embed: embed,
	}
}

// FindMetadata return executable binary metadata.
// If executable binary file name is defined in embed data, return it.
// Otherwise, find metadata return repository name as executable binary file name.
func (f Factory) FindMetadata(repo repository.Repository, platform platform.Platform) (execbinary.Metadata, error) {
	if f.embed.HasExecBinaryMetadata(repo, platform) {
		return f.embed.FindExecBinaryMetadata(repo, platform)
	}
	return execbinary.NewMetadata(repo.Name(), platform), nil
}
