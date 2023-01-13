package asset

import (
	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
)

// FindMetadata is interface for specification to find GitHub asset metadata.
type FindMetadata interface {
	IsSatisfied(asset.Metadata) (bool, error)
}

// FindMetadataByPlatform is specification to find GitHub asset metadata by platform guessed by asset name.
type FindMetadataByPlatform struct {
	platform platform.Platform
}

// NewFindMetadataByPlatform return FindMetadataByPlatform instance.
func NewFindMetadataByPlatform(platform platform.Platform) FindMetadataByPlatform {
	return FindMetadataByPlatform{
		platform: platform,
	}
}

// IsSatisfied return true if asset has executable binary and its platform is same to specification.
func (s FindMetadataByPlatform) IsSatisfied(meta asset.Metadata) (bool, error) {
	platform, err := meta.Platform()
	if err != nil {
		// return not error but false if guessing platform by asset name fail.
		return false, nil
	}
	return (meta.IsExecBinary() || meta.IsArchived() || meta.IsCompressed()) && s.platform.Equal(platform), nil
}
