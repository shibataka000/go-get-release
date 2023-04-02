package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssetFileExecBinary(t *testing.T) {
	tests := []struct {
		name           string
		assetFilePath  string
		execBinaryMeta ExecBinaryMeta
		execBinaryFile ExecBinaryFile
	}{
		{
			name:           "./testdata/test",
			assetFilePath:  "./testdata/test",
			execBinaryMeta: NewExecBinaryMeta("test"),
			execBinaryFile: NewExecBinaryFile("test", []byte("helloworld\n")),
		},
		{
			name:           "./testdata/test.gz",
			assetFilePath:  "./testdata/test.gz",
			execBinaryMeta: NewExecBinaryMeta("test"),
			execBinaryFile: NewExecBinaryFile("test", []byte("helloworld\n")),
		},
		{
			name:           "./testdata/test.tar.gz",
			assetFilePath:  "./testdata/test.tar.gz",
			execBinaryMeta: NewExecBinaryMeta("test"),
			execBinaryFile: NewExecBinaryFile("test", []byte("helloworld\n")),
		},
		{
			name:           "./testdata/test.tar.xz",
			assetFilePath:  "./testdata/test.tar.xz",
			execBinaryMeta: NewExecBinaryMeta("test"),
			execBinaryFile: NewExecBinaryFile("test", []byte("helloworld\n")),
		},
		{
			name:           "./testdata/test.zip",
			assetFilePath:  "./testdata/test.zip",
			execBinaryMeta: NewExecBinaryMeta("test"),
			execBinaryFile: NewExecBinaryFile("test", []byte("helloworld\n")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			file, err := ReadTestFile(t, tt.assetFilePath)
			assert.NoError(err)
			asset := AssetFile(file)
			execBinary, err := asset.ExecBinary(tt.execBinaryMeta.Name)
			assert.NoError(err)
			assert.Equal(tt.execBinaryFile, execBinary)
		})
	}
}
