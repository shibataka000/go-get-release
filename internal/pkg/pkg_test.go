package pkg

import (
	"os"
	"reflect"
	"testing"
)

func TestParsePkgName(t *testing.T) {
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
			owner, repo, version, err := parsePkgName(tt.in)
			if err != nil {
				t.Error(err)
				return
			}
			actual := []string{owner, repo, version}
			if !reflect.DeepEqual(actual, tt.out) {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestGetPkgInfo(t *testing.T) {
	tests := []struct {
		pkgName string
		pkgInfo *pkgInfo
		option  Option
	}{
		{
			pkgName: "shibataka000/go-get-release-test",
			pkgInfo: &pkgInfo{
				owner:       "shibataka000",
				repo:        "go-get-release-test",
				tag:         "v0.0.2",
				asset:       "go-get-release-test_v0.0.2_linux_amd64",
				downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release-test_v0.0.2_linux_amd64",
				binary:      "go-get-release-test",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "shibataka000/go-get-release-test=v0.0.1",
			pkgInfo: &pkgInfo{
				owner:       "shibataka000",
				repo:        "go-get-release-test",
				tag:         "v0.0.1",
				asset:       "go-get-release-test_v0.0.1_linux_amd64",
				downloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.1/go-get-release-test_v0.0.1_linux_amd64",
				binary:      "go-get-release-test",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "argo=v2.4.3",
			pkgInfo: &pkgInfo{
				owner:       "argoproj",
				repo:        "argo",
				tag:         "v2.4.3",
				asset:       "argo-linux-amd64",
				downloadURL: "https://github.com/argoproj/argo/releases/download/v2.4.3/argo-linux-amd64",
				binary:      "argo",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "argo=v2.4.3",
			pkgInfo: &pkgInfo{
				owner:       "argoproj",
				repo:        "argo",
				tag:         "v2.4.3",
				asset:       "argo-windows-amd64",
				downloadURL: "https://github.com/argoproj/argo/releases/download/v2.4.3/argo-windows-amd64",
				binary:      "argo.exe",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "argo=v2.4.3",
			pkgInfo: &pkgInfo{
				owner:       "argoproj",
				repo:        "argo",
				tag:         "v2.4.3",
				asset:       "argo-darwin-amd64",
				downloadURL: "https://github.com/argoproj/argo/releases/download/v2.4.3/argo-darwin-amd64",
				binary:      "argo",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "docker/compose=1.25.4",
			pkgInfo: &pkgInfo{
				owner:       "docker",
				repo:        "compose",
				tag:         "1.25.4",
				asset:       "docker-compose-Linux-x86_64",
				downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
				binary:      "docker-compose",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "docker/compose=1.25.4",
			pkgInfo: &pkgInfo{
				owner:       "docker",
				repo:        "compose",
				tag:         "1.25.4",
				asset:       "docker-compose-Windows-x86_64.exe",
				downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
				binary:      "docker-compose.exe",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "docker/compose=1.25.4",
			pkgInfo: &pkgInfo{
				owner:       "docker",
				repo:        "compose",
				tag:         "1.25.4",
				asset:       "docker-compose-Darwin-x86_64",
				downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
				binary:      "docker-compose",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "docker/machine=v0.16.2",
			pkgInfo: &pkgInfo{
				owner:       "docker",
				repo:        "machine",
				tag:         "v0.16.2",
				asset:       "docker-machine-Linux-x86_64",
				downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
				binary:      "docker-machine",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "docker/machine=v0.16.2",
			pkgInfo: &pkgInfo{
				owner:       "docker",
				repo:        "machine",
				tag:         "v0.16.2",
				asset:       "docker-machine-Windows-x86_64.exe",
				downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
				binary:      "docker-machine.exe",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "docker/machine=v0.16.2",
			pkgInfo: &pkgInfo{
				owner:       "docker",
				repo:        "machine",
				tag:         "v0.16.2",
				asset:       "docker-machine-Darwin-x86_64",
				downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
				binary:      "docker-machine",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "helm=v3.1.0",
			pkgInfo: &pkgInfo{
				owner:       "helm",
				repo:        "helm",
				tag:         "v3.1.0",
				asset:       "helm-v3.1.0-linux-amd64.tar.gz",
				downloadURL: "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
				binary:      "helm",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "helm=v3.1.0",
			pkgInfo: &pkgInfo{
				owner:       "helm",
				repo:        "helm",
				tag:         "v3.1.0",
				asset:       "helm-v3.1.0-windows-amd64.zip",
				downloadURL: "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
				binary:      "helm.exe",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "helm=v3.1.0",
			pkgInfo: &pkgInfo{
				owner:       "helm",
				repo:        "helm",
				tag:         "v3.1.0",
				asset:       "helm-v3.1.0-darwin-amd64.tar.gz",
				downloadURL: "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
				binary:      "helm",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "istio=1.3.8",
			pkgInfo: &pkgInfo{
				owner:       "istio",
				repo:        "istio",
				tag:         "1.3.8",
				asset:       "istio-1.3.8-linux.tar.gz",
				downloadURL: "https://github.com/istio/istio/releases/download/1.3.8/istio-1.3.8-linux.tar.gz",
				binary:      "istioctl",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "istio=1.3.8",
			pkgInfo: &pkgInfo{
				owner:       "istio",
				repo:        "istio",
				tag:         "1.3.8",
				asset:       "istio-1.3.8-win.zip",
				downloadURL: "https://github.com/istio/istio/releases/download/1.3.8/istio-1.3.8-win.zip",
				binary:      "istioctl.exe",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "istio=1.3.8",
			pkgInfo: &pkgInfo{
				owner:       "istio",
				repo:        "istio",
				tag:         "1.3.8",
				asset:       "istio-1.3.8-osx.tar.gz",
				downloadURL: "https://github.com/istio/istio/releases/download/1.3.8/istio-1.3.8-osx.tar.gz",
				binary:      "istioctl",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "kubectl-bindrole=v1.2.0",
			pkgInfo: &pkgInfo{
				owner:       "Ladicle",
				repo:        "kubectl-bindrole",
				tag:         "v1.2.0",
				asset:       "kubectl-bindrole_linux-amd64.tar.gz",
				downloadURL: "https://github.com/Ladicle/kubectl-bindrole/releases/download/v1.2.0/kubectl-bindrole_linux-amd64.tar.gz",
				binary:      "kubectl-bindrole",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "kubectl-bindrole=v1.2.0",
			pkgInfo: &pkgInfo{
				owner:       "Ladicle",
				repo:        "kubectl-bindrole",
				tag:         "v1.2.0",
				asset:       "kubectl-bindrole_windows-amd64.tar.gz",
				downloadURL: "https://github.com/Ladicle/kubectl-bindrole/releases/download/v1.2.0/kubectl-bindrole_windows-amd64.tar.gz",
				binary:      "kubectl-bindrole.exe",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "kubectl-bindrole=v1.2.0",
			pkgInfo: &pkgInfo{
				owner:       "Ladicle",
				repo:        "kubectl-bindrole",
				tag:         "v1.2.0",
				asset:       "kubectl-bindrole_darwin-amd64.tar.gz",
				downloadURL: "https://github.com/Ladicle/kubectl-bindrole/releases/download/v1.2.0/kubectl-bindrole_darwin-amd64.tar.gz",
				binary:      "kubectl-bindrole",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "kustomize=kustomize/v3.5.4",
			pkgInfo: &pkgInfo{
				owner:       "kubernetes-sigs",
				repo:        "kustomize",
				tag:         "kustomize/v3.5.4",
				asset:       "kustomize_v3.5.4_linux_amd64.tar.gz",
				downloadURL: "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v3.5.4/kustomize_v3.5.4_linux_amd64.tar.gz",
				binary:      "kustomize",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "kustomize=kustomize/v3.5.4",
			pkgInfo: &pkgInfo{
				owner:       "kubernetes-sigs",
				repo:        "kustomize",
				tag:         "kustomize/v3.5.4",
				asset:       "kustomize_v3.5.4_windows_amd64.tar.gz",
				downloadURL: "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v3.5.4/kustomize_v3.5.4_windows_amd64.tar.gz",
				binary:      "kustomize.exe",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "kustomize=kustomize/v3.5.4",
			pkgInfo: &pkgInfo{
				owner:       "kubernetes-sigs",
				repo:        "kustomize",
				tag:         "kustomize/v3.5.4",
				asset:       "kustomize_v3.5.4_darwin_amd64.tar.gz",
				downloadURL: "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v3.5.4/kustomize_v3.5.4_darwin_amd64.tar.gz",
				binary:      "kustomize",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "stern=1.11.0",
			pkgInfo: &pkgInfo{
				owner:       "wercker",
				repo:        "stern",
				tag:         "1.11.0",
				asset:       "stern_linux_amd64",
				downloadURL: "https://github.com/wercker/stern/releases/download/1.11.0/stern_linux_amd64",
				binary:      "stern",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "stern=1.11.0",
			pkgInfo: &pkgInfo{
				owner:       "wercker",
				repo:        "stern",
				tag:         "1.11.0",
				asset:       "stern_windows_amd64.exe",
				downloadURL: "https://github.com/wercker/stern/releases/download/1.11.0/stern_windows_amd64.exe",
				binary:      "stern.exe",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "stern=1.11.0",
			pkgInfo: &pkgInfo{
				owner:       "wercker",
				repo:        "stern",
				tag:         "1.11.0",
				asset:       "stern_darwin_amd64",
				downloadURL: "https://github.com/wercker/stern/releases/download/1.11.0/stern_darwin_amd64",
				binary:      "stern",
				isArchived:  false,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "terraform=v0.12.20",
			pkgInfo: &pkgInfo{
				owner:       "hashicorp",
				repo:        "terraform",
				tag:         "v0.12.20",
				asset:       "terraform_0.12.20_linux_amd64.zip",
				downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
				binary:      "terraform",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "linux",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "terraform=v0.12.20",
			pkgInfo: &pkgInfo{
				owner:       "hashicorp",
				repo:        "terraform",
				tag:         "v0.12.20",
				asset:       "terraform_0.12.20_windows_amd64.zip",
				downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
				binary:      "terraform.exe",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "windows",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
		{
			pkgName: "terraform=v0.12.20",
			pkgInfo: &pkgInfo{
				owner:       "hashicorp",
				repo:        "terraform",
				tag:         "v0.12.20",
				asset:       "terraform_0.12.20_darwin_amd64.zip",
				downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
				binary:      "terraform",
				isArchived:  true,
			},
			option: Option{
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				OS:          "darwin",
				Arch:        "amd64",
				ShowPrompt:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.pkgName, func(t *testing.T) {
			actual, err := getPkgInfo(tt.pkgName, &tt.option)
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(actual, tt.pkgInfo) {
				t.Errorf("Expected is %v but actual is %v", tt.pkgInfo, actual)
			}
		})
	}
}

func TestIsArchived(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64",
			out: false,
		},
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64.zip",
			out: true,
		},
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64.tar.gz",
			out: true,
		},
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64.exe",
			out: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			actual := isArchived(tt.in)
			if actual != tt.out {
				t.Errorf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestGetGoosByAsset(t *testing.T) {
	tests := []struct {
		asset  string
		goos   string
		goarch string
	}{
		{
			asset:  "argo-linux-amd64",
			goos:   "linux",
			goarch: "amd64",
		},
		{
			asset:  "argo-windows-amd64",
			goos:   "windows",
			goarch: "amd64",
		},
		{
			asset:  "argo-darwin-amd64",
			goos:   "darwin",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Linux-x86_64",
			goos:   "linux",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Windows-x86_64.exe",
			goos:   "windows",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Darwin-x86_64",
			goos:   "darwin",
			goarch: "amd64",
		},
		{
			asset:  "istio-1.3.8-linux.tar.gz",
			goos:   "linux",
			goarch: "amd64",
		},
		{
			asset:  "istio-1.3.8-win.zip",
			goos:   "windows",
			goarch: "amd64",
		},
		{
			asset:  "istio-1.3.8-osx.tar.gz",
			goos:   "darwin",
			goarch: "amd64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.asset, func(t *testing.T) {
			goos, err := getGoosByAsset(tt.asset)
			if err != nil {
				t.Error(err)
				return
			}
			if goos != tt.goos {
				t.Errorf("Expected is %v but actual is %v", tt.goos, goos)
				return
			}
			goarch, err := getGoarchByAsset(tt.asset)
			if err != nil {
				t.Error(err)
				return
			}
			if goarch != tt.goarch {
				t.Errorf("Expected is %v but actual is %v", tt.goarch, goarch)
				return
			}
		})
	}
}
