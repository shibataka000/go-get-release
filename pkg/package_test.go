package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPackageStringToPrompt(t *testing.T) {
	tests := []struct {
		name   string
		pkg    Package
		prompt string
	}{
		{
			name: "hashicorp/terraform",
			pkg: New(
				NewGitHubRepository("hashicorp", "terraform"),
				NewGitHubRelease(0, "0.12.20"),
				NewAssetMeta("https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip"),
				NewExecBinaryMeta("terraform"),
			),
			prompt: "Repo:\thashicorp/terraform\nTag:\t0.12.20\nAsset:\tterraform_0.12.20_linux_amd64.zip\nBinary:\tterraform",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.pkg.StringToPrompt(), tt.prompt)
		})
	}
}
