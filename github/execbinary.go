package github

import (
	"slices"

	"github.com/shibataka000/go-get-release/file"
	"github.com/shibataka000/go-get-release/platform"
	"gopkg.in/yaml.v3"
)

// ExecutableBinary represents an executable binary in a GitHub release asset.
type ExecutableBinary struct {
	BaseName file.Name `yaml:"name"`
	OS       platform.OS
}

// newExecutableBinary returns a new executable binary metadata object.
func newExecutableBinary(baseName file.Name, os platform.OS) ExecutableBinary {
	return ExecutableBinary{
		BaseName: baseName,
		OS:       os,
	}
}

// newExecutableBinaryMetaFromRepository returns a new executable binary metadata object from Repository.
func newExecutableBinaryMetaFromRepository(repo Repository, os platform.OS) ExecutableBinary {
	return newExecutableBinary(file.Name(repo.name), os)
}

// ExecutableBinaryRepository is a repository for executable binary.
type ExecutableBinaryRepository struct{}

// NewExecutableBinaryRepository returns a new ExecutableBinaryRepository object.
func NewExecutableBinaryRepository() *ExecutableBinaryRepository {
	return &ExecutableBinaryRepository{}
}

// find an executable binary metadata from built-in data.
func (r *ExecutableBinaryRepository) find(repo Repository, os platform.OS) (ExecutableBinary, error) {
	// Define structure about record in built-in data to unmarshal them.
	type Record struct {
		Repository       Repository       `yaml:"repository"`
		ExecutableBinary ExecutableBinary `yaml:"executableBinary"`
	}

	// Unmarshal built-in data.
	records := []Record{}
	err := yaml.Unmarshal(builtin, &records)
	if err != nil {
		return ExecutableBinary{}, err
	}

	// Find record in built-in data by repository.
	index := slices.IndexFunc(records, func(r Record) bool {
		return r.Repository.owner == repo.owner && r.Repository.name == repo.name && r.ExecutableBinary.BaseName != ""
	})
	if index == -1 {
		return ExecutableBinary{}, &ExecutableBinaryNotFoundError{}
	}
	bin := records[index].ExecutableBinary

	return newExecutableBinary(bin.BaseName, os), nil
}
