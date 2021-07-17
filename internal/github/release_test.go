package github

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestAssets(t *testing.T) {
	tests := []struct {
		description  string
		owner        string
		repo         string
		tag          string
		downloadURLs []string
	}{
		{
			description: "hashicorp/terraform(v0.12.20)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			downloadURLs: []string{
				"https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64",
				"https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
				"https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe",
			},
		},
		{
			description: "hashicorp/terraform(v0.12.20)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			downloadURLs: []string{
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS.348FFC4C.sig",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS.72D7468F.sig",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS.sig",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_freebsd_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_freebsd_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_freebsd_arm.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_arm.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_openbsd_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_openbsd_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_solaris_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
			},
		},
		{
			description: "hashicorp/terraform(v0.12.20)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			downloadURLs: []string{
				"https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-386.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-arm.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-arm64.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-ppc64le.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-s390x.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
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

			downloadURLs := []string{}
			for _, a := range assets {
				downloadURLs = append(downloadURLs, a.DownloadURL())
			}

			if !reflect.DeepEqual(tt.downloadURLs, downloadURLs) {
				t.Fatalf("Expected is %v but actual is %v", tt.downloadURLs, downloadURLs)
			}
		})
	}
}

func TestAsset(t *testing.T) {
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

func TestFindAssetByPlatform(t *testing.T) {
	tests := []struct {
		owner       string
		repo        string
		tag         string
		goos        string
		goarch      string
		downloadURL string
		binaryName  string
	}{
		// aquasecurity/trivy
		{
			owner:       "aquasecurity",
			repo:        "trivy",
			tag:         "v0.17.2",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_Linux-64bit.tar.gz",
			binaryName:  "trivy",
		},
		{
			owner:       "aquasecurity",
			repo:        "trivy",
			tag:         "v0.17.2",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_macOS-64bit.tar.gz",
			binaryName:  "trivy",
		},
		// argoproj/argo-workflows
		{
			owner:       "argoproj",
			repo:        "argo-workflows",
			tag:         "v3.0.7",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-linux-amd64.gz",
			binaryName:  "argo",
		},
		{
			owner:       "argoproj",
			repo:        "argo-workflows",
			tag:         "v3.0.7",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-darwin-amd64.gz",
			binaryName:  "argo",
		},
		{
			owner:       "argoproj",
			repo:        "argo-workflows",
			tag:         "v3.0.7",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-windows-amd64.gz",
			binaryName:  "argo.exe",
		},
		// argoproj/argo-cd
		{
			owner:       "argoproj",
			repo:        "argo-cd",
			tag:         "v2.0.4",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-linux-amd64",
			binaryName:  "argocd",
		},
		{
			owner:       "argoproj",
			repo:        "argo-cd",
			tag:         "v2.0.4",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-darwin-amd64",
			binaryName:  "argocd",
		},
		{
			owner:       "argoproj",
			repo:        "argo-cd",
			tag:         "v2.0.4",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-windows-amd64.exe",
			binaryName:  "argocd.exe",
		},
		// rgoproj/argo-rollouts
		{
			owner:       "argoproj",
			repo:        "argo-rollouts",
			tag:         "v0.9.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-linux-amd64",
			binaryName:  "kubectl-argo-rollouts",
		},
		{
			owner:       "argoproj",
			repo:        "argo-rollouts",
			tag:         "v0.9.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-darwin-amd64",
			binaryName:  "kubectl-argo-rollouts",
		},
		// aws/amazon-ec2-instance-selector
		{
			owner:       "aws",
			repo:        "amazon-ec2-instance-selector",
			tag:         "v2.0.2",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-linux-amd64",
			binaryName:  "ec2-instance-selector",
		},
		{
			owner:       "aws",
			repo:        "amazon-ec2-instance-selector",
			tag:         "v2.0.2",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-darwin-amd64",
			binaryName:  "ec2-instance-selector",
		},
		{
			owner:       "aws",
			repo:        "amazon-ec2-instance-selector",
			tag:         "v2.0.2",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-windows-amd64",
			binaryName:  "ec2-instance-selector.exe",
		},
		// buildpacks/pack
		{
			owner:       "buildpacks",
			repo:        "pack",
			tag:         "v0.19.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-linux.tgz",
			binaryName:  "pack",
		},
		{
			owner:       "buildpacks",
			repo:        "pack",
			tag:         "v0.19.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-macos.tgz",
			binaryName:  "pack",
		},
		{
			owner:       "buildpacks",
			repo:        "pack",
			tag:         "v0.19.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-windows.zip",
			binaryName:  "pack.exe",
		},
		// CircleCI-Public/circleci-cli
		{
			owner:       "CircleCI-Public",
			repo:        "circleci-cli",
			tag:         "v0.1.8764",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_linux_amd64.tar.gz",
			binaryName:  "circleci",
		},
		{
			owner:       "CircleCI-Public",
			repo:        "circleci-cli",
			tag:         "v0.1.8764",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_darwin_amd64.tar.gz",
			binaryName:  "circleci",
		},
		{
			owner:       "CircleCI-Public",
			repo:        "circleci-cli",
			tag:         "v0.1.8764",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_windows_amd64.zip",
			binaryName:  "circleci.exe",
		},
		// cli/cli
		{
			owner:       "cli",
			repo:        "cli",
			tag:         "v1.12.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_linux_amd64.tar.gz",
			binaryName:  "gh",
		},
		{
			owner:       "cli",
			repo:        "cli",
			tag:         "v1.12.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_macOS_amd64.tar.gz",
			binaryName:  "gh",
		},
		{
			owner:       "cli",
			repo:        "cli",
			tag:         "v1.12.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_windows_amd64.zip",
			binaryName:  "gh.exe",
		},
		// docker/compose
		{
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
			binaryName:  "docker-compose",
		},
		{
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
			binaryName:  "docker-compose",
		},
		{
			owner:       "docker",
			repo:        "compose",
			tag:         "1.25.4",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
			binaryName:  "docker-compose.exe",
		},
		// docker/machine
		{
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
			binaryName:  "docker-machine",
		},
		{
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
			binaryName:  "docker-machine",
		},
		{
			owner:       "docker",
			repo:        "machine",
			tag:         "v0.16.2",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
			binaryName:  "docker-machine.exe",
		},
		// fluxcd/flux2
		{
			owner:       "fluxcd",
			repo:        "flux2",
			tag:         "v0.8.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_linux_amd64.tar.gz",
			binaryName:  "flux",
		},
		{
			owner:       "fluxcd",
			repo:        "flux2",
			tag:         "v0.8.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_darwin_amd64.tar.gz",
			binaryName:  "flux",
		},
		{
			owner:       "fluxcd",
			repo:        "flux2",
			tag:         "v0.8.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_windows_amd64.zip",
			binaryName:  "flux.exe",
		},
		// goodwithtech/dockle
		{
			owner:       "goodwithtech",
			repo:        "dockle",
			tag:         "v0.3.1",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Linux-64bit.tar.gz",
			binaryName:  "dockle",
		},
		{
			owner:       "goodwithtech",
			repo:        "dockle",
			tag:         "v0.3.1",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_macOS-64bit.tar.gz",
			binaryName:  "dockle",
		},
		{
			owner:       "goodwithtech",
			repo:        "dockle",
			tag:         "v0.3.1",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Windows-64bit.zip",
			binaryName:  "dockle.exe",
		},
		// hashicorp/terraform
		{
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
			binaryName:  "terraform",
		},
		{
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
			binaryName:  "terraform",
		},
		{
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
			binaryName:  "terraform.exe",
		},
		// helm/helm
		{
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
			binaryName:  "helm",
		},
		{
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
			binaryName:  "helm",
		},
		{
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
			binaryName:  "helm.exe",
		},
		// istio/istio
		{
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-linux-amd64.tar.gz",
			binaryName:  "istioctl",
		},
		{
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-osx.tar.gz",
			binaryName:  "istioctl",
		},
		{
			owner:       "istio",
			repo:        "istio",
			tag:         "1.6.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-win.zip",
			binaryName:  "istioctl.exe",
		},
		// mikefarah/yq
		{
			owner:       "mikefarah",
			repo:        "yq",
			tag:         "v4.7.1",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_linux_amd64",
			binaryName:  "yq",
		},
		{
			owner:       "mikefarah",
			repo:        "yq",
			tag:         "v4.7.1",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_darwin_amd64",
			binaryName:  "yq",
		},
		{
			owner:       "mikefarah",
			repo:        "yq",
			tag:         "v4.7.1",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_windows_amd64.exe",
			binaryName:  "yq.exe",
		},
		// mozilla/sops
		{
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
			binaryName:  "sops",
		},
		{
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
			binaryName:  "sops",
		},
		{
			owner:       "mozilla",
			repo:        "sops",
			tag:         "v3.5.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
			binaryName:  "sops.exe",
		},
		// open-policy-agent/conftest
		{
			owner:       "open-policy-agent",
			repo:        "conftest",
			tag:         "v0.21.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Linux_x86_64.tar.gz",
			binaryName:  "conftest",
		},
		{
			owner:       "open-policy-agent",
			repo:        "conftest",
			tag:         "v0.21.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Darwin_x86_64.tar.gz",
			binaryName:  "conftest",
		},
		{
			owner:       "open-policy-agent",
			repo:        "conftest",
			tag:         "v0.21.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Windows_x86_64.zip",
			binaryName:  "conftest.exe",
		},
		// open-policy-agent/opa
		{
			owner:       "open-policy-agent",
			repo:        "opa",
			tag:         "v0.29.4",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_linux_amd64",
			binaryName:  "opa",
		},
		{
			owner:       "open-policy-agent",
			repo:        "opa",
			tag:         "v0.29.4",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_darwin_amd64",
			binaryName:  "opa",
		},
		{
			owner:       "open-policy-agent",
			repo:        "opa",
			tag:         "v0.29.4",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_windows_amd64.exe",
			binaryName:  "opa.exe",
		},
		// protocolbuffers/protobuf
		{
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
			binaryName:  "protoc",
		},
		{
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip",
			binaryName:  "protoc",
		},
		{
			owner:       "protocolbuffers",
			repo:        "protobuf",
			tag:         "v3.11.4",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-win64.zip",
			binaryName:  "protoc.exe",
		},
		// starship/starship
		{
			owner:       "starship",
			repo:        "starship",
			tag:         "v0.47.0",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-unknown-linux-gnu.tar.gz",
			binaryName:  "starship",
		},
		{
			owner:       "starship",
			repo:        "starship",
			tag:         "v0.47.0",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-apple-darwin.tar.gz",
			binaryName:  "starship",
		},
		{
			owner:       "starship",
			repo:        "starship",
			tag:         "v0.47.0",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-pc-windows-msvc.zip",
			binaryName:  "starship.exe",
		},
		// viaduct-ai/kustomize-sops
		{
			owner:       "viaduct-ai",
			repo:        "kustomize-sops",
			tag:         "v2.3.3",
			goos:        "linux",
			goarch:      "amd64",
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Linux_x86_64.tar.gz",
			binaryName:  "ksops",
		},
		{
			owner:       "viaduct-ai",
			repo:        "kustomize-sops",
			tag:         "v2.3.3",
			goos:        "darwin",
			goarch:      "amd64",
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Darwin_x86_64.tar.gz",
			binaryName:  "ksops",
		},
		{
			owner:       "viaduct-ai",
			repo:        "kustomize-sops",
			tag:         "v2.3.3",
			goos:        "windows",
			goarch:      "amd64",
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Windows_x86_64.tar.gz",
			binaryName:  "ksops.exe",
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s/%s", tt.owner, tt.repo)
		t.Run(name, func(t *testing.T) {
			repo, err := c.Repository(tt.owner, tt.repo)
			if err != nil {
				t.Fatal(err)
			}
			release, err := repo.Release(tt.tag)
			if err != nil {
				t.Fatal(err)
			}

			asset, err := release.FindAssetByPlatform(tt.goos, tt.goarch)
			if err != nil {
				t.Fatal(err)
			}
			if asset.DownloadURL() != tt.downloadURL {
				t.Fatalf("Expected is %s but actual is %s", tt.downloadURL, asset.DownloadURL())
			}

			binaryName, err := asset.BinaryName()
			if err != nil {
				t.Fatal(err)
			}
			if binaryName != tt.binaryName {
				t.Fatalf("Expected is %s but actual is %s", tt.binaryName, binaryName)
			}

			if !asset.IsReleaseBinary() {
				t.Fatalf("Asset does not contain release binary")
			}
		})
	}
}
