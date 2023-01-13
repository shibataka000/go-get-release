package application

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/stretchr/testify/require"
)

func TestInstall(t *testing.T) {
	tests := []struct {
		query         string
		verifyCommand []string
	}{
		{
			query:         "terraform",
			verifyCommand: []string{"terraform", "version"},
		},
		{
			query:         "istio",
			verifyCommand: []string{"istioctl", "version", "--remote=false"},
		},
		{
			query:         "protocolbuffers/protobuf",
			verifyCommand: []string{"protoc", "--version"},
		},
		{
			query:         "vmware-tanzu/velero",
			verifyCommand: []string{"velero", "--help"},
		},
		{
			query:         "argoproj/argo-workflows",
			verifyCommand: []string{"argo", "version"},
		},
		{
			query:         "buildpacks/pack",
			verifyCommand: []string{"pack", "--version"},
		},
		{
			query:         "koalaman/shellcheck",
			verifyCommand: []string{"shellcheck", "--version"},
		},
	}

	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	platform := platform.New(os.Getenv("GOOS"), os.Getenv("GOARCH"))
	interactive := false

	for _, tt := range tests {
		name := tt.query
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			var err error

			cmd := exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			err = cmd.Run()
			assert.Error(err, "binary is already installed")

			installDir, err := os.MkdirTemp("", "*")
			assert.NoError(err)
			t.Setenv("PATH", installDir)
			defer os.RemoveAll(installDir)

			service, err := NewService(ctx, token)
			assert.NoError(err)
			command, err := NewCommandFromQuery(tt.query, platform.OS(), platform.Arch(), installDir, interactive)
			assert.NoError(err)
			err = service.Install(ctx, command)
			assert.NoError(err)

			cmd = exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			err = cmd.Run()
			assert.NoError(err)
		})
	}
}

