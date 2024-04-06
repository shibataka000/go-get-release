package github

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func ReadTestFile(t *testing.T, path string) (File, error) {
	t.Helper()
	name := filepath.Base(path)
	content, err := os.ReadFile(path)
	if err != nil {
		return File{}, err
	}
	return NewFile(FileName(name), content), nil
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
		{
			name:         "test.xz",
			testFilePath: "./testdata/test.xz",
			extracted:    NewFile("test", []byte("helloworld\n")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			file, err := ReadTestFile(t, tt.testFilePath)
			require.NoError(err)
			extracted, err := file.Extract()
			require.NoError(err)
			require.Equal(tt.extracted, extracted)
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
		{
			name:         "test.tar",
			testFilePath: "./testdata/test.tar",
			target:       "test",
			found:        NewFile("test", []byte("helloworld\n")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			file, err := ReadTestFile(t, tt.testFilePath)
			require.NoError(err)
			found, err := file.FindFile(tt.target)
			require.NoError(err)
			require.Equal(tt.found, found)
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
			filename: "test",
			ext:      "",
		},
		{
			name:     "test.exe",
			filename: "test.exe",
			ext:      ".exe",
		},
		{
			name:     "test.tar.gz",
			filename: "test.tar.gz",
			ext:      ".gz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.ext, tt.filename.Ext())
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
			filename: "test",
			trimmed:  "test",
		},
		{
			name:     "test.exe",
			filename: "test.exe",
			trimmed:  "test",
		},
		{
			name:     "test.tar.gz",
			filename: "test.tar.gz",
			trimmed:  "test.tar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.trimmed, tt.filename.TrimExt())
		})
	}
}

func TestFileNameAddExt(t *testing.T) {
	tests := []struct {
		name     string
		filename FileName
		ext      string
		added    FileName
	}{
		{
			name:     ".exe",
			filename: "test",
			ext:      ".exe",
			added:    "test.exe",
		},
		{
			name:     "exe",
			filename: "test",
			ext:      "exe",
			added:    "test.exe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.added, tt.filename.AddExt(tt.ext))
		})
	}

}

func TestFileNameNormalize(t *testing.T) {
	tests := []struct {
		name       string
		filename   FileName
		normilized FileName
	}{
		{
			name:       "test.exe",
			filename:   "test.exe",
			normilized: "test.exe",
		},
		{
			name:       "test.tar.gz",
			filename:   "test.tar.gz",
			normilized: "test.tar.gz",
		},
		{
			name:       "test.tgz",
			filename:   "test.tgz",
			normilized: "test.tar.gz",
		},
		{
			name:       "test.txz",
			filename:   "test.txz",
			normilized: "test.tar.xz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.normilized, tt.filename.Normalize())
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
			filename:     "test",
			isExecBinary: true,
			isCompressed: false,
			isArchived:   false,
			isTarBall:    false,
		},
		{
			name:         "test.exe",
			filename:     "test.exe",
			isExecBinary: true,
			isCompressed: false,
			isArchived:   false,
			isTarBall:    false,
		},
		{
			name:         "test.gz",
			filename:     "test.gz",
			isExecBinary: false,
			isCompressed: true,
			isArchived:   false,
			isTarBall:    false,
		},
		{
			name:         "test.tar.gz",
			filename:     "test.tar.gz",
			isExecBinary: false,
			isCompressed: true,
			isArchived:   true,
			isTarBall:    true,
		},
		{
			name:         "test.tgz",
			filename:     "test.tgz",
			isExecBinary: false,
			isCompressed: true,
			isArchived:   true,
			isTarBall:    true,
		},
		{
			name:         "test.tar",
			filename:     "test.tar",
			isExecBinary: false,
			isCompressed: false,
			isArchived:   true,
			isTarBall:    true,
		},
		{
			name:         "test.zip",
			filename:     "test.zip",
			isExecBinary: false,
			isCompressed: true,
			isArchived:   true,
			isTarBall:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.isExecBinary, tt.filename.IsExecBinary())
			require.Equal(tt.isArchived, tt.filename.IsArchived())
			require.Equal(tt.isCompressed, tt.filename.IsCompressed())
		})
	}
}

func TestFileNamePlatform(t *testing.T) {
	tests := []struct {
		name     string
		filename FileName
		os       OS
		arch     Arch
		platform Platform
	}{
		{
			name:     "gh_2.21.0_linux_amd64.tar.gz",
			filename: "gh_2.21.0_linux_amd64.tar.gz",
			os:       "linux",
			arch:     "amd64",
			platform: NewPlatform("linux", "amd64"),
		},
		{
			name:     "Empty",
			filename: "",
			os:       UnknownOS,
			arch:     "amd64",
			platform: NewPlatform(UnknownOS, "amd64"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.os, tt.filename.OS())
			require.Equal(tt.arch, tt.filename.Arch())
			require.Equal(tt.platform, tt.filename.Platform())
		})
	}
}
