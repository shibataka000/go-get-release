package github

import (
	"fmt"
	"slices"

	"github.com/shibataka000/go-get-release/file"
	"github.com/shibataka000/go-get-release/platform"
	"gopkg.in/yaml.v3"
)

// ExecutableBinaryMeta is metadata of executable binary in a GitHub release asset.
type ExecutableBinaryMeta struct {
	BaseName file.Name   `yaml:"name"`
	OS       platform.OS `yaml:"os"`
}

// newExecutableBinaryMeta return new ExecutableBinaryMeta object.
func newExecutableBinaryMeta(baseName file.Name, os platform.OS) ExecutableBinaryMeta {
	return ExecutableBinaryMeta{
		BaseName: baseName,
		OS:       os,
	}
}

// newExecutableBinaryMetaFromRepository return new ExecutableBinaryMeta object from Repository.
func newExecutableBinaryMetaFromRepository(repo Repository, os platform.OS) ExecutableBinaryMeta {
	return newExecutableBinaryMeta(file.Name(repo.Name), os)
}

// Name return file name.
func (b ExecutableBinaryMeta) Name() file.Name {
	if b.OS == "windows" {
		return file.Name(fmt.Sprintf("%s.exe", b.BaseName))
	}
	return b.BaseName
}

// ExecutableBinaryRepository is repository for executable binary.
type ExecutableBinaryRepository struct{}

// NewExecutableBinaryRepository return new ExecutableBinaryRepository object.
func NewExecutableBinaryRepository() *ExecutableBinaryRepository {
	return &ExecutableBinaryRepository{}
}

// find metadata of executable binary from built-in data.
func (r *ExecutableBinaryRepository) find(repo Repository, os platform.OS) (ExecutableBinaryMeta, error) {
	type Record struct {
		Repository       Repository           `yaml:"repository"`
		ExecutableBinary ExecutableBinaryMeta `yaml:"executableBinary"`
	}
	records := []Record{}
	err := yaml.Unmarshal(builtin, &records)
	if err != nil {
		return ExecutableBinaryMeta{}, err
	}
	index := slices.IndexFunc(records, func(r Record) bool {
		return r.Repository.Owner == repo.Owner && r.Repository.Name == repo.Name && r.ExecutableBinary.BaseName != ""
	})
	if index == -1 {
		return ExecutableBinaryMeta{}, &ExecutableBinaryNotFoundError{}
	}
	execBinary := records[index].ExecutableBinary

	// Set OS.
	execBinary.OS = os

	return execBinary, nil
}
