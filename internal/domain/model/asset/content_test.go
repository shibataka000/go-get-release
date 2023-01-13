package asset

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContentFindExecBinary(t *testing.T) {
	tests := []struct {
		testfile           string
		execBinaryFileName string
		execBinary         ExecBinary
	}{
		{
			testfile:           "testdata/test",
			execBinaryFileName: "test",
			execBinary:         NewExecBinary([]byte("helloworld\n")),
		},
		{
			testfile:           "testdata/test.gz",
			execBinaryFileName: "test",
			execBinary:         NewExecBinary([]byte("helloworld\n")),
		},
		{
			testfile:           "testdata/test.tar.xz",
			execBinaryFileName: "test",
			execBinary:         NewExecBinary([]byte("helloworld\n")),
		},
		{
			testfile:           "testdata/test.zip",
			execBinaryFileName: "test",
			execBinary:         NewExecBinary([]byte("helloworld\n")),
		},
	}

	for _, tt := range tests {
		name := tt.testfile
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			body, err := os.ReadFile(tt.testfile)
			assert.NoError(err)
			filename := filepath.Base(tt.testfile)
			assetContent := NewContent(filename, body)
			execBinary, err := assetContent.FindExecBinary(tt.execBinaryFileName)
			assert.NoError(err)
			assert.Equal(tt.execBinary, execBinary)
		})
	}
}
