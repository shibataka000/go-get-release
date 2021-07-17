package github

import (
	"os"
	"testing"
)

func TestGoosAndGoarch(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		asset       string
		goos        string
		goarch      string
	}{
		{
			// todo: implements
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
			asset, err := release.Asset(tt.asset)
			if err != nil {
				t.Fatal(err)
			}
			goos, err := asset.Goos()
			if err != nil {
				t.Fatal(err)
			}
			goarch, err := asset.Goarch()
			if err != nil {
				t.Fatal(err)
			}
			if goos != tt.goos || goarch != tt.goarch {
				t.Fatalf("Expected is %s/%s but actual is %s/%s", tt.goos, tt.goarch, goos, goarch)
			}
		})
	}
}

func TestIsReleaseBinary(t *testing.T) {
	// todo: implements
}

func TestHasExt(t *testing.T) {
	// todo: implements
}
