package github

import (
	"context"
	"testing"

	"github.com/shibataka000/go-get-release/mime"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/stretchr/testify/require"
)

func TestApplicationServiceSearch(t *testing.T) {
	tests := []struct {
		repo    Repository
		release Release
		os      platform.OS
		arch    platform.Arch
		asset   Asset
	}{
		// aquasecurity/tfsec
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.1.5"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-linux-amd64"), mime.OctedStream),
		},
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.1.5"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-darwin-amd64"), mime.OctedStream),
		},
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.1.5"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.1.5/tfsec-windows-amd64.exe"), mime.OctedStream),
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
		// kubernetes/kubernetes (for kubectl)
		{
			query:       "kubernetes/kubernetes=v1.28.2",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://dl.k8s.io/release/v1.28.2/bin/linux/amd64/kubectl",
			execBinary:  "kubectl",
		},
		{
			query:       "kubernetes/kubernetes=v1.28.2",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://dl.k8s.io/release/v1.28.2/bin/darwin/amd64/kubectl",
			execBinary:  "kubectl",
		},
		{
			query:       "kubernetes/kubernetes=v1.28.2",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://dl.k8s.io/release/v1.28.2/bin/windows/amd64/kubectl.exe",
			execBinary:  "kubectl.exe",
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
		// getsops/sops
		{
			query:       "getsops/sops=v3.5.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/getsops/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
			execBinary:  "sops",
		},
		{
			query:       "getsops/sops=v3.5.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/getsops/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
			execBinary:  "sops",
		},
		{
			query:       "getsops/sops=v3.5.0",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/getsops/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
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
		// open-policy-agent/gatekeeper
		{
			query:       "open-policy-agent/gatekeeper=v3.12.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/gatekeeper/releases/download/v3.12.0/gator-v3.12.0-linux-amd64.tar.gz",
			execBinary:  "gator",
		},
		{
			query:       "open-policy-agent/gatekeeper=v3.12.0",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/open-policy-agent/gatekeeper/releases/download/v3.12.0/gator-v3.12.0-darwin-amd64.tar.gz",
			execBinary:  "gator",
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
		// openshift-pipelines/pipelines-as-code
		{
			query:       "openshift-pipelines/pipelines-as-code=v0.21.1",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/openshift-pipelines/pipelines-as-code/releases/download/v0.21.1/tkn-pac_0.21.1_linux_x86_64.tar.gz",
			execBinary:  "tkn-pac",
		},
		{
			query:       "openshift-pipelines/pipelines-as-code=v0.21.1",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/openshift-pipelines/pipelines-as-code/releases/download/v0.21.1/tkn-pac_0.21.1_darwin_all.tar.gz",
			execBinary:  "tkn-pac",
		},
		{
			query:       "openshift-pipelines/pipelines-as-code=v0.21.1",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/openshift-pipelines/pipelines-as-code/releases/download/v0.21.1/tkn-pac_0.21.1_windows_x86_64.zip",
			execBinary:  "tkn-pac.exe",
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
			query:       "starship/starship=v1.16.0",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/starship/starship/releases/download/v1.16.0/starship-x86_64-unknown-linux-gnu.tar.gz",
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
		// tektoncd/cli
		{
			query:       "tektoncd/cli=v0.31.1",
			platform:    NewPlatform("linux", "amd64"),
			downloadURL: "https://github.com/tektoncd/cli/releases/download/v0.31.1/tkn_0.31.1_Linux_x86_64.tar.gz",
			execBinary:  "tkn",
		},
		{
			query:       "tektoncd/cli=v0.31.1",
			platform:    NewPlatform("darwin", "amd64"),
			downloadURL: "https://github.com/tektoncd/cli/releases/download/v0.31.1/tkn_0.31.1_Darwin_all.tar.gz",
			execBinary:  "tkn",
		},
		{
			query:       "tektoncd/cli=v0.31.1",
			platform:    NewPlatform("windows", "amd64"),
			downloadURL: "https://github.com/tektoncd/cli/releases/download/v0.31.1/tkn_0.31.1_Windows_x86_64.zip",
			execBinary:  "tkn.exe",
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
