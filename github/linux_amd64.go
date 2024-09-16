package github

var defaultPatternsForLinuxAmd64 = map[string]string{
	"trivy_0.53.0_Linux-64bit.tar.gz":          "trivy",
	"argocd-linux-amd64":                       "argocd",
	"kubectl-argo-rollouts-linux-amd64":        "kubectl-argo-rollouts",
	"argo-linux-amd64.gz":                      "argo",
	"pack-v0.34.2-linux.tgz":                   "pack",
	"gh_2.52.0_linux_amd64.tar.gz":             "gh",
	"buildx-v0.15.1.linux-amd64":               "docker-buildx",
	"docker-compose-linux-x86_64":              "docker-compose",
	"sops-v3.9.0.linux.amd64":                  "sops",
	"dockle_0.4.14_Linux-64bit.tar.gz":         "dockle",
	"istioctl-1.22.2-linux-amd64.tar.gz":       "istioctl",
	"yq_linux_amd64.tar.gz":                    "yq",
	"conftest_0.53.0_Linux_x86_64.tar.gz":      "conftest",
	"gator-v3.16.3-linux-amd64.tar.gz":         "gator",
	"opa_linux_amd64":                          "opa",
	"protoc-27.2-linux-x86_64.zip":             "protoc",
	"snyk-linux":                               "snyk",
	"starship-x86_64-unknown-linux-gnu.tar.gz": "starship",
}

func init() {
	DefaultPatterns = defaultPatternsForLinuxAmd64
}
