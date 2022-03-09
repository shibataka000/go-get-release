package archive

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		description string
		filename    string
	}{
		{
			description: "testdata.zip",
			filename:    "./testdata/testdata.zip",
		},
		{
			description: "testdata.tar.gz",
			filename:    "./testdata/testdata.tar.gz",
		},
		{
			description: "testdata.tgz",
			filename:    "./testdata/testdata.tgz",
		},
		{
			description: "testdata.tar.xz",
			filename:    "./testdata/testdata.tar.xz",
		},
		{
			description: "testdata.txt.gz",
			filename:    "./testdata/testdata.txt.gz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			tempDir, err := ioutil.TempDir("", "test-extract-")
			if err != nil {
				t.Fatal(err)
			}
			err = Extract(tt.filename, tempDir)
			if err != nil {
				t.Fatal(err)
			}

			if IsArchived(tt.filename) {
				err = diff("./testdata/testdata", filepath.Join(tempDir, "testdata"))
			} else {
				err = diff("./testdata/testdata/testdata.txt", filepath.Join(tempDir, "testdata.txt"))
			}
			if err != nil {
				t.Fatal(err)
			}

			err = os.RemoveAll(tempDir)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func diff(file1, file2 string) error {
	cmd := exec.Command("diff", "-r", file1, file2)
	return cmd.Run()
}
