package github

import (
	"os"
	"testing"
)

func TestAssets(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		assets      []asset
	}{
		{
			description: "shibataka000/go-get-release",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			assets: []asset{
				{
					name:        "go-get-release_v0.0.2_darwin_amd64",
					downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64",
				},
				{
					name:        "go-get-release_v0.0.2_linux_amd64",
					downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
				},
				{
					name:        "go-get-release_v0.0.2_windows_amd64.exe",
					downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe",
				},
			},
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo, err := c.Repository(tt.owner, tt.repo)
			if err != nil {
				t.Fatal(err)
			}
			release, err := repo.Release(tt.tag)
			if err != nil {
				t.Fatal(err)
			}

			assets, err := release.Assets()
			if err != nil {
				t.Fatal(err)
			}
			if len(assets) != len(tt.assets) {
				t.Fatalf("Expected is %s but actual is %s", tt.assets, assets)
			}
			for i := range assets {
				if assets[i].Name() != tt.assets[i].Name() || assets[i].DownloadURL() != tt.assets[i].DownloadURL() {
					t.Fatalf("Expected is %s but actual is %s", tt.assets, assets)
				}
			}
		})
	}
}
