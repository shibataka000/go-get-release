package pkg

import (
	"os"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		input  *FindInput
		output *FindOutput
	}{
		{
			input: &FindInput{
				Name:        "shibataka000/go-get-release-test",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "shibataka000",
				Repo:        "go-get-release-test",
				Tag:         "v0.0.2",
				Asset:       "go-get-release_v0.0.2_linux_amd64",
				DownloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
				BinaryName:  "go-get-release-test",
				IsArchived:  false,
			},
		},
		// shibataka000/go-get-release-test
		{
			input: &FindInput{
				Name:        "shibataka000/go-get-release-test=v0.0.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "shibataka000",
				Repo:        "go-get-release-test",
				Tag:         "v0.0.2",
				Asset:       "go-get-release_v0.0.2_linux_amd64",
				DownloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
				BinaryName:  "go-get-release-test",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "shibataka000/go-get-release-test=v0.0.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "shibataka000",
				Repo:        "go-get-release-test",
				Tag:         "v0.0.2",
				Asset:       "go-get-release_v0.0.2_darwin_amd64",
				DownloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64",
				BinaryName:  "go-get-release-test",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "shibataka000/go-get-release-test=v0.0.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "shibataka000",
				Repo:        "go-get-release-test",
				Tag:         "v0.0.2",
				Asset:       "go-get-release_v0.0.2_windows_amd64.exe",
				DownloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe",
				BinaryName:  "go-get-release-test.exe",
				IsArchived:  false,
			},
		},
		// docker/compose
		{
			input: &FindInput{
				Name:        "docker/compose=1.25.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "docker",
				Repo:        "compose",
				Tag:         "1.25.4",
				Asset:       "docker-compose-Linux-x86_64",
				DownloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
				BinaryName:  "docker-compose",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "docker/compose=1.25.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "docker",
				Repo:        "compose",
				Tag:         "1.25.4",
				Asset:       "docker-compose-Darwin-x86_64",
				DownloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
				BinaryName:  "docker-compose",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "docker/compose=1.25.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "docker",
				Repo:        "compose",
				Tag:         "1.25.4",
				Asset:       "docker-compose-Windows-x86_64.exe",
				DownloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
				BinaryName:  "docker-compose.exe",
				IsArchived:  false,
			},
		},
		// docker/machine
		{
			input: &FindInput{
				Name:        "docker/machine=v0.16.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "docker",
				Repo:        "machine",
				Tag:         "v0.16.2",
				Asset:       "docker-machine-Linux-x86_64",
				DownloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
				BinaryName:  "docker-machine",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "docker/machine=v0.16.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "docker",
				Repo:        "machine",
				Tag:         "v0.16.2",
				Asset:       "docker-machine-Darwin-x86_64",
				DownloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
				BinaryName:  "docker-machine",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "docker/machine=v0.16.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "docker",
				Repo:        "machine",
				Tag:         "v0.16.2",
				Asset:       "docker-machine-Windows-x86_64.exe",
				DownloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
				BinaryName:  "docker-machine.exe",
				IsArchived:  false,
			},
		},
		// helm/helm
		{
			input: &FindInput{
				Name:        "helm/helm=v3.1.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "helm",
				Repo:        "helm",
				Tag:         "v3.1.0",
				Asset:       "helm-v3.1.0-linux-amd64.tar.gz",
				DownloadURL: "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
				BinaryName:  "helm",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "helm/helm=v3.1.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "helm",
				Repo:        "helm",
				Tag:         "v3.1.0",
				Asset:       "helm-v3.1.0-darwin-amd64.tar.gz",
				DownloadURL: "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
				BinaryName:  "helm",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "helm/helm=v3.1.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "helm",
				Repo:        "helm",
				Tag:         "v3.1.0",
				Asset:       "helm-v3.1.0-windows-amd64.zip",
				DownloadURL: "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
				BinaryName:  "helm.exe",
				IsArchived:  true,
			},
		},
		// istio/istio
		{
			input: &FindInput{
				Name:        "istio/istio=1.6.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "istio",
				Repo:        "istio",
				Tag:         "1.6.0",
				Asset:       "istioctl-1.6.0-linux-amd64.tar.gz",
				DownloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-linux-amd64.tar.gz",
				BinaryName:  "istioctl",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "istio/istio=1.6.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "istio",
				Repo:        "istio",
				Tag:         "1.6.0",
				Asset:       "istioctl-1.6.0-osx.tar.gz",
				DownloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-osx.tar.gz",
				BinaryName:  "istioctl",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "istio/istio=1.6.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "istio",
				Repo:        "istio",
				Tag:         "1.6.0",
				Asset:       "istioctl-1.6.0-win.zip",
				DownloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-win.zip",
				BinaryName:  "istioctl.exe",
				IsArchived:  true,
			},
		},
		// hashicorp/terraform
		{
			input: &FindInput{
				Name:        "hashicorp/terraform=v0.12.20",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "hashicorp",
				Repo:        "terraform",
				Tag:         "v0.12.20",
				Asset:       "terraform_0.12.20_linux_amd64.zip",
				DownloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
				BinaryName:  "terraform",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "hashicorp/terraform=v0.12.20",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "hashicorp",
				Repo:        "terraform",
				Tag:         "v0.12.20",
				Asset:       "terraform_0.12.20_darwin_amd64.zip",
				DownloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
				BinaryName:  "terraform",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "hashicorp/terraform=v0.12.20",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "hashicorp",
				Repo:        "terraform",
				Tag:         "v0.12.20",
				Asset:       "terraform_0.12.20_windows_amd64.zip",
				DownloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
				BinaryName:  "terraform.exe",
				IsArchived:  true,
			},
		},
		// argoproj/argo-cd
		{
			input: &FindInput{
				Name:        "argoproj/argo-cd=v2.0.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-cd",
				Tag:         "v2.0.4",
				Asset:       "argocd-linux-amd64",
				DownloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-linux-amd64",
				BinaryName:  "argocd",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "argoproj/argo-cd=v2.0.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-cd",
				Tag:         "v2.0.4",
				Asset:       "argocd-darwin-amd64",
				DownloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-darwin-amd64",
				BinaryName:  "argocd",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "argoproj/argo-cd=v2.0.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-cd",
				Tag:         "v2.0.4",
				Asset:       "argocd-windows-amd64.exe",
				DownloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-windows-amd64.exe",
				BinaryName:  "argocd.exe",
				IsArchived:  false,
			},
		},
		// protocolbuffers/protobuf
		{
			input: &FindInput{
				Name:        "protocolbuffers/protobuf=v3.11.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "protocolbuffers",
				Repo:        "protobuf",
				Tag:         "v3.11.4",
				Asset:       "protoc-3.11.4-linux-x86_64.zip",
				DownloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
				BinaryName:  "protoc",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "protocolbuffers/protobuf=v3.11.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "protocolbuffers",
				Repo:        "protobuf",
				Tag:         "v3.11.4",
				Asset:       "protoc-3.11.4-osx-x86_64.zip",
				DownloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip",
				BinaryName:  "protoc",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "protocolbuffers/protobuf=v3.11.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "protocolbuffers",
				Repo:        "protobuf",
				Tag:         "v3.11.4",
				Asset:       "protoc-3.11.4-win64.zip",
				DownloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-win64.zip",
				BinaryName:  "protoc.exe",
				IsArchived:  true,
			},
		},
		// mozilla/sops
		{
			input: &FindInput{
				Name:        "mozilla/sops=v3.5.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "mozilla",
				Repo:        "sops",
				Tag:         "v3.5.0",
				Asset:       "sops-v3.5.0.linux",
				DownloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
				BinaryName:  "sops",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "mozilla/sops=v3.5.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "mozilla",
				Repo:        "sops",
				Tag:         "v3.5.0",
				Asset:       "sops-v3.5.0.darwin",
				DownloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
				BinaryName:  "sops",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "mozilla/sops=v3.5.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "mozilla",
				Repo:        "sops",
				Tag:         "v3.5.0",
				Asset:       "sops-v3.5.0.exe",
				DownloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
				BinaryName:  "sops.exe",
				IsArchived:  false,
			},
		},
		// CircleCI-Public/circleci-cli
		{
			input: &FindInput{
				Name:        "CircleCI-Public/circleci-cli=v0.1.8764",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "CircleCI-Public",
				Repo:        "circleci-cli",
				Tag:         "v0.1.8764",
				Asset:       "circleci-cli_0.1.8764_linux_amd64.tar.gz",
				DownloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_linux_amd64.tar.gz",
				BinaryName:  "circleci",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "CircleCI-Public/circleci-cli=v0.1.8764",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "CircleCI-Public",
				Repo:        "circleci-cli",
				Tag:         "v0.1.8764",
				Asset:       "circleci-cli_0.1.8764_darwin_amd64.tar.gz",
				DownloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_darwin_amd64.tar.gz",
				BinaryName:  "circleci",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "CircleCI-Public/circleci-cli=v0.1.8764",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "CircleCI-Public",
				Repo:        "circleci-cli",
				Tag:         "v0.1.8764",
				Asset:       "circleci-cli_0.1.8764_windows_amd64.zip",
				DownloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_windows_amd64.zip",
				BinaryName:  "circleci.exe",
				IsArchived:  true,
			},
		},
		// rgoproj/argo-rollouts
		{
			input: &FindInput{
				Name:        "argoproj/argo-rollouts=v0.9.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-rollouts",
				Tag:         "v0.9.0",
				Asset:       "kubectl-argo-rollouts-linux-amd64",
				DownloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-linux-amd64",
				BinaryName:  "kubectl-argo-rollouts",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "argoproj/argo-rollouts=v0.9.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-rollouts",
				Tag:         "v0.9.0",
				Asset:       "kubectl-argo-rollouts-darwin-amd64",
				DownloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-darwin-amd64",
				BinaryName:  "kubectl-argo-rollouts",
				IsArchived:  false,
			},
		},
		// open-policy-agent/conftest
		{
			input: &FindInput{
				Name:        "open-policy-agent/conftest=v0.21.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "open-policy-agent",
				Repo:        "conftest",
				Tag:         "v0.21.0",
				Asset:       "conftest_0.21.0_Linux_x86_64.tar.gz",
				DownloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Linux_x86_64.tar.gz",
				BinaryName:  "conftest",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "open-policy-agent/conftest=v0.21.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "open-policy-agent",
				Repo:        "conftest",
				Tag:         "v0.21.0",
				Asset:       "conftest_0.21.0_Darwin_x86_64.tar.gz",
				DownloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Darwin_x86_64.tar.gz",
				BinaryName:  "conftest",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "open-policy-agent/conftest=v0.21.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "open-policy-agent",
				Repo:        "conftest",
				Tag:         "v0.21.0",
				Asset:       "conftest_0.21.0_Windows_x86_64.zip",
				DownloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Windows_x86_64.zip",
				BinaryName:  "conftest.exe",
				IsArchived:  true,
			},
		},
		// goodwithtech/dockle
		{
			input: &FindInput{
				Name:        "goodwithtech/dockle=v0.3.1",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "goodwithtech",
				Repo:        "dockle",
				Tag:         "v0.3.1",
				Asset:       "dockle_0.3.1_Linux-64bit.tar.gz",
				DownloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Linux-64bit.tar.gz",
				BinaryName:  "dockle",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "goodwithtech/dockle=v0.3.1",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "goodwithtech",
				Repo:        "dockle",
				Tag:         "v0.3.1",
				Asset:       "dockle_0.3.1_macOS-64bit.tar.gz",
				DownloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_macOS-64bit.tar.gz",
				BinaryName:  "dockle",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "goodwithtech/dockle=v0.3.1",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "goodwithtech",
				Repo:        "dockle",
				Tag:         "v0.3.1",
				Asset:       "dockle_0.3.1_Windows-64bit.zip",
				DownloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Windows-64bit.zip",
				BinaryName:  "dockle.exe",
				IsArchived:  true,
			},
		},
		// starship/starship
		{
			input: &FindInput{
				Name:        "starship/starship=v0.47.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "starship",
				Repo:        "starship",
				Tag:         "v0.47.0",
				Asset:       "starship-x86_64-unknown-linux-gnu.tar.gz",
				DownloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-unknown-linux-gnu.tar.gz",
				BinaryName:  "starship",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "starship/starship=v0.47.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "starship",
				Repo:        "starship",
				Tag:         "v0.47.0",
				Asset:       "starship-x86_64-apple-darwin.tar.gz",
				DownloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-apple-darwin.tar.gz",
				BinaryName:  "starship",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "starship/starship=v0.47.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "starship",
				Repo:        "starship",
				Tag:         "v0.47.0",
				Asset:       "starship-x86_64-pc-windows-msvc.zip",
				DownloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-pc-windows-msvc.zip",
				BinaryName:  "starship.exe",
				IsArchived:  true,
			},
		},
		// viaduct-ai/kustomize-sops
		{
			input: &FindInput{
				Name:        "viaduct-ai/kustomize-sops=v2.3.3",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "viaduct-ai",
				Repo:        "kustomize-sops",
				Tag:         "v2.3.3",
				Asset:       "ksops_2.3.3_Linux_x86_64.tar.gz",
				DownloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Linux_x86_64.tar.gz",
				BinaryName:  "ksops",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "viaduct-ai/kustomize-sops=v2.3.3",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "viaduct-ai",
				Repo:        "kustomize-sops",
				Tag:         "v2.3.3",
				Asset:       "ksops_2.3.3_Darwin_x86_64.tar.gz",
				DownloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Darwin_x86_64.tar.gz",
				BinaryName:  "ksops",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "viaduct-ai/kustomize-sops=v2.3.3",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "viaduct-ai",
				Repo:        "kustomize-sops",
				Tag:         "v2.3.3",
				Asset:       "ksops_2.3.3_Windows_x86_64.tar.gz",
				DownloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Windows_x86_64.tar.gz",
				BinaryName:  "ksops.exe",
				IsArchived:  true,
			},
		},
		// fluxcd/flux2
		{
			input: &FindInput{
				Name:        "fluxcd/flux2=v0.8.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "fluxcd",
				Repo:        "flux2",
				Tag:         "v0.8.0",
				Asset:       "flux_0.8.0_linux_amd64.tar.gz",
				DownloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_linux_amd64.tar.gz",
				BinaryName:  "flux",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "fluxcd/flux2=v0.8.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "fluxcd",
				Repo:        "flux2",
				Tag:         "v0.8.0",
				Asset:       "flux_0.8.0_darwin_amd64.tar.gz",
				DownloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_darwin_amd64.tar.gz",
				BinaryName:  "flux",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "fluxcd/flux2=v0.8.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "fluxcd",
				Repo:        "flux2",
				Tag:         "v0.8.0",
				Asset:       "flux_0.8.0_windows_amd64.zip",
				DownloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_windows_amd64.zip",
				BinaryName:  "flux.exe",
				IsArchived:  true,
			},
		},
		// mikefarah/yq
		{
			input: &FindInput{
				Name:        "mikefarah/yq=v4.7.1",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "mikefarah",
				Repo:        "yq",
				Tag:         "v4.7.1",
				Asset:       "yq_linux_amd64",
				DownloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_linux_amd64",
				BinaryName:  "yq",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "mikefarah/yq=v4.7.1",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "mikefarah",
				Repo:        "yq",
				Tag:         "v4.7.1",
				Asset:       "yq_darwin_amd64",
				DownloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_darwin_amd64",
				BinaryName:  "yq",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "mikefarah/yq=v4.7.1",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "mikefarah",
				Repo:        "yq",
				Tag:         "v4.7.1",
				Asset:       "yq_windows_amd64.exe",
				DownloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_windows_amd64.exe",
				BinaryName:  "yq.exe",
				IsArchived:  false,
			},
		},
		// aquasecurity/trivy
		{
			input: &FindInput{
				Name:        "aquasecurity/trivy=v0.17.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "aquasecurity",
				Repo:        "trivy",
				Tag:         "v0.17.2",
				Asset:       "trivy_0.17.2_Linux-64bit.tar.gz",
				DownloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_Linux-64bit.tar.gz",
				BinaryName:  "trivy",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "aquasecurity/trivy=v0.17.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "aquasecurity",
				Repo:        "trivy",
				Tag:         "v0.17.2",
				Asset:       "trivy_0.17.2_macOS-64bit.tar.gz",
				DownloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_macOS-64bit.tar.gz",
				BinaryName:  "trivy",
				IsArchived:  true,
			},
		},
		// aws/amazon-ec2-instance-selector
		{
			input: &FindInput{
				Name:        "aws/amazon-ec2-instance-selector=v2.0.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "aws",
				Repo:        "amazon-ec2-instance-selector",
				Tag:         "v2.0.2",
				Asset:       "ec2-instance-selector-linux-amd64",
				DownloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-linux-amd64",
				BinaryName:  "ec2-instance-selector",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "aws/amazon-ec2-instance-selector=v2.0.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "aws",
				Repo:        "amazon-ec2-instance-selector",
				Tag:         "v2.0.2",
				Asset:       "ec2-instance-selector-darwin-amd64",
				DownloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-darwin-amd64",
				BinaryName:  "ec2-instance-selector",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "aws/amazon-ec2-instance-selector=v2.0.2",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "aws",
				Repo:        "amazon-ec2-instance-selector",
				Tag:         "v2.0.2",
				Asset:       "ec2-instance-selector-windows-amd64",
				DownloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-windows-amd64",
				BinaryName:  "ec2-instance-selector.exe",
				IsArchived:  false,
			},
		},
		// argoproj/argo-workflows
		{
			input: &FindInput{
				Name:        "argoproj/argo-workflows=v3.0.7",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-workflows",
				Tag:         "v3.0.7",
				Asset:       "argo-linux-amd64.gz",
				DownloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-linux-amd64.gz",
				BinaryName:  "argo",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "argoproj/argo-workflows=v3.0.7",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-workflows",
				Tag:         "v3.0.7",
				Asset:       "argo-darwin-amd64.gz",
				DownloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-darwin-amd64.gz",
				BinaryName:  "argo",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "argoproj/argo-workflows=v3.0.7",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "argoproj",
				Repo:        "argo-workflows",
				Tag:         "v3.0.7",
				Asset:       "argo-windows-amd64.gz",
				DownloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-windows-amd64.gz",
				BinaryName:  "argo.exe",
				IsArchived:  true,
			},
		},
		// open-policy-agent/opa
		{
			input: &FindInput{
				Name:        "open-policy-agent/opa=v0.29.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "open-policy-agent",
				Repo:        "opa",
				Tag:         "v0.29.4",
				Asset:       "opa_linux_amd64",
				DownloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_linux_amd64",
				BinaryName:  "opa",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "open-policy-agent/opa=v0.29.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "open-policy-agent",
				Repo:        "opa",
				Tag:         "v0.29.4",
				Asset:       "opa_darwin_amd64",
				DownloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_darwin_amd64",
				BinaryName:  "opa",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "open-policy-agent/opa=v0.29.4",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "open-policy-agent",
				Repo:        "opa",
				Tag:         "v0.29.4",
				Asset:       "opa_windows_amd64.exe",
				DownloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_windows_amd64.exe",
				BinaryName:  "opa.exe",
				IsArchived:  false,
			},
		},
		// buildpacks/pack
		{
			input: &FindInput{
				Name:        "buildpacks/pack=v0.19.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "buildpacks",
				Repo:        "pack",
				Tag:         "v0.19.0",
				Asset:       "pack-v0.19.0-linux.tgz",
				DownloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-linux.tgz",
				BinaryName:  "pack",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "buildpacks/pack=v0.19.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "darwin",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "buildpacks",
				Repo:        "pack",
				Tag:         "v0.19.0",
				Asset:       "pack-v0.19.0-macos.tgz",
				DownloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-macos.tgz",
				BinaryName:  "pack",
				IsArchived:  true,
			},
		},
		{
			input: &FindInput{
				Name:        "buildpacks/pack=v0.19.0",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "windows",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "buildpacks",
				Repo:        "pack",
				Tag:         "v0.19.0",
				Asset:       "pack-v0.19.0-windows.zip",
				DownloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-windows.zip",
				BinaryName:  "pack.exe",
				IsArchived:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input.Name, func(t *testing.T) {
			actual, err := Find(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(actual, tt.output) {
				t.Fatalf("Expected is %v but actual is %v", tt.output, actual)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			in:  "go-get-release",
			out: []string{"", "go-get-release", ""},
		},
		{
			in:  "shibataka000/go-get-release",
			out: []string{"shibataka000", "go-get-release", ""},
		},
		{
			in:  "go-get-release=1.0.0",
			out: []string{"", "go-get-release", "1.0.0"},
		},
		{
			in:  "shibataka000/go-get-release=1.0.0",
			out: []string{"shibataka000", "go-get-release", "1.0.0"},
		},
		{
			in:  "shibataka000/go-get-release=alpha/1.0.0",
			out: []string{"shibataka000", "go-get-release", "alpha/1.0.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			owner, repo, version, err := parse(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			actual := []string{owner, repo, version}
			if !reflect.DeepEqual(actual, tt.out) {
				t.Fatalf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}
