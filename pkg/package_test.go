package pkg

import (
	"fmt"
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
				NewRepository("hashicorp", "terraform"),
				NewRelease("0.12.20"),
				NewAsset("https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip"),
				NewExecBinary("terraform"),
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

func TestReleaseSemVer(t *testing.T) {
	tests := []struct {
		name    string
		release Release
		semver  string
		err     error
	}{
		{
			name:    "v1.2.3",
			release: NewRelease("v1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "1.2.3",
			release: NewRelease("1.2.3"),
			semver:  "1.2.3",
		},
		{
			name:    "x.y.z",
			release: NewRelease("x.y.z"),
			semver:  "",
			err:     fmt.Errorf("x.y.z is not valid semver"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			semver, err := tt.release.SemVer()
			if tt.err == nil {
				assert.NoError(err)
				assert.Equal(tt.semver, semver)
			} else {
				assert.EqualError(err, tt.err.Error())
			}
		})
	}
}
