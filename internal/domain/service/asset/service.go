package asset

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	specification "github.com/shibataka000/go-get-release/internal/domain/specification/asset"
)

// Service is domain service for GitHub asset.
type Service struct {
}

// New return domain service instance for GitHub asset.
func New() Service {
	return Service{}
}

// FindMetadata find asset metadata which satisfy specification.
func (s Service) FindMetadata(meta []asset.Metadata, spec specification.FindMetadata) (asset.Metadata, error) {
	filtered, err := s.filterMetadata(meta, spec)
	if err != nil {
		return asset.Metadata{}, err
	}
	switch len(filtered) {
	case 0:
		return asset.Metadata{}, fmt.Errorf("no asset was found")
	case 1:
		return filtered[0], nil
	case 2:
		if filtered[0].IsExecBinary() && (filtered[1].IsArchived() || filtered[1].IsCompressed()) {
			return filtered[0], nil
		} else if filtered[1].IsExecBinary() && (filtered[0].IsArchived() || filtered[0].IsCompressed()) {
			return filtered[1], nil
		}
		fallthrough
	default:
		names := []string{}
		for _, meta := range filtered {
			names = append(names, meta.Name())
		}
		return asset.Metadata{}, fmt.Errorf("%d assets were found: %v", len(names), names)
	}
}

// filterMetadata return list of asset metadata which satisfy specification.
func (s Service) filterMetadata(meta []asset.Metadata, spec specification.FindMetadata) ([]asset.Metadata, error) {
	filtered := []asset.Metadata{}
	for _, m := range meta {
		satisfied, err := spec.IsSatisfied(m)
		if err != nil {
			return nil, err
		}
		if satisfied {
			filtered = append(filtered, m)
		}
	}
	return filtered, nil
}
