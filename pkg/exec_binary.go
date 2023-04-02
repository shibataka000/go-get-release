package pkg

import "fmt"

// ExecBinaryMeta is metadata about executable binary.
type ExecBinaryMeta struct {
	Name FileName
}

// ExecBinaryFile is file about executable binary.
type ExecBinaryFile File

// NewExecBinaryMeta return new metadata instance about executable binary.
func NewExecBinaryMeta(name FileName) ExecBinaryMeta {
	return ExecBinaryMeta{
		Name: name,
	}
}

// NewExecBinaryMetaWithPlatform return new metadata instance about executable binary.
// If os is windows, extension ".exe" is added.
func NewExecBinaryMetaWithPlatform(baseName FileName, platform Platform) ExecBinaryMeta {
	name := baseName
	if platform.OS == "windows" {
		name = NewFileName(fmt.Sprintf("%s.exe", baseName.String()))
	}
	return NewExecBinaryMeta(name)
}

// NewExecBinaryFile return new file instance about executable binary.
func NewExecBinaryFile(name FileName, body []byte) ExecBinaryFile {
	return ExecBinaryFile{
		Name: name,
		Body: body,
	}
}
