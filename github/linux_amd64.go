package github

var (
	DefaultAssetPatterns = []string{
		"trivy_0.53.0_Linux-64bit.tar.gz",
		"argocd-linux-amd64",
		"kubectl-argo-rollouts-linux-amd64",
		"argo-linux-amd64.gz",
		"pack-v0.34.2-linux.tgz",
		"gh_2.52.0_linux_amd64.tar.gz",
		"buildx-v0.15.1.linux-amd64",
		"docker-compose-linux-x86_64",
		"sops-v3.9.0.linux.amd64",
		"dockle_0.4.14_Linux-64bit.tar.gz",
		"istioctl-1.22.2-linux-amd64.tar.gz",
		"yq_linux_amd64.tar.gz",
		"conftest_0.53.0_Linux_x86_64.tar.gz",
		"gator-v3.16.3-linux-amd64.tar.gz",
		"opa_linux_amd64",
		"protoc-27.2-linux-x86_64.zip",
		"snyk-linux",
		"starship-x86_64-unknown-linux-gnu.tar.gz",
	}

	DefaultExecBinaryPatterns = []string{
		"trivy",
		"argocd",
		"kubectl-argo-rollouts",
		"argo",
		"pack",
		"gh",
		"docker-buildx",
		"docker-compose",
		"sops",
		"dockle",
		"istioctl",
		"yq",
		"conftest",
		"gator",
		"opa",
		"protoc",
		"snyk",
		"starship",
	}
)
