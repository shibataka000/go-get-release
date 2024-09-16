package github

import (
	"context"
	"io"
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
		repoFullName string
		tag          string

		asset      Asset
		execBinary ExecBinary

		test *exec.Cmd
	}{
		{
			repoFullName: "aquasecurity/trivy",
			tag:          "v0.53.0",
			asset:        newAsset(0, "trivy_0.53.0_Linux-64bit.tar.gz"),
			execBinary:   newExecBinary("trivy"),
			test:         exec.Command("./trivy", "version"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.repoFullName, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()

			dir, err := os.MkdirTemp("", "")
			require.NoError(err)
			defer os.RemoveAll(dir)

			tt.test.Dir = dir

			require.Error(tt.test.Run(), "executable binary was already installed")

			app := NewApplicationService(
				NewAssetRepository(ctx, githubTokenForTest),
				NewExecBinaryRepository(),
			)

			asset, execBinary, err := app.Find(ctx, tt.repoFullName, tt.tag, DefaultAssetPatterns, DefaultExecBinaryPatterns)
			require.NoError(err)

			err = app.Install(ctx, tt.repoFullName, asset, execBinary, dir, io.Discard)
			require.NoError(err)

			require.NoError(tt.test.Run())

		})
	}

}

// githubTokenForTest is authentication token for github.com API requests. This can be used for test only.
var githubTokenForTest = os.Getenv("GITHUB_TOKEN")
