package github

import (
	"os"
	"path/filepath"
)

// ExecBinary represents a executable binary in GitHub release asset.
type ExecBinary struct {
	name string
}

func newExecBinary(name string) ExecBinary {
	return ExecBinary{
		name: name,
	}
}

// ExecBinaryContent represents a executable binary content in GitHub release asset content.
type ExecBinaryContent []byte

type ExecBinaryPattern string

type ExecBinaryPatternList []ExecBinaryPattern

func newExecBinaryPatternList(exprs []string) ExecBinaryPatternList {
	patterns := ExecBinaryPatternList{}
	for _, expr := range exprs {
		patterns = append(patterns, ExecBinaryPattern(expr))
	}
	return patterns
}

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
