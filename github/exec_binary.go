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

// ExecBinaryTemplate represents a template of [ExecBinary].
type ExecBinaryTemplate ExecBinary

// newExecBinaryTemplate returns a new [ExecBinaryTemplate] object.
// `name` must be template used by [regexp.Regexp.Expand].
func newExecBinaryTemplate(name string) ExecBinaryTemplate {
	return ExecBinaryTemplate(newExecBinary(name))
}

// execute applies a [ExecBinaryTemplate] to the [Asset] and [AssetPattern] object, and returns a new [ExecBinary] object.
func (t ExecBinaryTemplate) execute(asset Asset, pattern AssetPattern) ExecBinary {
	name := pattern.expand(string(t.name), asset.name)
	return newExecBinary(name)
}

// ExecBinaryTemplateList is a list of [ExecBinaryTemplate].
type ExecBinaryTemplateList []ExecBinaryTemplate

// newExecBinaryTemplateList returns a new [ExecBinaryTemplateList] object.
func newExecBinaryTemplateList(names []string) ExecBinaryTemplateList {
	tmpls := ExecBinaryTemplateList{}
	for _, name := range names {
		tmpls = append(tmpls, newExecBinaryTemplate(name))
	}
	return tmpls
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
