package github

import (
	"os"
	"testing"
)

func TestGetAsset(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
		asset       string
		downloadURL string
		binaryName  string
		goos        string
		goarch      string
	}{
		{
			description: "shibataka000/go-get-release-test:v0.0.2(linux,amd64)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			asset:       "go-get-release_v0.0.2_linux_amd64",
			downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
			binaryName:  "go-get-release-test",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "shibataka000/go-get-release-test:v0.0.2(windows,amd64)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			asset:       "go-get-release_v0.0.2_windows_amd64.exe",
			downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe",
			binaryName:  "go-get-release-test.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "shibataka000/go-get-release-test:v0.0.2(darwin,amd64)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			asset:       "go-get-release_v0.0.2_darwin_amd64",
			downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64",
			binaryName:  "go-get-release-test",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "docker/compose:1.25.4(linux,amd64)",
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			asset:       "docker-compose-Linux-x86_64",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
			binaryName:  "docker-compose",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "docker/compose:1.25.4(windows,amd64)",
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			asset:       "docker-compose-Windows-x86_64.exe",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
			binaryName:  "docker-compose.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "docker/compose:1.25.4(darwin,amd64)",
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			asset:       "docker-compose-Darwin-x86_64",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
			binaryName:  "docker-compose",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "docker/machine:v0.16.2(linux,amd64)",
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			asset:       "docker-machine-Linux-x86_64",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
			binaryName:  "docker-machine",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "docker/machine:v0.16.2(windows,amd64)",
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			asset:       "docker-machine-Windows-x86_64.exe",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
			binaryName:  "docker-machine.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "docker/machine:v0.16.2(darwin,amd64)",
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			asset:       "docker-machine-Darwin-x86_64",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
			binaryName:  "docker-machine",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "helm/helm(linux,amd64)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			asset:       "helm-v3.1.0-linux-amd64.tar.gz",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
			binaryName:  "helm",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "helm/helm(windows,amd64)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			asset:       "helm-v3.1.0-windows-amd64.zip",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
			binaryName:  "helm.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "helm/helm(darwin,amd64)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			asset:       "helm-v3.1.0-darwin-amd64.tar.gz",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
			binaryName:  "helm",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "istio/istio:1.6.0(linux,amd64)",
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			asset:       "istioctl-1.6.0-linux-amd64.tar.gz",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-linux-amd64.tar.gz",
			binaryName:  "istioctl",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "istio/istio:1.6.0(windows,amd64)",
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			asset:       "istioctl-1.6.0-win.zip",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-win.zip",
			binaryName:  "istioctl.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "istio/istio:1.6.0(darwin,amd64)",
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			asset:       "istioctl-1.6.0-osx.tar.gz",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-osx.tar.gz",
			binaryName:  "istioctl",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "hashicorp/terraform:v0.12.20(linux,amd64)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			asset:       "terraform_0.12.20_linux_amd64.zip",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
			binaryName:  "terraform",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "hashicorp/terraform:v0.12.20(windows,amd64)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			asset:       "terraform_0.12.20_windows_amd64.zip",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
			binaryName:  "terraform.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "hashicorp/terraform:v0.12.20(darwin,amd64)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			asset:       "terraform_0.12.20_darwin_amd64.zip",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
			binaryName:  "terraform",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-cd:v1.4.2(linux,amd64)",
			owner:       "argoproj",
			repo:        "argo-cd",
			tag:         "v1.4.2",
			asset:       "argocd-linux-amd64",
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v1.4.2/argocd-linux-amd64",
			binaryName:  "argocd",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "argoproj/argo-cd:v1.4.2(darwin,amd64)",
			owner:       "argoproj",
			repo:        "argo-cd",
			tag:         "v1.4.2",
			asset:       "argocd-darwin-amd64",
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v1.4.2/argocd-darwin-amd64",
			binaryName:  "argocd",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "protocolbuffers/protobuf:v3.11.4(linux,amd64)",
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			asset:       "protoc-3.11.4-linux-x86_64.zip",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
			binaryName:  "protoc",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "protocolbuffers/protobuf:v3.11.4(windows,amd64)",
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			asset:       "protoc-3.11.4-win64.zip",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-win64.zip",
			binaryName:  "protoc.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "protocolbuffers/protobuf:v3.11.4(darwin,amd64)",
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			asset:       "protoc-3.11.4-osx-x86_64.zip",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip",
			binaryName:  "protoc",
			goos:        "darwin",
			goarch:      "amd64",
		},
		{
			description: "mozilla/sops:v3.5.0(linux,amd64)",
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			asset:       "sops-v3.5.0.linux",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
			binaryName:  "sops",
			goos:        "linux",
			goarch:      "amd64",
		},
		{
			description: "mozilla/sops:v3.5.0(windows,amd64)",
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			asset:       "sops-v3.5.0.exe",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
			binaryName:  "sops.exe",
			goos:        "windows",
			goarch:      "amd64",
		},
		{
			description: "mozilla/sops:v3.5.0(darwin,amd64)",
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			asset:       "sops-v3.5.0.darwin",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
			binaryName:  "sops",
			goos:        "darwin",
			goarch:      "amd64",
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := c.GetRepository(tt.owner, tt.repo)
			release, err := repo.GetRelease(tt.tag)
			if err != nil {
				t.Error(err)
				return
			}
			asset, err := release.GetAsset(tt.goos, tt.goarch)
			if asset.Name() != tt.asset || asset.DownloadURL() != tt.downloadURL || asset.BinaryName() != tt.binaryName {
				t.Errorf("Expected is {%s %s %s} but actual is {%s %s %s}", tt.asset, tt.downloadURL, tt.binaryName, asset.Name(), asset.DownloadURL(), asset.BinaryName())
				return
			}
		})
	}
}
