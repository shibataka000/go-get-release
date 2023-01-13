package application

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// Metadata about GitHub release asset.
type Metadata struct {
	repo       repository.Repository
	release    release.Release
	asset      asset.Metadata
	execBinary execbinary.Metadata
}

// NewMetadata return metadata instance about GitHub release asset.
func NewMetadata(repo repository.Repository, release release.Release, asset asset.Metadata, execBinary execbinary.Metadata) Metadata {
	return Metadata{
		repo:       repo,
		release:    release,
		asset:      asset,
		execBinary: execBinary,
	}
}

// Asset which is related to this metadata instance.
func (m Metadata) Asset() asset.Metadata {
	return m.asset
}

// ExecBinary which is related to this metadata instance.
func (m Metadata) ExecBinary() execbinary.Metadata {
	return m.execBinary
}

// Format return string to prompt metadata.
func (m Metadata) Format() string {
	return fmt.Sprintf("Repo:\t%s/%s\nTag:\t%s\nAsset:\t%s\nBinary:\t%s", m.repo.Owner(), m.repo.Name(), m.release.Tag(), m.asset.Name(), m.execBinary.Name())
}
