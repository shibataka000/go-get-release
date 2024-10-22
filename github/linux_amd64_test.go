package github

import (
	"context"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAndInstallOnLinuxAmd64(t *testing.T) {
	tests := []struct {
		repoFullName string
		tag          string

		asset      Asset
		execBinary ExecBinary

		test *exec.Cmd
	}{
		{
			repoFullName: "aquasecurity/trivy",
			tag:          "v0.53.0",
			asset:        must(newAssetFromString(176981043, "trivy_0.53.0_Linux-64bit.tar.gz")),
			execBinary:   newExecBinary("trivy"),
			test:         exec.Command("./trivy", "version"),
		},
		{
			repoFullName: "argoproj/argo-cd",
			tag:          "v2.9.18",
			asset:        must(newAssetFromString(177293568, "argocd-linux-amd64")),
			execBinary:   newExecBinary("argocd"),
			test:         exec.Command("./argocd", "version"),
		},
		{
			repoFullName: "argoproj/argo-rollouts",
			tag:          "v1.7.1",
			asset:        must(newAssetFromString(175717897, "kubectl-argo-rollouts-linux-amd64")),
			execBinary:   newExecBinary("kubectl-argo-rollouts"),
			test:         exec.Command("./kubectl-argo-rollouts"),
		},
		{
			repoFullName: "argoproj/argo-workflows",
			tag:          "v3.5.8",
			asset:        must(newAssetFromString(174415137, "argo-linux-amd64.gz")),
			execBinary:   newExecBinary("argo"),
			test:         exec.Command("./argo"),
		},
		{
			repoFullName: "buildpacks/pack",
			tag:          "v0.34.2",
			asset:        must(newAssetFromString(172104571, "pack-v0.34.2-linux.tgz")),
			execBinary:   newExecBinary("pack"),
			test:         exec.Command("./pack"),
		},
		{
			repoFullName: "cli/cli",
			tag:          "v2.52.0",
			asset:        must(newAssetFromString(175682889, "gh_2.52.0_linux_amd64.tar.gz")),
			execBinary:   newExecBinary("gh"),
			test:         exec.Command("./gh"),
		},
		{
			repoFullName: "docker/buildx",
			tag:          "v0.15.1",
			asset:        must(newAssetFromString(174531956, "buildx-v0.15.1.linux-amd64")),
			execBinary:   newExecBinary("docker-buildx"),
			test:         exec.Command("./docker-buildx"),
		},
		{
			repoFullName: "docker/compose",
			tag:          "v2.28.1",
			asset:        must(newAssetFromString(175627020, "docker-compose-linux-x86_64")),
			execBinary:   newExecBinary("docker-compose"),
			test:         exec.Command("./docker-compose"),
		},
		{
			repoFullName: "getsops/sops",
			tag:          "v3.9.0",
			asset:        must(newAssetFromString(176438234, "sops-v3.9.0.linux.amd64")),
			execBinary:   newExecBinary("sops"),
			test:         exec.Command("./sops"),
		},
		{
			repoFullName: "goodwithtech/dockle",
			tag:          "v0.4.14",
			asset:        must(newAssetFromString(149683239, "dockle_0.4.14_Linux-64bit.tar.gz")),
			execBinary:   newExecBinary("dockle"),
			test:         exec.Command("./dockle"),
		},
		{
			repoFullName: "istio/istio",
			tag:          "v1.22.2",
			asset:        must(newAssetFromString(174040551, "istioctl-1.22.2-linux-amd64.tar.gz")),
			execBinary:   newExecBinary("istioctl"),
			test:         exec.Command("./istioctl"),
		},
		{
			repoFullName: "mikefarah/yq",
			tag:          "v4.44.2",
			asset:        must(newAssetFromString(174040551, "yq_linux_amd64.tar.gz")),
			execBinary:   newExecBinary("yq"),
			test:         exec.Command("./yq"),
		},
		{
			repoFullName: "open-policy-agent/conftest",
			tag:          "v0.53.0",
			asset:        must(newAssetFromString(172540735, "conftest_0.53.0_Linux_x86_64.tar.gz")),
			execBinary:   newExecBinary("conftest"),
			test:         exec.Command("./conftest"),
		},
		{
			repoFullName: "open-policy-agent/gatekeeper",
			tag:          "v3.16.3",
			asset:        must(newAssetFromString(169950399, "gator-v3.16.3-linux-amd64.tar.gz")),
			execBinary:   newExecBinary("gator"),
			test:         exec.Command("./gator"),
		},
		{
			repoFullName: "open-policy-agent/opa",
			tag:          "v0.66.0",
			asset:        must(newAssetFromString(176292835, "opa_linux_amd64")),
			execBinary:   newExecBinary("opa"),
			test:         exec.Command("./opa"),
		},
		{
			repoFullName: "protocolbuffers/protobuf",
			tag:          "v27.2",
			asset:        must(newAssetFromString(175919234, "protoc-27.2-linux-x86_64.zip")),
			execBinary:   newExecBinary("protoc"),
			test:         exec.Command("protoc"),
		},
		{
			repoFullName: "snyk/cli",
			tag:          "v1.1292.1",
			asset:        must(newAssetFromString(176276540, "snyk-linux")),
			execBinary:   newExecBinary("snyk"),
			test:         exec.Command("snyk"),
		},
		{
			repoFullName: "starship/starship",
			tag:          "v1.19.0",
			asset:        must(newAssetFromString(168103285, "starship-x86_64-unknown-linux-gnu.tar.gz")),
			execBinary:   newExecBinary("starship"),
			test:         exec.Command("./starship"),
		},
		{
			repoFullName: "viaduct-ai/kustomize-sops",
			tag:          "v4.3.2",
			asset:        must(newAssetFromString(0, "ksops_4.3.2_Linux_x86_64.tar.gz")),
			execBinary:   newExecBinary("ksops"),
			test:         exec.Command("./ksops"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.repoFullName, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()

			dir, err := os.MkdirTemp("", "")
			require.NoError(err)
			defer os.RemoveAll(dir)

			tt.test.Dir = dir

			require.Error(tt.test.Run(), "executable binary was already installed")

			app := NewApplicationService(
				NewAssetRepository(ctx, githubTokenForTest),
				NewExecBinaryRepository(),
			)

			asset, execBinary, err := app.Find(ctx, tt.repoFullName, tt.tag, DefaultPatterns)
			require.NoError(err)
			require.Equal(tt.asset, asset)
			require.Equal(tt.execBinary, execBinary)

			err = app.Install(ctx, tt.repoFullName, asset, execBinary, dir, io.Discard)
			require.NoError(err)

			require.NoError(tt.test.Run())

		})
	}

}
