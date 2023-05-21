package pkg

import (
	"context"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewApplicationServiceForTest(ctx context.Context, t *testing.T) *ApplicationService {
	t.Helper()
	repository := NewInfrastructureRepository(ctx, os.Getenv("GITHUB_TOKEN"))
	factory := NewFactory()
	return NewApplicationService(repository, factory)
}

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
		{
			query:         "mozilla/sops",
			verifyCommand: []string{"sops", "--version"},
		},
	}

	for _, tt := range tests {
		name := tt.query
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)

			cmd := exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			err := cmd.Run()
			assert.Error(err, "binary is already installed")

			dir, err := os.MkdirTemp("", "*")
			assert.NoError(err)
			t.Setenv("PATH", dir)
			defer os.RemoveAll(dir)

			ctx := context.Background()
			app := NewApplicationServiceForTest(ctx, t)
			query, err := ParseQuery(tt.query)
			assert.NoError(err)
			platform := NewPlatform(os.Getenv("GOOS"), os.Getenv("GOARCH"))
			pkg, err := app.Search(ctx, query, platform)
			assert.NoError(err)

			err = app.Install(pkg, dir, io.Discard)
			assert.NoError(err)

			cmd = exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			err = cmd.Run()
			assert.NoError(err)
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		query       string
		platform    Platform
		downloadURL URL
		execBinary  FileName
	}{
		// aquasecurity/tfsec
		{
			query:       "aquasecurity/tfsec=v1.1.5",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-linux-amd64",
			execBinary:  "tfsec",
		},
		{
			query:       "aquasecurity/tfsec=v1.1.5",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-darwin-amd64",
			execBinary:  "tfsec",
		},
		{
			query:       "aquasecurity/tfsec=v1.1.5",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-windows-amd64.exe",
			execBinary:  "tfsec.exe",
		},
		// aquasecurity/trivy
		{
			query:       "aquasecurity/trivy=v0.17.2",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_Linux-64bit.tar.gz",
			execBinary:  "trivy",
		},
		{
			query:       "aquasecurity/trivy=v0.17.2",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/aquasecurity/trivy/releases/download/v0.17.2/trivy_0.17.2_macOS-64bit.tar.gz",
			execBinary:  "trivy",
		},
		// argoproj/argo-cd
		{
			query:       "argoproj/argo-cd=v2.6.7",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.6.7/argocd-linux-amd64",
			execBinary:  "argocd",
		},
		{
			query:       "argoproj/argo-cd=v2.6.7",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.6.7/argocd-darwin-amd64",
			execBinary:  "argocd",
		},
		{
			query:       "argoproj/argo-cd=v2.6.7",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-cd/releases/download/v2.6.7/argocd-windows-amd64.exe",
			execBinary:  "argocd.exe",
		},
		// argoproj/argo-rollouts
		{
			query:       "argoproj/argo-rollouts=v0.9.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-linux-amd64",
			execBinary:  "kubectl-argo-rollouts",
		},
		{
			query:       "argoproj/argo-rollouts=v0.9.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-darwin-amd64",
			execBinary:  "kubectl-argo-rollouts",
		},
		// argoproj/argo-workflows
		{
			query:       "argoproj/argo-workflows=v3.0.7",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-linux-amd64.gz",
			execBinary:  "argo",
		},
		{
			query:       "argoproj/argo-workflows=v3.0.7",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-darwin-amd64.gz",
			execBinary:  "argo",
		},
		{
			query:       "argoproj/argo-workflows=v3.0.7",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/argoproj/argo-workflows/releases/download/v3.0.7/argo-windows-amd64.gz",
			execBinary:  "argo.exe",
		},
		// aws/amazon-ec2-instance-selector
		{
			query:       "aws/amazon-ec2-instance-selector=v2.0.2",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-linux-amd64",
			execBinary:  "ec2-instance-selector",
		},
		{
			query:       "aws/amazon-ec2-instance-selector=v2.0.2",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-darwin-amd64",
			execBinary:  "ec2-instance-selector",
		},
		{
			query:       "aws/amazon-ec2-instance-selector=v2.0.2",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.0.2/ec2-instance-selector-windows-amd64",
			execBinary:  "ec2-instance-selector.exe",
		},
		// bitnami-labs/sealed-secrets
		{
			query:       "bitnami-labs/sealed-secrets=v0.20.2",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.20.2/kubeseal-0.20.2-linux-amd64.tar.gz",
			execBinary:  "kubeseal",
		},
		{
			query:       "bitnami-labs/sealed-secrets=v0.20.2",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.20.2/kubeseal-0.20.2-darwin-amd64.tar.gz",
			execBinary:  "kubeseal",
		},
		{
			query:       "bitnami-labs/sealed-secrets=v0.20.2",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.20.2/kubeseal-0.20.2-windows-amd64.tar.gz",
			execBinary:  "kubeseal.exe",
		},
		// buildpacks/pack
		{
			query:       "buildpacks/pack=v0.19.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-linux.tgz",
			execBinary:  "pack",
		},
		{
			query:       "buildpacks/pack=v0.19.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-macos.tgz",
			execBinary:  "pack",
		},
		{
			query:       "buildpacks/pack=v0.19.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/buildpacks/pack/releases/download/v0.19.0/pack-v0.19.0-windows.zip",
			execBinary:  "pack.exe",
		},
		// CircleCI-Public/circleci-cli
		{
			query:       "CircleCI-Public/circleci-cli=v0.1.8764",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_linux_amd64.tar.gz",
			execBinary:  "circleci",
		},
		{
			query:       "CircleCI-Public/circleci-cli=v0.1.8764",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_darwin_amd64.tar.gz",
			execBinary:  "circleci",
		},
		{
			query:       "CircleCI-Public/circleci-cli=v0.1.8764",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.8764/circleci-cli_0.1.8764_windows_amd64.zip",
			execBinary:  "circleci.exe",
		},
		// cli/cli
		{
			query:       "cli/cli=v1.12.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_linux_amd64.tar.gz",
			execBinary:  "gh",
		},
		{
			query:       "cli/cli=v1.12.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_macOS_amd64.tar.gz",
			execBinary:  "gh",
		},
		{
			query:       "cli/cli=v1.12.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/cli/cli/releases/download/v1.12.0/gh_1.12.0_windows_amd64.zip",
			execBinary:  "gh.exe",
		},
		// docker/buildx
		{
			query:       "docker/buildx=v0.10.4",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/docker/buildx/releases/download/v0.10.4/buildx-v0.10.4.linux-amd64",
			execBinary:  "docker-buildx",
		},
		{
			query:       "docker/buildx=v0.10.4",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/docker/buildx/releases/download/v0.10.4/buildx-v0.10.4.darwin-amd64",
			execBinary:  "docker-buildx",
		},
		{
			query:       "docker/buildx=v0.10.4",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/docker/buildx/releases/download/v0.10.4/buildx-v0.10.4.windows-amd64.exe",
			execBinary:  "docker-buildx.exe",
		},
		// docker/compose
		{
			query:       "docker/compose=1.25.4",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
			execBinary:  "docker-compose",
		},
		{
			query:       "docker/compose=1.25.4",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
			execBinary:  "docker-compose",
		},
		{
			query:       "docker/compose=1.25.4",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
			execBinary:  "docker-compose.exe",
		},
		// docker/machine
		{
			query:       "docker/machine=v0.16.2",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
			execBinary:  "docker-machine",
		},
		{
			query:       "docker/machine=v0.16.2",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
			execBinary:  "docker-machine",
		},
		{
			query:       "docker/machine=v0.16.2",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
			execBinary:  "docker-machine.exe",
		},
		// docker/scan-cli-plugin
		{
			query:       "docker/scan-cli-plugin=v0.17.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_linux_amd64",
			execBinary:  "docker-scan",
		},
		{
			query:       "docker/scan-cli-plugin=v0.17.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_darwin_amd64",
			execBinary:  "docker-scan",
		},
		{
			query:       "docker/scan-cli-plugin=v0.17.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_windows_amd64.exe",
			execBinary:  "docker-scan.exe",
		},
		// fluxcd/flux2
		{
			query:       "fluxcd/flux2=v0.8.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_linux_amd64.tar.gz",
			execBinary:  "flux",
		},
		{
			query:       "fluxcd/flux2=v0.8.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_darwin_amd64.tar.gz",
			execBinary:  "flux",
		},
		{
			query:       "fluxcd/flux2=v0.8.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_windows_amd64.zip",
			execBinary:  "flux.exe",
		},
		// goodwithtech/dockle
		{
			query:       "goodwithtech/dockle=v0.3.1",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Linux-64bit.tar.gz",
			execBinary:  "dockle",
		},
		{
			query:       "goodwithtech/dockle=v0.3.1",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_macOS-64bit.tar.gz",
			execBinary:  "dockle",
		},
		{
			query:       "goodwithtech/dockle=v0.3.1",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Windows-64bit.zip",
			execBinary:  "dockle.exe",
		},
		// gravitational/teleport
		{
			query:       "gravitational/teleport=v8.1.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://get.gravitational.com/teleport-v8.1.0-linux-amd64-bin.tar.gz",
			execBinary:  "tsh",
		},
		{
			query:       "gravitational/teleport=v8.1.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://get.gravitational.com/teleport-v8.1.0-darwin-amd64-bin.tar.gz",
			execBinary:  "tsh",
		},
		{
			query:       "gravitational/teleport=v8.1.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://get.gravitational.com/teleport-v8.1.0-windows-amd64-bin.zip",
			execBinary:  "tsh.exe",
		},
		// hashicorp/terraform
		{
			query:       "hashicorp/terraform=v0.12.20",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
			execBinary:  "terraform",
		},
		{
			query:       "hashicorp/terraform=v0.12.20",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
			execBinary:  "terraform",
		},
		{
			query:       "hashicorp/terraform=v0.12.20",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
			execBinary:  "terraform.exe",
		},
		// helm/helm
		{
			query:       "helm/helm=v3.1.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
			execBinary:  "helm",
		},
		{
			query:       "helm/helm=v3.1.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
			execBinary:  "helm",
		},
		{
			query:       "helm/helm=v3.1.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
			execBinary:  "helm.exe",
		},
		// istio/istio
		{
			query:       "istio/istio=1.6.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istio-1.6.0-linux-amd64.tar.gz",
			execBinary:  "istioctl",
		},
		{
			query:       "istio/istio=1.6.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istio-1.6.0-osx.tar.gz",
			execBinary:  "istioctl",
		},
		{
			query:       "istio/istio=1.6.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istio-1.6.0-win.zip",
			execBinary:  "istioctl.exe",
		},
		// mikefarah/yq
		{
			query:       "mikefarah/yq=v4.7.1",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_linux_amd64",
			execBinary:  "yq",
		},
		{
			query:       "mikefarah/yq=v4.7.1",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_darwin_amd64",
			execBinary:  "yq",
		},
		{
			query:       "mikefarah/yq=v4.7.1",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_windows_amd64.exe",
			execBinary:  "yq.exe",
		},
		// mozilla/sops
		{
			query:       "mozilla/sops=v3.5.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
			execBinary:  "sops",
		},
		{
			query:       "mozilla/sops=v3.5.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
			execBinary:  "sops",
		},
		{
			query:       "mozilla/sops=v3.5.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/mozilla/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
			execBinary:  "sops.exe",
		},
		// open-policy-agent/conftest
		{
			query:       "open-policy-agent/conftest=v0.21.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Linux_x86_64.tar.gz",
			execBinary:  "conftest",
		},
		{
			query:       "open-policy-agent/conftest=v0.21.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Darwin_x86_64.tar.gz",
			execBinary:  "conftest",
		},
		{
			query:       "open-policy-agent/conftest=v0.21.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Windows_x86_64.zip",
			execBinary:  "conftest.exe",
		},
		// open-policy-agent/opa
		{
			query:       "open-policy-agent/opa=v0.29.4",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_linux_amd64",
			execBinary:  "opa",
		},
		{
			query:       "open-policy-agent/opa=v0.29.4",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_darwin_amd64",
			execBinary:  "opa",
		},
		{
			query:       "open-policy-agent/opa=v0.29.4",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_windows_amd64.exe",
			execBinary:  "opa.exe",
		},
		// protocolbuffers/protobuf
		{
			query:       "protocolbuffers/protobuf=v3.11.4",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
			execBinary:  "protoc",
		},
		{
			query:       "protocolbuffers/protobuf=v3.11.4",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip",
			execBinary:  "protoc",
		},
		{
			query:       "protocolbuffers/protobuf=v3.11.4",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-win64.zip",
			execBinary:  "protoc.exe",
		},
		// starship/starship
		{
			query:       "starship/starship=v0.47.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-unknown-linux-gnu.tar.gz",
			execBinary:  "starship",
		},
		{
			query:       "starship/starship=v0.47.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-apple-darwin.tar.gz",
			execBinary:  "starship",
		},
		{
			query:       "starship/starship=v0.47.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-pc-windows-msvc.zip",
			execBinary:  "starship.exe",
		},
		// snyk/cli
		{
			query:       "snyk/cli=v1.1140.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/snyk/cli/releases/download/v1.1140.0/snyk-linux",
			execBinary:  "snyk",
		},
		{
			query:       "snyk/cli=v1.1140.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/snyk/cli/releases/download/v1.1140.0/snyk-macos",
			execBinary:  "snyk",
		},
		{
			query:       "snyk/cli=v1.1140.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/snyk/cli/releases/download/v1.1140.0/snyk-win.exe",
			execBinary:  "snyk.exe",
		},
		// viaduct-ai/kustomize-sops
		{
			query:       "viaduct-ai/kustomize-sops=v2.3.3",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Linux_x86_64.tar.gz",
			execBinary:  "ksops",
		},
		{
			query:       "viaduct-ai/kustomize-sops=v2.3.3",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Darwin_x86_64.tar.gz",
			execBinary:  "ksops",
		},
		{
			query:       "viaduct-ai/kustomize-sops=v2.3.3",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Windows_x86_64.tar.gz",
			execBinary:  "ksops.exe",
		},
	}

	for _, tt := range tests {
		name := tt.query
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			app := NewApplicationServiceForTest(ctx, t)
			query, err := ParseQuery(tt.query)
			assert.NoError(err)
			pkg, err := app.Search(ctx, query, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.downloadURL, pkg.Asset.DownloadURL)
			assert.Equal(tt.execBinary, pkg.ExecBinary.Name)
		})
	}
}

