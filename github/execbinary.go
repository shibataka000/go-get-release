package github

import (
	"slices"

	"github.com/shibataka000/go-get-release/file"
	"github.com/shibataka000/go-get-release/platform"
	"gopkg.in/yaml.v3"
)

// ExecutableBinaryMeta represents an executable binary in a GitHub release asset.
type ExecutableBinaryMeta struct {
	BaseName file.Name `yaml:"name"`
	OS       platform.OS
}

// newExecutableBinaryMeta returns a new executable binary metadata object.
func newExecutableBinaryMeta(baseName file.Name, os platform.OS) ExecutableBinaryMeta {
	return ExecutableBinaryMeta{
		BaseName: baseName,
		OS:       os,
	}
}

// newExecutableBinaryMetaFromRepository returns a new executable binary metadata object from Repository.
func newExecutableBinaryMetaFromRepository(repo Repository, os platform.OS) ExecutableBinaryMeta {
	return newExecutableBinaryMeta(file.Name(repo.Name), os)
}

// ExecutableBinaryRepository is a repository for executable binary.
type ExecutableBinaryRepository struct{}

// NewExecutableBinaryRepository returns a new ExecutableBinaryRepository object.
func NewExecutableBinaryRepository() *ExecutableBinaryRepository {
	return &ExecutableBinaryRepository{}
}

// find an executable binary metadata from built-in data.
func (r *ExecutableBinaryRepository) find(repo Repository, os platform.OS) (ExecutableBinaryMeta, error) {
	// Define structure about record in built-in data to unmarshal them.
	type Record struct {
		Repository       Repository           `yaml:"repository"`
		ExecutableBinary ExecutableBinaryMeta `yaml:"executableBinary"`
	}

	// Unmarshal built-in data.
	records := []Record{}
	err := yaml.Unmarshal(builtin, &records)
	if err != nil {
		return ExecutableBinaryMeta{}, err
	}

	// Find record in built-in data by repository.
	index := slices.IndexFunc(records, func(r Record) bool {
		return r.Repository.Owner == repo.Owner && r.Repository.Name == repo.Name && r.ExecutableBinary.BaseName != ""
	})
	if index == -1 {
		return ExecutableBinaryMeta{}, &ExecutableBinaryNotFoundError{}
	}
	bin := records[index].ExecutableBinary

	return newExecutableBinaryMeta(bin.BaseName, os), nil
}
