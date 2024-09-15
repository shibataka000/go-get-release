package github

import (
	"os"
	"path/filepath"
)

// ExecBinary represents a executable binary in a GitHub release asset.
type ExecBinary struct {
	name string
}

// newExecBinary returns a new [ExecBinary] object.
func newExecBinary(name string) ExecBinary {
	return ExecBinary{
		name: name,
	}
}

// ExecBinaryContent represents a executable binary content in a GitHub release asset content.
type ExecBinaryContent []byte

// ExecBinaryRepository is a repository for [ExecBinary] and [ExecBinaryContent].
type ExecBinaryRepository struct{}

// NewExecBinaryRepository returns a new [ExecBinaryRepository] object.
func NewExecBinaryRepository() *ExecBinaryRepository {
	return &ExecBinaryRepository{}
}

// write [ExecBinaryContent] to file in given directory.
func (r *ExecBinaryRepository) write(meta ExecBinary, content ExecBinaryContent, dir string) error {
	path := filepath.Join(dir, meta.name)
	os.WriteFile(path, content, 0755)
	return nil
}
