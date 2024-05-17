package github

import "github.com/shibataka000/go-get-release/file"

// ExecutableBinaryMeta is metadata of executable binary in a GitHub release asset.
type ExecutableBinaryMeta struct {
	BaseName file.Name
}

// newExecutableBinaryMeta return new ExecutableBinaryMeta object.
func newExecutableBinaryMeta(baseName file.Name) ExecutableBinaryMeta {
	return ExecutableBinaryMeta{
		BaseName: baseName,
	}
}

// newExecutableBinaryMetaFromRepository return new ExecutableBinaryMeta object from Repository.
func newExecutableBinaryMetaFromRepository(repo Repository) ExecutableBinaryMeta {
	return newExecutableBinaryMeta(file.Name(repo.Name))
}
