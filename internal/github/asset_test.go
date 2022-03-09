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
			asset, err := release.AssetByName(tt.assetName)
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
			asset, err := release.AssetByName(tt.assetName)
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
		{
			description: "protocolbuffers/protobuf",
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.17.3",
			assetName:   "protoc-3.17.3-linux-ppcle_64.zip",
			goos:        "linux",
			goarch:      "ppc64le",
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
			asset, err := release.AssetByName(tt.assetName)
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

func TestContainReleaseBinary(t *testing.T) {
	tests := []struct {
		description          string
		owner                string
		repo                 string
		tag                  string
		assetName            string
		containReleaseBinary bool
	}{
		{
			description:          "shibataka000/go-get-release-test",
			owner:                "shibataka000",
			repo:                 "go-get-release-test",
			tag:                  "v0.0.2",
			assetName:            "go-get-release_v0.0.2_linux_amd64",
			containReleaseBinary: true,
		},
		{
			description:          "shibataka000/go-get-release-test",
			owner:                "shibataka000",
			repo:                 "go-get-release-test",
			tag:                  "v0.0.2",
			assetName:            "go-get-release_v0.0.2_windows_amd64.exe",
			containReleaseBinary: true,
		},
		{
			description:          "hashicorp/terraform",
			owner:                "hashicorp",
			repo:                 "terraform",
			tag:                  "v0.12.20",
			assetName:            "terraform_0.12.20_linux_amd64.zip",
			containReleaseBinary: true,
		},
		{
			description:          "docker/compose",
			owner:                "docker",
			repo:                 "compose",
			tag:                  "1.25.4",
			assetName:            "docker-compose-Linux-x86_64.sha256",
			containReleaseBinary: false,
		},
		{
			description:          "mozilla/sops",
			owner:                "mozilla",
			repo:                 "sops",
			tag:                  "v3.7.1",
			assetName:            "sops-v3.7.1.linux",
			containReleaseBinary: true,
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
			asset, err := release.AssetByName(tt.assetName)
			if err != nil {
				t.Fatal(err)
			}
			if asset.ContainReleaseBinary() != tt.containReleaseBinary {
				t.Fatalf("Expected is %v but actual is %v", tt.containReleaseBinary, asset.ContainReleaseBinary())
			}
		})
	}
}

func TestIsArchived(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		assetName   string
		isArchived  bool
	}{
		{
			description: "hashicorp/terraform",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			assetName:   "terraform_0.12.20_linux_amd64.zip",
			isArchived:  true,
		},
		{
			description: "helm/helm",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			assetName:   "helm-v3.1.0-linux-amd64.tar.gz",
			isArchived:  true,
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
			asset, err := release.AssetByName(tt.assetName)
			if err != nil {
				t.Fatal(err)
			}
			if asset.IsArchived() != tt.isArchived {
				t.Fatalf("Expected is %v but actual is %v", tt.isArchived, asset.IsArchived())
			}
		})
	}
}

func TestIsCompressed(t *testing.T) {
	tests := []struct {
		description  string
		owner        string
		repo         string
		tag          string
		assetName    string
		isCompressed bool
	}{
		{
			description:  "hashicorp/terraform",
			owner:        "hashicorp",
			repo:         "terraform",
			tag:          "v0.12.20",
			assetName:    "terraform_0.12.20_linux_amd64.zip",
			isCompressed: true,
		},
		{
			description:  "helm/helm",
			owner:        "helm",
			repo:         "helm",
			tag:          "v3.1.0",
			assetName:    "helm-v3.1.0-linux-amd64.tar.gz",
			isCompressed: true,
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
			asset, err := release.AssetByName(tt.assetName)
			if err != nil {
				t.Fatal(err)
			}
			if asset.IsCompressed() != tt.isCompressed {
				t.Fatalf("Expected is %v but actual is %v", tt.isCompressed, asset.IsCompressed())
			}
		})
	}
}

func TestIsExecBinary(t *testing.T) {
	tests := []struct {
		description  string
		owner        string
		repo         string
		tag          string
		assetName    string
		isExecBinary bool
	}{
		{
			description:  "shibataka000/go-get-release-test",
			owner:        "shibataka000",
			repo:         "go-get-release-test",
			tag:          "v0.0.2",
			assetName:    "go-get-release_v0.0.2_linux_amd64",
			isExecBinary: true,
		},
		{
			description:  "shibataka000/go-get-release-test",
			owner:        "shibataka000",
			repo:         "go-get-release-test",
			tag:          "v0.0.2",
			assetName:    "go-get-release_v0.0.2_windows_amd64.exe",
			isExecBinary: true,
		},
		{
			description:  "mozilla/sops",
			owner:        "mozilla",
			repo:         "sops",
			tag:          "v3.7.1",
			assetName:    "sops-v3.7.1.linux",
			isExecBinary: true,
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
			asset, err := release.AssetByName(tt.assetName)
			if err != nil {
				t.Fatal(err)
			}
			if asset.IsExecBinary() != tt.isExecBinary {
				t.Fatalf("Expected is %v but actual is %v", tt.isExecBinary, asset.IsExecBinary())
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
			description: "a.exe's_ext_is_not_zip",
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
