package pkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func ReadTestFile(t *testing.T, path string) (File, error) {
	t.Helper()
	body, err := os.ReadFile(path)
	if err != nil {
		return File{}, err
	}
	base := filepath.Base(path)
	filename := NewFileName(base)
	return NewFile(filename, body), nil
}

func TestFileExtract(t *testing.T) {
	tests := []struct {
		name         string
		testFilePath string
		extracted    File
	}{
		{
			name:         "test.gz",
			testFilePath: "./testdata/test.gz",
			extracted:    NewFile("test", []byte("helloworld\n")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			file, err := ReadTestFile(t, tt.testFilePath)
			assert.NoError(err)
			extracted, err := file.Extract()
			assert.NoError(err)
			assert.Equal(tt.extracted, extracted)
		})
	}
}

func TestFileFindFile(t *testing.T) {
	tests := []struct {
		name         string
		testFilePath string
		target       FileName
		found        File
	}{
		{
			name:         "test.zip",
			testFilePath: "./testdata/test.zip",
			target:       "test",
			found:        NewFile("test", []byte("helloworld\n")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			file, err := ReadTestFile(t, tt.testFilePath)
			assert.NoError(err)
			found, err := file.FindFile(tt.target)
			assert.NoError(err)
			assert.Equal(tt.found, found)
		})
	}
}

func TestFileNameExt(t *testing.T) {
	tests := []struct {
		name     string
		filename FileName
		ext      string
	}{
		{
			name:     "test",
			filename: NewFileName("test"),
			ext:      "",
		},
		{
			name:     "test.exe",
			filename: NewFileName("test.exe"),
			ext:      ".exe",
		},
		{
			name:     "test.tar.gz",
			filename: NewFileName("test.tar.gz"),
			ext:      ".gz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.ext, tt.filename.Ext())
		})
	}
}

func TestFileNameTrimExt(t *testing.T) {
	tests := []struct {
		name     string
		filename FileName
		trimmed  FileName
	}{
		{
			name:     "test",
			filename: NewFileName("test"),
			trimmed:  NewFileName("test"),
		},
		{
			name:     "test.exe",
			filename: NewFileName("test.exe"),
			trimmed:  NewFileName("test"),
		},
		{
			name:     "test.tar.gz",
			filename: NewFileName("test.tar.gz"),
			trimmed:  NewFileName("test.tar"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.trimmed, tt.filename.TrimExt())
		})
	}
}

func TestFileNormalize(t *testing.T) {
	tests := []struct {
		name       string
		filename   FileName
		normilized FileName
	}{
		{
			name:       "test.exe",
			filename:   NewFileName("test.exe"),
			normilized: NewFileName("test.exe"),
		},
		{
			name:       "test.tar.gz",
			filename:   NewFileName("test.tar.gz"),
			normilized: NewFileName("test.tar.gz"),
		},
		{
			name:       "test.tgz",
			filename:   NewFileName("test.tgz"),
			normilized: NewFileName("test.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.normilized, tt.filename.Normalize())
		})
	}
}

func TestFileNameIsSomething(t *testing.T) {
	tests := []struct {
		name         string
		filename     FileName
		isExecBinary bool
		isCompressed bool
		isArchived   bool
		isTarBall    bool
	}{
		{
			name:         "test",
			filename:     NewFileName("test"),
			isExecBinary: true,
			isCompressed: false,
			isArchived:   false,
			isTarBall:    false,
		},
		{
			name:         "test.exe",
			filename:     NewFileName("test.exe"),
			isExecBinary: true,
			isCompressed: false,
			isArchived:   false,
			isTarBall:    false,
		},
		{
			name:         "test.gz",
			filename:     NewFileName("test.gz"),
			isExecBinary: false,
			isCompressed: true,
			isArchived:   false,
			isTarBall:    false,
		},
		{
			name:         "test.tar.gz",
			filename:     NewFileName("test.tar.gz"),
			isExecBinary: false,
			isCompressed: true,
			isArchived:   true,
			isTarBall:    true,
		},
		{
			name:         "test.tgz",
			filename:     NewFileName("test.tgz"),
			isExecBinary: false,
			isCompressed: true,
			isArchived:   true,
			isTarBall:    true,
		},
		{
			name:         "test.tar",
			filename:     NewFileName("test.tar"),
			isExecBinary: false,
			isCompressed: false,
			isArchived:   true,
			isTarBall:    true,
		},
		{
			name:         "test.zip",
			filename:     NewFileName("test.zip"),
			isExecBinary: false,
			isCompressed: true,
			isArchived:   true,
			isTarBall:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.isExecBinary, tt.filename.IsExecBinary())
			assert.Equal(tt.isArchived, tt.filename.IsArchived())
			assert.Equal(tt.isCompressed, tt.filename.IsCompressed())
		})
	}
}

func TestFileNamePlatform(t *testing.T) {
	tests := []struct {
		name     string
		filename FileName
		platform Platform
	}{
		{
			name:     "gh_2.21.0_linux_amd64.tar.gz",
			filename: NewFileName("gh_2.21.0_linux_amd64.tar.gz"),
			platform: NewPlatform("linux", "amd64"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			platform, err := tt.filename.Platform()
			assert.NoError(err)
			assert.Equal(tt.platform, platform)
		})
	}
}
