package github

import (
	"context"
	"fmt"
	"os"
	"testing"

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
			release: newRelease("v1.28.6"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.28.6/tfsec-darwin-amd64")),
		},
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.28.6"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.28.6/tfsec-darwin-arm64")),
		},
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.28.6"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.28.6/tfsec-linux-amd64")),
		},
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.28.6"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.28.6/tfsec-linux-arm64")),
		},
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.28.6"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.28.6/tfsec-windows-amd64.exe")),
		},
		{
			repo:    newRepository("aquasecurity", "tfsec"),
			release: newRelease("v1.28.6"),
			os:      "windows",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/tfsec/releases/download/v1.28.6/tfsec-windows-arm64.exe")),
		},
		// aquasecurity/trivy
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "freebsd",
			arch:    "amd32",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_FreeBSD-32bit.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "freebsd",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_FreeBSD-64bit.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "linux",
			arch:    "amd32",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_Linux-32bit.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_Linux-64bit.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "linux",
			arch:    "arm32",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_Linux-ARM.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_Linux-ARM64.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "linux",
			arch:    "ppc64le",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_Linux-PPC64LE.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "linux",
			arch:    "s390x",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_Linux-s390x.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_macOS-64bit.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_macOS-ARM64.tar.gz")),
		},
		{
			repo:    newRepository("aquasecurity", "trivy"),
			release: newRelease("v0.52.2"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_windows-64bit.zip")),
		},
		// argoproj/argo-cd
		{
			repo:    newRepository("argoproj", "argocd"),
			release: newRelease("v2.11.3"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-cd/releases/download/v2.11.3/argocd-darwin-amd64")),
		},
		{
			repo:    newRepository("argoproj", "argocd"),
			release: newRelease("v2.11.3"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-cd/releases/download/v2.11.3/argocd-darwin-arm64")),
		},
		{
			repo:    newRepository("argoproj", "argocd"),
			release: newRelease("v2.11.3"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-cd/releases/download/v2.11.3/argocd-linux-amd64")),
		},
		{
			repo:    newRepository("argoproj", "argocd"),
			release: newRelease("v2.11.3"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-cd/releases/download/v2.11.3/argocd-linux-arm64")),
		},
		{
			repo:    newRepository("argoproj", "argocd"),
			release: newRelease("v2.11.3"),
			os:      "linux",
			arch:    "ppc64le",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-cd/releases/download/v2.11.3/argocd-linux-ppc64le")),
		},
		{
			repo:    newRepository("argoproj", "argocd"),
			release: newRelease("v2.11.3"),
			os:      "linux",
			arch:    "s390x",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-cd/releases/download/v2.11.3/argocd-linux-s390x")),
		},
		{
			repo:    newRepository("argoproj", "argocd"),
			release: newRelease("v2.11.3"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-cd/releases/download/v2.11.3/argocd-windows-amd64.exe")),
		},
		// argoproj/argo-rollouts
		{
			repo:    newRepository("argoproj", "argo-rollouts"),
			release: newRelease("v1.7.0"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-darwin-amd64")),
		},
		{
			repo:    newRepository("argoproj", "argo-rollouts"),
			release: newRelease("v1.7.0"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-darwin-arm64")),
		},
		{
			repo:    newRepository("argoproj", "argo-rollouts"),
			release: newRelease("v1.7.0"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-linux-amd64")),
		},
		{
			repo:    newRepository("argoproj", "argo-rollouts"),
			release: newRelease("v1.7.0"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-linux-arm64")),
		},
		{
			repo:    newRepository("argoproj", "argo-rollouts"),
			release: newRelease("v1.7.0"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-rollouts/releases/download/v0.9.0/kubectl-argo-rollouts-windows-amd64")),
		},
		// argoproj/argo-workflows
		{
			repo:    newRepository("argoproj", "argo-workflows"),
			release: newRelease("v3.5.8"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-darwin-amd64.gz")),
		},
		{
			repo:    newRepository("argoproj", "argo-workflows"),
			release: newRelease("v3.5.8"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-darwin-arm64.gz")),
		},
		{
			repo:    newRepository("argoproj", "argo-workflows"),
			release: newRelease("v3.5.8"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-linux-amd64.gz")),
		},
		{
			repo:    newRepository("argoproj", "argo-workflows"),
			release: newRelease("v3.5.8"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-linux-arm64.gz")),
		},
		{
			repo:    newRepository("argoproj", "argo-workflows"),
			release: newRelease("v3.5.8"),
			os:      "linux",
			arch:    "ppc64le",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-linux-ppc64le.gz")),
		},
		{
			repo:    newRepository("argoproj", "argo-workflows"),
			release: newRelease("v3.5.8"),
			os:      "linux",
			arch:    "s390x",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-linux-s390x.gz")),
		},
		{
			repo:    newRepository("argoproj", "argo-workflows"),
			release: newRelease("v3.5.8"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/argoproj/argo-workflows/releases/download/v3.5.8/argo-windows-amd64.exe.gz")),
		},
		// aws/amazon-ec2-instance-selector
		{
			repo:    newRepository("aws", "amazon-ec2-instance-selector"),
			release: newRelease("v2.4.1"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.4.1/ec2-instance-selector-darwin-amd64")),
		},
		{
			repo:    newRepository("aws", "amazon-ec2-instance-selector"),
			release: newRelease("v2.4.1"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.4.1/ec2-instance-selector-darwin-arm64")),
		},
		{
			repo:    newRepository("aws", "amazon-ec2-instance-selector"),
			release: newRelease("v2.4.1"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.4.1/ec2-instance-selector-linux-amd64")),
		},
		{
			repo:    newRepository("aws", "amazon-ec2-instance-selector"),
			release: newRelease("v2.4.1"),
			os:      "linux",
			arch:    "arm32",
			asset:   newAsset(newURL("https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.4.1/ec2-instance-selector-linux-arm")),
		},
		{
			repo:    newRepository("aws", "amazon-ec2-instance-selector"),
			release: newRelease("v2.4.1"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.4.1/ec2-instance-selector-linux-arm64")),
		},
		{
			repo:    newRepository("aws", "amazon-ec2-instance-selector"),
			release: newRelease("v2.4.1"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/aws/amazon-ec2-instance-selector/releases/download/v2.4.1/ec2-instance-selector-windows-amd64")),
		},
		// bitnami-labs/sealed-secrets
		{
			repo:    newRepository("bitnami-labs", "sealed-secrets"),
			release: newRelease("v0.27.0"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.27.0/kubeseal-0.27.0-darwin-amd64.tar.gz")),
		},
		{
			repo:    newRepository("bitnami-labs", "sealed-secrets"),
			release: newRelease("v0.27.0"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.27.0/kubeseal-0.27.0-darwin-arm64.tar.gz")),
		},
		{
			repo:    newRepository("bitnami-labs", "sealed-secrets"),
			release: newRelease("v0.27.0"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.27.0/kubeseal-0.27.0-linux-amd64.tar.gz")),
		},
		{
			repo:    newRepository("bitnami-labs", "sealed-secrets"),
			release: newRelease("v0.27.0"),
			os:      "linux",
			arch:    "arm32",
			asset:   newAsset(newURL("https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.27.0/kubeseal-0.27.0-linux-arm.tar.gz")),
		},
		{
			repo:    newRepository("bitnami-labs", "sealed-secrets"),
			release: newRelease("v0.27.0"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.27.0/kubeseal-0.27.0-linux-arm64.tar.gz")),
		},
		{
			repo:    newRepository("bitnami-labs", "sealed-secrets"),
			release: newRelease("v0.27.0"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.27.0/kubeseal-0.27.0-windows-amd64.tar.gz")),
		},
		// buildpacks/pack
		{
			repo:    newRepository("buikdpacks", "pack"),
			release: newRelease("v0.34.2"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-linux-arm64.tgz")),
		},
		{
			repo:    newRepository("buikdpacks", "pack"),
			release: newRelease("v0.34.2"),
			os:      "linux",
			arch:    "ppc64le",
			asset:   newAsset(newURL("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-linux-ppc64le.tgz")),
		},
		{
			repo:    newRepository("buikdpacks", "pack"),
			release: newRelease("v0.34.2"),
			os:      "linux",
			arch:    "s390x",
			asset:   newAsset(newURL("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-linux-s390x.tgz")),
		},
		{
			repo:    newRepository("buikdpacks", "pack"),
			release: newRelease("v0.34.2"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-linux.tgz")),
		},
		{
			repo:    newRepository("buikdpacks", "pack"),
			release: newRelease("v0.34.2"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-macos-arm64.tgz")),
		},
		{
			repo:    newRepository("buikdpacks", "pack"),
			release: newRelease("v0.34.2"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-macos.tgz")),
		},
		{
			repo:    newRepository("buikdpacks", "pack"),
			release: newRelease("v0.34.2"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/buildpacks/pack/releases/download/v0.34.2/pack-v0.34.2-windows.zip")),
		},
		// CircleCI-Public/circleci-cli
		{
			repo:    newRepository("CircleCI-Public", "circleci-cli"),
			release: newRelease("v0.1.30549"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.30549/circleci-cli_0.1.30549_darwin_amd64.tar.gz")),
		},
		{
			repo:    newRepository("CircleCI-Public", "circleci-cli"),
			release: newRelease("v0.1.30549"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.30549/circleci-cli_0.1.30549_darwin_arm64.tar.gz")),
		},
		{
			repo:    newRepository("CircleCI-Public", "circleci-cli"),
			release: newRelease("v0.1.30549"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.30549/circleci-cli_0.1.30549_linux_amd64.tar.gz")),
		},
		{
			repo:    newRepository("CircleCI-Public", "circleci-cli"),
			release: newRelease("v0.1.30549"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.30549/circleci-cli_0.1.30549_linux_arm64.tar.gz")),
		},
		{
			repo:    newRepository("CircleCI-Public", "circleci-cli"),
			release: newRelease("v0.1.30549"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.30549/circleci-cli_0.1.30549_windows_amd64.zip")),
		},
		{
			repo:    newRepository("CircleCI-Public", "circleci-cli"),
			release: newRelease("v0.1.30549"),
			os:      "windows",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/CircleCI-Public/circleci-cli/releases/download/v0.1.30549/circleci-cli_0.1.30549_windows_arm64.zip")),
		},
		// cli/cli
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "linux",
			arch:    "amd32",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_linux_386.tar.gz")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "linux",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_linux_amd64.tar.gz")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_linux_arm64.tar.gz")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "linux",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_linux_armv6.tar.gz")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "darwin",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_macOS_amd64.zip")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "darwin",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_macOS_arm64.zip")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "windows",
			arch:    "amd32",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_windows_386.zip")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "windows",
			arch:    "amd64",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_linux_amd64.zip")),
		},
		{
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.51.0"),
			os:      "windows",
			arch:    "arm64",
			asset:   newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.0/gh_2.51.0_linux_arm64.zip")),
		},
		// // docker/buildx
		// {
		// 	query:       "docker/buildx=v0.10.4",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/docker/buildx/releases/download/v0.10.4/buildx-v0.10.4.linux-amd64",
		// 	execBinary:  "docker-buildx",
		// },
		// {
		// 	query:       "docker/buildx=v0.10.4",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/docker/buildx/releases/download/v0.10.4/buildx-v0.10.4.darwin-amd64",
		// 	execBinary:  "docker-buildx",
		// },
		// {
		// 	query:       "docker/buildx=v0.10.4",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/docker/buildx/releases/download/v0.10.4/buildx-v0.10.4.windows-amd64.exe",
		// 	execBinary:  "docker-buildx.exe",
		// },
		// // docker/compose
		// {
		// 	query:       "docker/compose=1.25.4",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Linux-x86_64",
		// 	execBinary:  "docker-compose",
		// },
		// {
		// 	query:       "docker/compose=1.25.4",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Darwin-x86_64",
		// 	execBinary:  "docker-compose",
		// },
		// {
		// 	query:       "docker/compose=1.25.4",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-Windows-x86_64.exe",
		// 	execBinary:  "docker-compose.exe",
		// },
		// // docker/machine
		// {
		// 	query:       "docker/machine=v0.16.2",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Linux-x86_64",
		// 	execBinary:  "docker-machine",
		// },
		// {
		// 	query:       "docker/machine=v0.16.2",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Darwin-x86_64",
		// 	execBinary:  "docker-machine",
		// },
		// {
		// 	query:       "docker/machine=v0.16.2",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/docker/machine/releases/download/v0.16.2/docker-machine-Windows-x86_64.exe",
		// 	execBinary:  "docker-machine.exe",
		// },
		// // docker/scan-cli-plugin
		// {
		// 	query:       "docker/scan-cli-plugin=v0.17.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_linux_amd64",
		// 	execBinary:  "docker-scan",
		// },
		// {
		// 	query:       "docker/scan-cli-plugin=v0.17.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_darwin_amd64",
		// 	execBinary:  "docker-scan",
		// },
		// {
		// 	query:       "docker/scan-cli-plugin=v0.17.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/docker/scan-cli-plugin/releases/download/v0.17.0/docker-scan_windows_amd64.exe",
		// 	execBinary:  "docker-scan.exe",
		// },
		// // fluxcd/flux2
		// {
		// 	query:       "fluxcd/flux2=v0.8.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_linux_amd64.tar.gz",
		// 	execBinary:  "flux",
		// },
		// {
		// 	query:       "fluxcd/flux2=v0.8.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_darwin_amd64.tar.gz",
		// 	execBinary:  "flux",
		// },
		// {
		// 	query:       "fluxcd/flux2=v0.8.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/fluxcd/flux2/releases/download/v0.8.0/flux_0.8.0_windows_amd64.zip",
		// 	execBinary:  "flux.exe",
		// },
		// // goodwithtech/dockle
		// {
		// 	query:       "goodwithtech/dockle=v0.3.1",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Linux-64bit.tar.gz",
		// 	execBinary:  "dockle",
		// },
		// {
		// 	query:       "goodwithtech/dockle=v0.3.1",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_macOS-64bit.tar.gz",
		// 	execBinary:  "dockle",
		// },
		// {
		// 	query:       "goodwithtech/dockle=v0.3.1",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/goodwithtech/dockle/releases/download/v0.3.1/dockle_0.3.1_Windows-64bit.zip",
		// 	execBinary:  "dockle.exe",
		// },
		// // gravitational/teleport
		// {
		// 	query:       "gravitational/teleport=v8.1.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://get.gravitational.com/teleport-v8.1.0-linux-amd64-bin.tar.gz",
		// 	execBinary:  "tsh",
		// },
		// {
		// 	query:       "gravitational/teleport=v8.1.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://get.gravitational.com/teleport-v8.1.0-darwin-amd64-bin.tar.gz",
		// 	execBinary:  "tsh",
		// },
		// {
		// 	query:       "gravitational/teleport=v8.1.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://get.gravitational.com/teleport-v8.1.0-windows-amd64-bin.zip",
		// 	execBinary:  "tsh.exe",
		// },
		// // hashicorp/terraform
		// {
		// 	query:       "hashicorp/terraform=v0.12.20",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
		// 	execBinary:  "terraform",
		// },
		// {
		// 	query:       "hashicorp/terraform=v0.12.20",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
		// 	execBinary:  "terraform",
		// },
		// {
		// 	query:       "hashicorp/terraform=v0.12.20",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
		// 	execBinary:  "terraform.exe",
		// },
		// // helm/helm
		// {
		// 	query:       "helm/helm=v3.1.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
		// 	execBinary:  "helm",
		// },
		// {
		// 	query:       "helm/helm=v3.1.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
		// 	execBinary:  "helm",
		// },
		// {
		// 	query:       "helm/helm=v3.1.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
		// 	execBinary:  "helm.exe",
		// },
		// // istio/istio
		// {
		// 	query:       "istio/istio=1.6.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istio-1.6.0-linux-amd64.tar.gz",
		// 	execBinary:  "istioctl",
		// },
		// {
		// 	query:       "istio/istio=1.6.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istio-1.6.0-osx.tar.gz",
		// 	execBinary:  "istioctl",
		// },
		// {
		// 	query:       "istio/istio=1.6.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/istio/istio/releases/download/1.6.0/istio-1.6.0-win.zip",
		// 	execBinary:  "istioctl.exe",
		// },
		// // kubernetes/kubernetes (for kubectl)
		// {
		// 	query:       "kubernetes/kubernetes=v1.28.2",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://dl.k8s.io/release/v1.28.2/bin/linux/amd64/kubectl",
		// 	execBinary:  "kubectl",
		// },
		// {
		// 	query:       "kubernetes/kubernetes=v1.28.2",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://dl.k8s.io/release/v1.28.2/bin/darwin/amd64/kubectl",
		// 	execBinary:  "kubectl",
		// },
		// {
		// 	query:       "kubernetes/kubernetes=v1.28.2",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://dl.k8s.io/release/v1.28.2/bin/windows/amd64/kubectl.exe",
		// 	execBinary:  "kubectl.exe",
		// },
		// // mikefarah/yq
		// {
		// 	query:       "mikefarah/yq=v4.7.1",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_linux_amd64",
		// 	execBinary:  "yq",
		// },
		// {
		// 	query:       "mikefarah/yq=v4.7.1",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_darwin_amd64",
		// 	execBinary:  "yq",
		// },
		// {
		// 	query:       "mikefarah/yq=v4.7.1",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_windows_amd64.exe",
		// 	execBinary:  "yq.exe",
		// },
		// // getsops/sops
		// {
		// 	query:       "getsops/sops=v3.5.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/getsops/sops/releases/download/v3.5.0/sops-v3.5.0.linux",
		// 	execBinary:  "sops",
		// },
		// {
		// 	query:       "getsops/sops=v3.5.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/getsops/sops/releases/download/v3.5.0/sops-v3.5.0.darwin",
		// 	execBinary:  "sops",
		// },
		// {
		// 	query:       "getsops/sops=v3.5.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/getsops/sops/releases/download/v3.5.0/sops-v3.5.0.exe",
		// 	execBinary:  "sops.exe",
		// },
		// // open-policy-agent/conftest
		// {
		// 	query:       "open-policy-agent/conftest=v0.21.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Linux_x86_64.tar.gz",
		// 	execBinary:  "conftest",
		// },
		// {
		// 	query:       "open-policy-agent/conftest=v0.21.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Darwin_x86_64.tar.gz",
		// 	execBinary:  "conftest",
		// },
		// {
		// 	query:       "open-policy-agent/conftest=v0.21.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/conftest/releases/download/v0.21.0/conftest_0.21.0_Windows_x86_64.zip",
		// 	execBinary:  "conftest.exe",
		// },
		// // open-policy-agent/gatekeeper
		// {
		// 	query:       "open-policy-agent/gatekeeper=v3.12.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/gatekeeper/releases/download/v3.12.0/gator-v3.12.0-linux-amd64.tar.gz",
		// 	execBinary:  "gator",
		// },
		// {
		// 	query:       "open-policy-agent/gatekeeper=v3.12.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/gatekeeper/releases/download/v3.12.0/gator-v3.12.0-darwin-amd64.tar.gz",
		// 	execBinary:  "gator",
		// },
		// // open-policy-agent/opa
		// {
		// 	query:       "open-policy-agent/opa=v0.29.4",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_linux_amd64",
		// 	execBinary:  "opa",
		// },
		// {
		// 	query:       "open-policy-agent/opa=v0.29.4",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_darwin_amd64",
		// 	execBinary:  "opa",
		// },
		// {
		// 	query:       "open-policy-agent/opa=v0.29.4",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/open-policy-agent/opa/releases/download/v0.29.4/opa_windows_amd64.exe",
		// 	execBinary:  "opa.exe",
		// },
		// // openshift-pipelines/pipelines-as-code
		// {
		// 	query:       "openshift-pipelines/pipelines-as-code=v0.21.1",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/openshift-pipelines/pipelines-as-code/releases/download/v0.21.1/tkn-pac_0.21.1_linux_x86_64.tar.gz",
		// 	execBinary:  "tkn-pac",
		// },
		// {
		// 	query:       "openshift-pipelines/pipelines-as-code=v0.21.1",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/openshift-pipelines/pipelines-as-code/releases/download/v0.21.1/tkn-pac_0.21.1_darwin_all.tar.gz",
		// 	execBinary:  "tkn-pac",
		// },
		// {
		// 	query:       "openshift-pipelines/pipelines-as-code=v0.21.1",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/openshift-pipelines/pipelines-as-code/releases/download/v0.21.1/tkn-pac_0.21.1_windows_x86_64.zip",
		// 	execBinary:  "tkn-pac.exe",
		// },
		// // protocolbuffers/protobuf
		// {
		// 	query:       "protocolbuffers/protobuf=v3.11.4",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip",
		// 	execBinary:  "protoc",
		// },
		// {
		// 	query:       "protocolbuffers/protobuf=v3.11.4",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip",
		// 	execBinary:  "protoc",
		// },
		// {
		// 	query:       "protocolbuffers/protobuf=v3.11.4",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-win64.zip",
		// 	execBinary:  "protoc.exe",
		// },
		// // starship/starship
		// {
		// 	query:       "starship/starship=v0.47.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-unknown-linux-gnu.tar.gz",
		// 	execBinary:  "starship",
		// },
		// {
		// 	query:       "starship/starship=v1.16.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/starship/starship/releases/download/v1.16.0/starship-x86_64-unknown-linux-gnu.tar.gz",
		// 	execBinary:  "starship",
		// },
		// {
		// 	query:       "starship/starship=v0.47.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-apple-darwin.tar.gz",
		// 	execBinary:  "starship",
		// },
		// {
		// 	query:       "starship/starship=v0.47.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/starship/starship/releases/download/v0.47.0/starship-x86_64-pc-windows-msvc.zip",
		// 	execBinary:  "starship.exe",
		// },
		// // snyk/cli
		// {
		// 	query:       "snyk/cli=v1.1140.0",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/snyk/cli/releases/download/v1.1140.0/snyk-linux",
		// 	execBinary:  "snyk",
		// },
		// {
		// 	query:       "snyk/cli=v1.1140.0",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/snyk/cli/releases/download/v1.1140.0/snyk-macos",
		// 	execBinary:  "snyk",
		// },
		// {
		// 	query:       "snyk/cli=v1.1140.0",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/snyk/cli/releases/download/v1.1140.0/snyk-win.exe",
		// 	execBinary:  "snyk.exe",
		// },
		// // tektoncd/cli
		// {
		// 	query:       "tektoncd/cli=v0.31.1",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/tektoncd/cli/releases/download/v0.31.1/tkn_0.31.1_Linux_x86_64.tar.gz",
		// 	execBinary:  "tkn",
		// },
		// {
		// 	query:       "tektoncd/cli=v0.31.1",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/tektoncd/cli/releases/download/v0.31.1/tkn_0.31.1_Darwin_all.tar.gz",
		// 	execBinary:  "tkn",
		// },
		// {
		// 	query:       "tektoncd/cli=v0.31.1",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/tektoncd/cli/releases/download/v0.31.1/tkn_0.31.1_Windows_x86_64.zip",
		// 	execBinary:  "tkn.exe",
		// },
		// // viaduct-ai/kustomize-sops
		// {
		// 	query:       "viaduct-ai/kustomize-sops=v2.3.3",
		// 	platform:    NewPlatform("linux", "amd64"),
		// 	downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Linux_x86_64.tar.gz",
		// 	execBinary:  "ksops",
		// },
		// {
		// 	query:       "viaduct-ai/kustomize-sops=v2.3.3",
		// 	platform:    NewPlatform("darwin", "amd64"),
		// 	downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Darwin_x86_64.tar.gz",
		// 	execBinary:  "ksops",
		// },
		// {
		// 	query:       "viaduct-ai/kustomize-sops=v2.3.3",
		// 	platform:    NewPlatform("windows", "amd64"),
		// 	downloadURL: "https://github.com/viaduct-ai/kustomize-sops/releases/download/v2.3.3/ksops_2.3.3_Windows_x86_64.tar.gz",
		// 	execBinary:  "ksops.exe",
		// },
	}

	for _, tt := range tests {
		name := tt.asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			app := NewApplicationService(
				NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN")),
			)
			asset, err := app.FindAsset(ctx, fmt.Sprintf("%s/%s", tt.repo.owner, tt.repo.name), tt.release.tag, tt.os, tt.arch)
			require.NoError(err)
			require.Equal(tt.asset, asset)
		})
	}
}