func TestParseQuery(t *testing.T) {
	tests := []struct {
		name     string
		queryStr string
		query    Query
	}{
		{
			name:     "shibataka000/go-get-release=v0.0.1",
			queryStr: "shibataka000/go-get-release=v0.0.1",
			query:    NewQuery(NewRepository("shibataka000", "go-get-release"), "v0.0.1"),
		},
		{
			name:     "go-get-release",
			queryStr: "go-get-release",
			query:    NewQuery(NewRepository("", "go-get-release"), ""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			query, err := ParseQuery(tt.queryStr)
			assert.NoError(err)
			assert.Equal(tt.query, query)
		})
	}
}

func TestQueryQueryToSearchGitHubRepository(t *testing.T) {
	tests := []struct {
		name                          string
		query                         Query
		queryToSearchGitHubRepository string
	}{
		{
			name:                          "shibataka000/go-get-release=v0.0.1",
			query:                         NewQuery(NewRepository("shibataka000", "go-get-release"), "v0.0.1"),
			queryToSearchGitHubRepository: "shibataka000/go-get-release",
		},
		{
			name:                          "go-get-release",
			query:                         NewQuery(NewRepository("", "go-get-release"), ""),
			queryToSearchGitHubRepository: "go-get-release",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.queryToSearchGitHubRepository, tt.query.QueryToSearchGitHubRepository())
		})
	}
}

func TestQueryIsSingleRepositorySpecified(t *testing.T) {
	tests := []struct {
		name                        string
		query                       Query
		isSingleRepositorySpecified bool
	}{
		{
			name:                        "shibataka000/go-get-release=v0.0.1",
			query:                       NewQuery(NewRepository("shibataka000", "go-get-release"), "v0.0.1"),
			isSingleRepositorySpecified: true,
		},
		{
			name:                        "go-get-release",
			query:                       NewQuery(NewRepository("", "go-get-release"), ""),
			isSingleRepositorySpecified: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.isSingleRepositorySpecified, tt.query.IsSingleRepositorySpecified())
		})
	}
}

func TestQueryHasTag(t *testing.T) {
	tests := []struct {
		name   string
		query  Query
		hasTag bool
	}{
		{
			name:   "shibataka000/go-get-release=v0.0.1",
			query:  NewQuery(NewRepository("shibataka000", "go-get-release"), "v0.0.1"),
			hasTag: true,
		},
		{
			name:   "go-get-release",
			query:  NewQuery(NewRepository("", "go-get-release"), ""),
			hasTag: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.hasTag, tt.query.HasTag())
		})
	}
}
