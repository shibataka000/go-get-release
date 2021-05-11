package github

var specialAssetMap = map[string]map[string]*asset{
	"docker/compose": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/docker-compose-Linux-x86_64", binaryName: "docker-compose"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/docker-compose-Darwin-x86_64", binaryName: "docker-compose"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/docker-compose-Windows-x86_64.exe", binaryName: "docker-compose.exe"},
	},
	"docker/machine": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/docker-machine-Linux-x86_64", binaryName: "docker-machine"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/docker-machine-Darwin-x86_64", binaryName: "docker-machine"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/docker-machine-Windows-x86_64.exe", binaryName: "docker-machine.exe"},
	},
	"helm/helm": {
		"default":       &asset{downloadURL: "https://get.helm.sh/helm-{{.Tag}}-{{.Goos}}-{{.Goarch}}.tar.gz", binaryName: "helm"},
		"windows/amd64": &asset{downloadURL: "https://get.helm.sh/helm-{{.Tag}}-{{.Goos}}-{{.Goarch}}.zip", binaryName: "helm.exe"},
	},
	"istio/istio": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/istioctl-{{.Tag}}-{{.Goos}}-{{.Goarch}}.tar.gz", binaryName: "istioctl"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/istioctl-{{.Tag}}-osx.tar.gz", binaryName: "istioctl"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/istioctl-{{.Tag}}-win.zip", binaryName: "istioctl.exe"},
	},
	"hashicorp/terraform": {
		"default":       &asset{downloadURL: "https://releases.hashicorp.com/terraform/{{.Version}}/terraform_{{.Version}}_{{.Goos}}_{{.Goarch}}.zip", binaryName: "terraform"},
		"windows/amd64": &asset{downloadURL: "https://releases.hashicorp.com/terraform/{{.Version}}/terraform_{{.Version}}_{{.Goos}}_{{.Goarch}}.zip", binaryName: "terraform.exe"},
	},
	"argoproj/argo-cd": {
		"default":       &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/argocd-{{.Goos}}-{{.Goarch}}", binaryName: "argocd"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/argocd-{{.Goos}}-{{.Goarch}}.exe", binaryName: "argocd.exe"},
	},
	"protocolbuffers/protobuf": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/protoc-{{.Version}}-linux-x86_64.zip", binaryName: "protoc"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/protoc-{{.Version}}-osx-x86_64.zip", binaryName: "protoc"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/protoc-{{.Version}}-win64.zip", binaryName: "protoc.exe"},
	},
	"mozilla/sops": {
		"default":       &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/sops-{{.Tag}}.{{.Goos}}", binaryName: "sops"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/sops-{{.Tag}}.exe", binaryName: "sops.exe"},
	},
	"CircleCI-Public/circleci-cli": {
		"default":       &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/circleci-cli_{{.Version}}_{{.Goos}}_{{.Goarch}}.tar.gz", binaryName: "circleci"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/circleci-cli_{{.Version}}_{{.Goos}}_{{.Goarch}}.zip", binaryName: "circleci.exe"},
	},
	"argoproj/argo-rollouts": {
		"default": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/kubectl-argo-rollouts-{{.Goos}}-{{.Goarch}}", binaryName: "kubectl-argo-rollouts"},
	},
	"open-policy-agent/conftest": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/conftest_{{.Version}}_Linux_x86_64.tar.gz", binaryName: "conftest"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/conftest_{{.Version}}_Darwin_x86_64.tar.gz", binaryName: "conftest"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/conftest_{{.Version}}_Windows_x86_64.zip", binaryName: "conftest.exe"},
	},
	"goodwithtech/dockle": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/dockle_{{.Version}}_Linux-64bit.tar.gz", binaryName: "dockle"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/dockle_{{.Version}}_macOS-64bit.tar.gz", binaryName: "dockle"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/dockle_{{.Version}}_Windows-64bit.zip", binaryName: "dockle.exe"},
	},
	"starship/starship": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/starship-x86_64-unknown-linux-gnu.tar.gz", binaryName: "starship"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/starship-x86_64-apple-darwin.tar.gz", binaryName: "starship"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/starship-x86_64-pc-windows-msvc.zip", binaryName: "starship.exe"},
	},
	"viaduct-ai/kustomize-sops": {
		"linux/amd64":   &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/ksops_{{.Version}}_Linux_x86_64.tar.gz", binaryName: "ksops"},
		"darwin/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/ksops_{{.Version}}_Darwin_x86_64.tar.gz", binaryName: "ksops"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/ksops_{{.Version}}_Windows_x86_64.tar.gz", binaryName: "ksops.exe"},
	},
	"fluxcd/flux2": {
		"default":       &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/flux_{{.Version}}_{{.Goos}}_{{.Goarch}}.tar.gz", binaryName: "flux"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/flux_{{.Version}}_{{.Goos}}_{{.Goarch}}.zip", binaryName: "flux.exe"},
	},
	"mikefarah/yq": {
		"default":       &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/yq_{{.Goos}}_{{.Goarch}}", binaryName: "yq"},
		"windows/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/yq_{{.Goos}}_{{.Goarch}}.exe", binaryName: "yq.exe"},
	},
	"aquasecurity/trivy": {
		"linux/amd64":  &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/trivy_{{.Version}}_Linux-64bit.tar.gz", binaryName: "trivy"},
		"darwin/amd64": &asset{downloadURL: "https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Tag}}/trivy_{{.Version}}_macOS-64bit.tar.gz", binaryName: "trivy"},
	},
}
