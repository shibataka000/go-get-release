package github

var defaultPatternsForLinuxAmd64 = PatternList{
	mustNewPatternFromString("trivy_0.53.0_Linux-64bit.tar.gz", "trivy"),
	mustNewPatternFromString("argocd-linux-amd64", "argocd"),
	mustNewPatternFromString("kubectl-argo-rollouts-linux-amd64", "kubectl-argo-rollouts"),
	mustNewPatternFromString("argo-linux-amd64.gz", "argo"),
	mustNewPatternFromString("pack-v0.34.2-linux.tgz", "pack"),
	mustNewPatternFromString("gh_2.52.0_linux_amd64.tar.gz", "gh"),
	mustNewPatternFromString("buildx-v0.15.1.linux-amd64", "docker-buildx"),
	mustNewPatternFromString("docker-compose-linux-x86_64", "docker-compose"),
	mustNewPatternFromString("sops-v3.9.0.linux.amd64", "sops"),
	mustNewPatternFromString("dockle_0.4.14_Linux-64bit.tar.gz", "dockle"),
	mustNewPatternFromString("istioctl-1.22.2-linux-amd64.tar.gz", "istioctl"),
	mustNewPatternFromString("yq_linux_amd64.tar.gz", "yq"),
	mustNewPatternFromString("conftest_0.53.0_Linux_x86_64.tar.gz", "conftest"),
	mustNewPatternFromString("gator-v3.16.3-linux-amd64.tar.gz", "gator"),
	mustNewPatternFromString("opa_linux_amd64", "opa"),
	mustNewPatternFromString("protoc-27.2-linux-x86_64.zip", "protoc"),
	mustNewPatternFromString("snyk-linux", "snyk"),
	mustNewPatternFromString("starship-x86_64-unknown-linux-gnu.tar.gz", "starship"),
}

func init() {
	for _, p := range defaultPatternsForLinuxAmd64 {
		DefaultAssetPatterns = append(DefaultAssetPatterns, p.asset.String())
		DefaultExecBinaryPatterns = append(DefaultExecBinaryPatterns, p.execBinary.Name())
	}
}
