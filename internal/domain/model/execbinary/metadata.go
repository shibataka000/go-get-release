package execbinary

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
)

// Metadata about executable binary in GitHub asset.
type Metadata struct {
	name     string
	platform platform.Platform
}

// NewMetadata return metadata instance about executable binary.
func NewMetadata(name string, platform platform.Platform) Metadata {
	return Metadata{
		name:     name,
		platform: platform,
	}
}

// Name return executable binary name.
func (m Metadata) Name() string {
	if m.platform.OS() == "windows" {
		return fmt.Sprintf("%s.exe", m.name)
	}
	return m.name
}
