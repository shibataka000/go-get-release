package execbinary

import (
	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// EmbedRepository is interface for embed data.
type EmbedRepository interface {
	FindExecBinaryMetadata(repo repository.Repository, platform platform.Platform) (execbinary.Metadata, error)
	HasExecBinaryMetadata(repo repository.Repository, platform platform.Platform) bool
}
