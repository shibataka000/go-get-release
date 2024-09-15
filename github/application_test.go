package github

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	// todo: implement this.
}

func TestInstall(t *testing.T) {
	// todo: implement this.
}

func TestFindAndInstallOnLinuxAmd64(t *testing.T) {
	tests := []struct {
		name string

		repoFullName       string
		tag                string
		assetPatterns      []string
		execBinaryPatterns []string

		asset      Asset
		execBinary ExecBinary

		test *exec.Cmd
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()

			dir, err := os.MkdirTemp("", "")
			require.NoError(err)
			defer os.RemoveAll(dir)

			t.Setenv("PATH", dir)

			require.Error(tt.test.Run(), "executable binary was already installed")

			app := NewApplicationService(
				NewAssetRepository(ctx, githubTokenForTest),
				NewExecBinaryRepository(),
			)

			asset, execBinary, err := app.Find(ctx, tt.repoFullName, tt.tag, tt.assetPatterns, tt.execBinaryPatterns)
			require.NoError(err)

			err = app.Install(ctx, tt.repoFullName, asset, execBinary, dir)
			require.NoError(err)

			require.NoError(tt.test.Run())

		})
	}

}

// githubTokenForTest is authentication token for github.com API requests. This can be used for test only.
var githubTokenForTest = os.Getenv("GITHUB_TOKEN")
