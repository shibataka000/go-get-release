package filesystem

import (
	"os"
	"path/filepath"

	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
)

// WriteExecBinaryContent write executable binary to filesystem.
func (r *Repository) WriteExecBinaryContent(meta execbinary.Metadata, content execbinary.Content, dir string) error {
	path := filepath.Join(dir, meta.Name())
	return os.WriteFile(path, content.Bytes(), 0755)
}
