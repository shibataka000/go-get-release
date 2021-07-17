package github

import (
	"os"
	"testing"
)

func TestAssetName(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		assetName   string
	}{
		{
			description: "shibataka000/go-get-release-test",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			assetName:   "go-get-release_v0.0.2_linux_amd64",
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
			asset, err := release.Asset(tt.assetName)
			if err != nil {
				t.Fatal(err)
			}
			if asset.Name() != tt.assetName {
				t.Fatalf("Expected is %s but actual is %s", tt.assetName, asset.Name())
			}
		})
	}
}

func TestBinaryName(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		assetName   string
		binaryName  string
	}{
		{
			description: "shibataka000/go-get-release-test",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			assetName:   "go-get-release_v0.0.2_linux_amd64",
			binaryName:  "go-get-release-test",
		},
		{
			description: "shibataka000/go-get-release-test_windows",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			assetName:   "go-get-release_v0.0.2_windows_amd64.exe",
			binaryName:  "go-get-release-test.exe",
		},
		{
			description: "docker/compose",
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			assetName:   "docker-compose-Linux-x86_64",
			binaryName:  "docker-compose",
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
			asset, err := release.Asset(tt.assetName)
			if err != nil {
				t.Fatal(err)
			}
			binaryName, err := asset.BinaryName()
			if err != nil {
				t.Fatal(err)
			}
			if binaryName != tt.binaryName {
				t.Fatalf("Expected is %s but actual is %s", tt.binaryName, binaryName)
			}
		})
	}
}

func TestGoosAndGoarch(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		assetName   string
		goos        string
		goarch      string
	}{
		{
			description: "shibataka000/go-get-release-test",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			assetName:   "go-get-release_v0.0.2_linux_amd64",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "shibataka000/go-get-release-test",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			assetName:   "go-get-release_v0.0.2_darwin_amd64",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "shibataka000/go-get-release-test",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			assetName:   "go-get-release_v0.0.2_windows_amd64.exe",
			goos:        "windows",
			goarch:      "amd64",
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
			asset, err := release.Asset(tt.assetName)
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
	tests := []struct {
		description     string
		owner           string
		repo            string
		tag             string
		assetName       string
		isReleaseBinary bool
	}{
		{
			description:     "shibataka000/go-get-release-test",
			owner:           "shibataka000",
			repo:            "go-get-release-test",
			tag:             "v0.0.2",
			assetName:       "go-get-release_v0.0.2_linux_amd64",
			isReleaseBinary: true,
		},
		{
			description:     "shibataka000/go-get-release-test",
			owner:           "shibataka000",
			repo:            "go-get-release-test",
			tag:             "v0.0.2",
			assetName:       "go-get-release_v0.0.2_windows_amd64.exe",
			isReleaseBinary: true,
		},
		{
			owner:           "hashicorp",
			repo:            "terraform",
			tag:             "v0.12.20",
			assetName:       "terraform_0.12.20_linux_amd64.zip",
			isReleaseBinary: true,
		},
		{
			owner:           "docker",
			repo:            "compose",
			tag:             "1.25.4",
			assetName:       "docker-compose-Linux-x86_64.sha256",
			isReleaseBinary: false,
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
			asset, err := release.Asset(tt.assetName)
			if err != nil {
				t.Fatal(err)
			}
			if asset.IsReleaseBinary() != tt.isReleaseBinary {
				t.Fatalf("Expected is %v but actual is %v", tt.isReleaseBinary, asset.IsReleaseBinary())
			}
		})
	}
}

func TestHasExt(t *testing.T) {
	tests := []struct {
		description string
		name        string
		exts        []string
		hasExt      bool
	}{
		{
			description: "a.exe",
			name:        "a.exe",
			exts:        []string{".exe"},
			hasExt:      true,
		},
		{
			description: "a.exe_does_not_zip",
			name:        "a.exe",
			exts:        []string{".zip"},
			hasExt:      false,
		},
		{
			description: "a",
			name:        "a",
			exts:        []string{""},
			hasExt:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := hasExt(tt.name, tt.exts)
			if actual != tt.hasExt {
				t.Fatalf("Expected is %v but actual is %v", tt.hasExt, actual)
			}
		})
	}
}