func TestFindMetadata(t *testing.T) {
	tests := []struct {
		query          string
		platform       platform.Platform
		downloadURL    string
		execBinaryName string
	}{
		// aquasecurity/tfsec
		{
			query:          "aquasecurity/tfsec=v1.1.5",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-linux-amd64",
			execBinaryName: "tfsec",
		},
		{
			query:          "aquasecurity/tfsec=v1.1.5",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-darwin-amd64",
			execBinaryName: "tfsec",
		},
		{
			query:          "aquasecurity/tfsec=v1.1.5",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-windows-amd64.exe",
			execBinaryName: "tfsec.exe",
		},
		// aquasecurity/trivy
		{
			query:          "aquasecurity/trivy=v0.17.2",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_Linux-64bit.tar.gz",
			execBinaryName: "trivy",
		},
		{
			query:          "aquasecurity/trivy=v0.17.2",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_macOS-64bit.tar.gz",
			execBinaryName: "trivy",
		},
		// argoproj/argo-cd
		{
			query:          "argoproj/argo-cd=v2.0.4",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-linux-amd64",
			execBinaryName: "argocd",
		},
		{
			query:          "argoproj/argo-cd=v2.0.4",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-darwin-amd64",
			execBinaryName: "argocd",
		},
		{
			query:          "argoproj/argo-cd=v2.0.4",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-cd/releases/download/v2.0.4/argocd-windows-amd64.exe",
			execBinaryName: "argocd.exe",
		},
		// argoproj/argo-rollouts
		{
			query:          "argoproj/argo-rollouts=v0.9.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-linux-amd64",
			execBinaryName: "kubectl-argo-rollouts",
		},
		{
			query:          "argoproj/argo-rollouts=v0.9.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-darwin-amd64",
			execBinaryName: "kubectl-argo-rollouts",
		},
		// argoproj/argo-workflows
		{
			query:          "argoproj/argo-workflows=v3.0.7",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-linux-amd64.gz",
			execBinaryName: "argo",
		},
		{
			query:          "argoproj/argo-workflows=v3.0.7",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-darwin-amd64.gz",
			execBinaryName: "argo",
		},
		{
			query:          "argoproj/argo-workflows=v3.0.7",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-windows-amd64.gz",
			execBinaryName: "argo.exe",
		},
		// aws/amazon-ec2-instance-selector
		{
			query:          "aws/amazon-ec2-instance-selector=v2.0.2",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-linux-amd64",
			execBinaryName: "ec2-instance-selector",
		},
		{
			query:          "aws/amazon-ec2-instance-selector=v2.0.2",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-darwin-amd64",
			execBinaryName: "ec2-instance-selector",
		},
		{
			query:          "aws/amazon-ec2-instance-selector=v2.0.2",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-windows-amd64",
			execBinaryName: "ec2-instance-selector.exe",
		},
		// buildpacks/pack
		{
			query:          "buildpacks/pack=v0.19.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-linux.tgz",
			execBinaryName: "pack",
		},
		{
			query:          "buildpacks/pack=v0.19.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-macos.tgz",
			execBinaryName: "pack",
		},
		{
			query:          "buildpacks/pack=v0.19.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-windows.zip",
			execBinaryName: "pack.exe",
		},
		// CircleCI-Public/circleci-cli
		{
			query:          "CircleCI-Public/circleci-cli=v0.1.8764",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_linux_amd64.tar.gz",
			execBinaryName: "circleci",
		},
		{
			query:          "CircleCI-Public/circleci-cli=v0.1.8764",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_darwin_amd64.tar.gz",
			execBinaryName: "circleci",
		},
		{
			query:          "CircleCI-Public/circleci-cli=v0.1.8764",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_windows_amd64.zip",
			execBinaryName: "circleci.exe",
		},
		// cli/cli
		{
			query:          "cli/cli=v1.12.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_linux_amd64.tar.gz",
			execBinaryName: "gh",
		},
		{
			query:          "cli/cli=v1.12.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_macOS_amd64.tar.gz",
			execBinaryName: "gh",
		},
		{
			query:          "cli/cli=v1.12.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_windows_amd64.zip",
			execBinaryName: "gh.exe",
		},
		// docker/compose
		{
			query:          "docker/compose=1.25.4",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
			execBinaryName: "docker-compose",
		},
		{
			query:          "docker/compose=1.25.4",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
			execBinaryName: "docker-compose",
		},
		{
			query:          "docker/compose=1.25.4",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
			execBinaryName: "docker-compose.exe",
		},
		// docker/machine
		{
			query:          "docker/machine=v0.16.2",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
			execBinaryName: "docker-machine",
		},
		{
			query:          "docker/machine=v0.16.2",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
			execBinaryName: "docker-machine",
		},
		{
			query:          "docker/machine=v0.16.2",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
			execBinaryName: "docker-machine.exe",
		},
		// docker/scan-cli-plugin
		{
			query:          "docker/scan-cli-plugin=v0.17.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_linux_amd64",
			execBinaryName: "docker-scan",
		},
		{
			query:          "docker/scan-cli-plugin=v0.17.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_darwin_amd64",
			execBinaryName: "docker-scan",
		},
		{
			query:          "docker/scan-cli-plugin=v0.17.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_windows_amd64.exe",
			execBinaryName: "docker-scan.exe",
		},
		// fluxcd/flux2
		{
			query:          "fluxcd/flux2=v0.8.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_linux_amd64.tar.gz",
			execBinaryName: "flux",
		},
		{
			query:          "fluxcd/flux2=v0.8.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_darwin_amd64.tar.gz",
			execBinaryName: "flux",
		},
		{
			query:          "fluxcd/flux2=v0.8.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_windows_amd64.zip",
			execBinaryName: "flux.exe",
		},
		// goodwithtech/dockle
		{
			query:          "goodwithtech/dockle=v0.3.1",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Linux-64bit.tar.gz",
			execBinaryName: "dockle",
		},
		{
			query:          "goodwithtech/dockle=v0.3.1",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_macOS-64bit.tar.gz",
			execBinaryName: "dockle",
		},
		{
			query:          "goodwithtech/dockle=v0.3.1",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Windows-64bit.zip",
			execBinaryName: "dockle.exe",
		},
		// gravitational/teleport
		{
			query:          "gravitational/teleport=v8.1.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://get.gravitational.com/teleport-v8.1.0-linux-amd64-bin.tar.gz",
			execBinaryName: "tsh",
		},
		{
			query:          "gravitational/teleport=v8.1.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://get.gravitational.com/teleport-v8.1.0-darwin-amd64-bin.tar.gz",
			execBinaryName: "tsh",
		},
		{
			query:          "gravitational/teleport=v8.1.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://get.gravitational.com/teleport-v8.1.0-windows-amd64-bin.zip",
			execBinaryName: "tsh.exe",
		},
		// hashicorp/terraform
		{
			query:          "hashicorp/terraform=v0.12.20",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
			execBinaryName: "terraform",
		},
		{
			query:          "hashicorp/terraform=v0.12.20",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
			execBinaryName: "terraform",
		},
		{
			query:          "hashicorp/terraform=v0.12.20",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
			execBinaryName: "terraform.exe",
		},
		// helm/helm
		{
			query:          "helm/helm=v3.1.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
			execBinaryName: "helm",
		},
		{
			query:          "helm/helm=v3.1.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
			execBinaryName: "helm",
		},
		{
			query:          "helm/helm=v3.1.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
			execBinaryName: "helm.exe",
		},
		// istio/istio
		{
			query:          "istio/istio=1.6.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-linux-amd64.tar.gz",
			execBinaryName: "istioctl",
		},
		{
			query:          "istio/istio=1.6.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-osx-amd64.tar.gz",
			execBinaryName: "istioctl",
		},
		{
			query:          "istio/istio=1.6.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/istio/istio/releases/download/1.6.0/istioctl-1.6.0-win-amd64.zip",
			execBinaryName: "istioctl.exe",
		},
		// mikefarah/yq
		{
			query:          "mikefarah/yq=v4.7.1",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_linux_amd64",
			execBinaryName: "yq",
		},
		{
			query:          "mikefarah/yq=v4.7.1",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_darwin_amd64",
			execBinaryName: "yq",
		},
		{
			query:          "mikefarah/yq=v4.7.1",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_windows_amd64.exe",
			execBinaryName: "yq.exe",
		},
		// mozilla/sops
		{
			query:          "mozilla/sops=v3.5.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
			execBinaryName: "sops",
		},
		{
			query:          "mozilla/sops=v3.5.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
			execBinaryName: "sops",
		},
		{
			query:          "mozilla/sops=v3.5.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
			execBinaryName: "sops.exe",
		},
		// open-policy-agent/conftest
		{
			query:          "open-policy-agent/conftest=v0.21.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Linux_x86_64.tar.gz",
			execBinaryName: "conftest",
		},
		{
			query:          "open-policy-agent/conftest=v0.21.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Darwin_x86_64.tar.gz",
			execBinaryName: "conftest",
		},
		{
			query:          "open-policy-agent/conftest=v0.21.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Windows_x86_64.zip",
			execBinaryName: "conftest.exe",
		},
		// open-policy-agent/opa
		{
			query:          "open-policy-agent/opa=v0.29.4",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_linux_amd64",
			execBinaryName: "opa",
		},
		{
			query:          "open-policy-agent/opa=v0.29.4",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_darwin_amd64",
			execBinaryName: "opa",
		},
		{
			query:          "open-policy-agent/opa=v0.29.4",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_windows_amd64.exe",
			execBinaryName: "opa.exe",
		},
		// protocolbuffers/protobuf
		{
			query:          "protocolbuffers/protobuf=v3.11.4",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
			execBinaryName: "protoc",
		},
		{
			query:          "protocolbuffers/protobuf=v3.11.4",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip",
			execBinaryName: "protoc",
		},
		{
			query:          "protocolbuffers/protobuf=v3.11.4",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-win64.zip",
			execBinaryName: "protoc.exe",
		},
		// starship/starship
		{
			query:          "starship/starship=v0.47.0",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-unknown-linux-gnu.tar.gz",
			execBinaryName: "starship",
		},
		{
			query:          "starship/starship=v0.47.0",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-apple-darwin.tar.gz",
			execBinaryName: "starship",
		},
		{
			query:          "starship/starship=v0.47.0",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-pc-windows-msvc.zip",
			execBinaryName: "starship.exe",
		},
		// viaduct-ai/kustomize-sops
		{
			query:          "viaduct-ai/kustomize-sops=v2.3.3",
			platform:       platform.New("linux", "amd64"),
			downloadURL:    "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Linux_x86_64.tar.gz",
			execBinaryName: "ksops",
		},
		{
			query:          "viaduct-ai/kustomize-sops=v2.3.3",
			platform:       platform.New("darwin", "amd64"),
			downloadURL:    "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Darwin_x86_64.tar.gz",
			execBinaryName: "ksops",
		},
		{
			query:          "viaduct-ai/kustomize-sops=v2.3.3",
			platform:       platform.New("windows", "amd64"),
			downloadURL:    "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Windows_x86_64.tar.gz",
			execBinaryName: "ksops.exe",
		},
	}

	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	installDir := ""
	interactive := false

	for _, tt := range tests {
		name := tt.query
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)

			service, err := NewService(ctx, token)
			assert.NoError(err)

			command, err := NewCommandFromQuery(tt.query, tt.platform.OS(), tt.platform.Arch(), installDir, interactive)
			assert.NoError(err)

			meta, err := service.findMetadata(ctx, command)
			assert.NoError(err)
			assert.Equal(tt.downloadURL, meta.Asset().DownloadURL())
			assert.Equal(tt.execBinaryName, meta.ExecBinary().Name())
		})
	}
}
