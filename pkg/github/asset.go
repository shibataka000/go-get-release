package github

// Asset in github release.
type Asset struct {
	Owner       string
	Repo        string
	Tag         string
	DownloadURL string
	BinaryName  string
	Goos        string
	Goarch      string
}

// RegisteredAsset is registered asset data.
type RegisteredAsset struct {
	Owner               string
	Repo                string
	DownloadURLTemplate string
	BinaryName          string
	Goos                string
	Goarch              string
}

// List of GOOS, which are same as result of following command.
// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\1/g" | sort | uniq`
var goosMap = map[string][]string{
	"aix":       {"aix"},
	"android":   {"android"},
	"darwin":    {"darwin", "macos", "osx"},
	"dragonfly": {"dragonfly"},
	"freebsd":   {"freebsd"},
	"illumos":   {"illumos"},
	"ios":       {"ios"},
	"js":        {"js"},
	"linux":     {"linux"},
	"netbsd":    {"netbsd"},
	"openbsd":   {"openbsd"},
	"plan9":     {"plan9"},
	"solaris":   {"solaris"},
	"windows":   {"windows", "win", ".exe"},
}

// List of GOARCH, which are same as result of following command.
// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\2/g" | sort | uniq`
var goarchMap = map[string][]string{
	"386":      {"386", "x86_32", "32bit", "win32"},
	"amd64":    {"amd64", "x86_64", "64bit", "win64"},
	"arm":      {"arm"},
	"arm64":    {"arm64", "aarch64", "aarch_64"},
	"mips":     {"mips"},
	"mips64":   {"mips64"},
	"mips64le": {"mips64le"},
	"mipsle":   {"mipsle"},
	"ppc64":    {"ppc64"},
	"ppc64le":  {"ppc64le", "ppcle_64"},
	"riscv64":  {"riscv64"},
	"s390x":    {"s390x", "s390"},
	"wasm":     {"wasm"},
}

var registeredAsset = []RegisteredAsset{
	// aquasecurity/tfsec
	{
		Owner:               "aquasecurity",
		Repo:                "tfsec",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/aquasecurity/tfsec/releases/download/{{.Tag}}/tfsec-linux-amd64",
		BinaryName:          "tfsec",
	},
	{
		Owner:               "aquasecurity",
		Repo:                "tfsec",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/aquasecurity/tfsec/releases/download/{{.Tag}}/tfsec-darwin-amd64",
		BinaryName:          "tfsec",
	},
	{
		Owner:               "aquasecurity",
		Repo:                "tfsec",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/aquasecurity/tfsec/releases/download/{{.Tag}}/tfsec-windows-amd64.exe",
		BinaryName:          "tfsec.exe",
	},
	// argoproj/argo-workflows
	{
		Owner:               "argoproj",
		Repo:                "argo-workflows",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-workflows/releases/download/{{.Tag}}/argo-linux-amd64.gz",
		BinaryName:          "argo",
	},
	{
		Owner:               "argoproj",
		Repo:                "argo-workflows",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-workflows/releases/download/{{.Tag}}/argo-darwin-amd64.gz",
		BinaryName:          "argo",
	},
	{
		Owner:               "argoproj",
		Repo:                "argo-workflows",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-workflows/releases/download/{{.Tag}}/argo-windows-amd64.gz",
		BinaryName:          "argo.exe",
	},
	// argoproj/argo-cd
	{
		Owner:               "argoproj",
		Repo:                "argo-cd",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-cd/releases/download/{{.Tag}}/argocd-linux-amd64",
		BinaryName:          "argocd",
	},
	{
		Owner:               "argoproj",
		Repo:                "argo-cd",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-cd/releases/download/{{.Tag}}/argocd-darwin-amd64",
		BinaryName:          "argocd",
	},
	{
		Owner:               "argoproj",
		Repo:                "argo-cd",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-cd/releases/download/{{.Tag}}/argocd-windows-amd64.exe",
		BinaryName:          "argocd.exe",
	},
	// argoproj/argo-rollouts
	{
		Owner:               "argoproj",
		Repo:                "argo-rollouts",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-rollouts/releases/download/{{.Tag}}/kubectl-argo-rollouts-linux-amd64",
		BinaryName:          "kubectl-argo-rollouts",
	},
	{
		Owner:               "argoproj",
		Repo:                "argo-rollouts",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/argoproj/argo-rollouts/releases/download/{{.Tag}}/kubectl-argo-rollouts-darwin-amd64",
		BinaryName:          "kubectl-argo-rollouts",
	},
	// aws/amazon-ec2-instance-selector
	{
		Owner:               "aws",
		Repo:                "amazon-ec2-instance-selector",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/{{.Tag}}/ec2-instance-selector-linux-amd64",
		BinaryName:          "ec2-instance-selector",
	},
	{
		Owner:               "aws",
		Repo:                "amazon-ec2-instance-selector",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/{{.Tag}}/ec2-instance-selector-darwin-amd64",
		BinaryName:          "ec2-instance-selector",
	},
	{
		Owner:               "aws",
		Repo:                "amazon-ec2-instance-selector",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/aws/amazon-ec2-instance-selector/releases/download/{{.Tag}}/ec2-instance-selector-windows-amd64",
		BinaryName:          "ec2-instance-selector.exe",
	},
	// CircleCI-Public/circleci-cli
	{
		Owner:               "CircleCI-Public",
		Repo:                "circleci-cli",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/CircleCI-Public/circleci-cli/releases/download/{{.Tag}}/circleci-cli_{{.Version}}_linux_amd64.tar.gz",
		BinaryName:          "circleci",
	},
	{
		Owner:               "CircleCI-Public",
		Repo:                "circleci-cli",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/CircleCI-Public/circleci-cli/releases/download/{{.Tag}}/circleci-cli_{{.Version}}_darwin_amd64.tar.gz",
		BinaryName:          "circleci",
	},
	{
		Owner:               "CircleCI-Public",
		Repo:                "circleci-cli",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/CircleCI-Public/circleci-cli/releases/download/{{.Tag}}/circleci-cli_{{.Version}}_windows_amd64.zip",
		BinaryName:          "circleci.exe",
	},
	// cli/cli
	{
		Owner:               "cli",
		Repo:                "cli",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/cli/cli/releases/download/{{.Tag}}/gh_{{.Version}}_linux_amd64.tar.gz",
		BinaryName:          "gh",
	},
	{
		Owner:               "cli",
		Repo:                "cli",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/cli/cli/releases/download/{{.Tag}}/gh_{{.Version}}_macOS_amd64.tar.gz",
		BinaryName:          "gh",
	},
	{
		Owner:               "cli",
		Repo:                "cli",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/cli/cli/releases/download/{{.Tag}}/gh_{{.Version}}_windows_amd64.zip",
		BinaryName:          "gh.exe",
	},
	// docker/compose
	{
		Owner:               "docker",
		Repo:                "compose",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/compose/releases/download/{{.Tag}}/docker-compose-Linux-x86_64",
		BinaryName:          "docker-compose",
	},
	{
		Owner:               "docker",
		Repo:                "compose",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/compose/releases/download/{{.Tag}}/docker-compose-Darwin-x86_64",
		BinaryName:          "docker-compose",
	},
	{
		Owner:               "docker",
		Repo:                "compose",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/compose/releases/download/{{.Tag}}/docker-compose-Windows-x86_64.exe",
		BinaryName:          "docker-compose.exe",
	},
	// docker/machine
	{
		Owner:               "docker",
		Repo:                "machine",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/machine/releases/download/{{.Tag}}/docker-machine-Linux-x86_64",
		BinaryName:          "docker-machine",
	},
	{
		Owner:               "docker",
		Repo:                "machine",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/machine/releases/download/{{.Tag}}/docker-machine-Darwin-x86_64",
		BinaryName:          "docker-machine",
	},
	{
		Owner:               "docker",
		Repo:                "machine",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/machine/releases/download/{{.Tag}}/docker-machine-Windows-x86_64.exe",
		BinaryName:          "docker-machine.exe",
	},
	// docker/scan-cli-plugin
	{
		Owner:               "docker",
		Repo:                "scan-cli-plugin",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/scan-cli-plugin/releases/download/{{.Tag}}/docker-scan_linux_amd64",
		BinaryName:          "docker-scan",
	},
	{
		Owner:               "docker",
		Repo:                "scan-cli-plugin",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/scan-cli-plugin/releases/download/{{.Tag}}/docker-scan_darwin_amd64",
		BinaryName:          "docker-scan",
	},
	{
		Owner:               "docker",
		Repo:                "scan-cli-plugin",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/docker/scan-cli-plugin/releases/download/{{.Tag}}/docker-scan_windows_amd64.exe",
		BinaryName:          "docker-scan.exe",
	},
	// fluxcd/flux2
	{
		Owner:               "fluxcd",
		Repo:                "flux2",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/fluxcd/flux2/releases/download/{{.Tag}}/flux_{{.Version}}_linux_amd64.tar.gz",
		BinaryName:          "flux",
	},
	{
		Owner:               "fluxcd",
		Repo:                "flux2",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/fluxcd/flux2/releases/download/{{.Tag}}/flux_{{.Version}}_darwin_amd64.tar.gz",
		BinaryName:          "flux",
	},
	{
		Owner:               "fluxcd",
		Repo:                "flux2",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/fluxcd/flux2/releases/download/{{.Tag}}/flux_{{.Version}}_windows_amd64.zip",
		BinaryName:          "flux.exe",
	},
	// gravitational/teleport
	{
		Owner:               "gravitational",
		Repo:                "teleport",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://get.gravitational.com/teleport-v{{.Version}}-linux-amd64-bin.tar.gz",
		BinaryName:          "tsh",
	},
	{
		Owner:               "gravitational",
		Repo:                "teleport",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://get.gravitational.com/teleport-v{{.Version}}-darwin-amd64-bin.tar.gz",
		BinaryName:          "tsh",
	},
	{
		Owner:               "gravitational",
		Repo:                "teleport",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://get.gravitational.com/teleport-v{{.Version}}-windows-amd64-bin.zip",
		BinaryName:          "tsh.exe",
	},
	// hashicorp/terraform
	{
		Owner:               "hashicorp",
		Repo:                "terraform",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://releases.hashicorp.com/terraform/{{.Version}}/terraform_{{.Version}}_linux_amd64.zip",
		BinaryName:          "terraform",
	},
	{
		Owner:               "hashicorp",
		Repo:                "terraform",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://releases.hashicorp.com/terraform/{{.Version}}/terraform_{{.Version}}_darwin_amd64.zip",
		BinaryName:          "terraform",
	},
	{
		Owner:               "hashicorp",
		Repo:                "terraform",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://releases.hashicorp.com/terraform/{{.Version}}/terraform_{{.Version}}_windows_amd64.zip",
		BinaryName:          "terraform.exe",
	},
	// helm/helm
	{
		Owner:               "helm",
		Repo:                "helm",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://get.helm.sh/helm-v{{.Version}}-linux-amd64.tar.gz",
		BinaryName:          "helm",
	},
	{
		Owner:               "helm",
		Repo:                "helm",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://get.helm.sh/helm-v{{.Version}}-darwin-amd64.tar.gz",
		BinaryName:          "helm",
	},
	{
		Owner:               "helm",
		Repo:                "helm",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://get.helm.sh/helm-v{{.Version}}-windows-amd64.zip",
		BinaryName:          "helm.exe",
	},
	// istio/istio
	{
		Owner:               "istio",
		Repo:                "istio",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/istio/istio/releases/download/{{.Tag}}/istioctl-{{.Version}}-linux-amd64.tar.gz",
		BinaryName:          "istioctl",
	},
	{
		Owner:               "istio",
		Repo:                "istio",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/istio/istio/releases/download/{{.Tag}}/istioctl-{{.Version}}-osx.tar.gz",
		BinaryName:          "istioctl",
	},
	{
		Owner:               "istio",
		Repo:                "istio",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/istio/istio/releases/download/{{.Tag}}/istioctl-{{.Version}}-win.zip",
		BinaryName:          "istioctl.exe",
	},
	// mikefarah/yq
	{
		Owner:               "mikefarah",
		Repo:                "yq",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/mikefarah/yq/releases/download/{{.Tag}}/yq_linux_amd64",
		BinaryName:          "yq",
	},
	{
		Owner:               "mikefarah",
		Repo:                "yq",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/mikefarah/yq/releases/download/{{.Tag}}/yq_darwin_amd64",
		BinaryName:          "yq",
	},
	{
		Owner:               "mikefarah",
		Repo:                "yq",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/mikefarah/yq/releases/download/{{.Tag}}/yq_windows_amd64.exe",
		BinaryName:          "yq.exe",
	},
	// mozilla/sops
	{
		Owner:               "mozilla",
		Repo:                "sops",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/mozilla/sops/releases/download/{{.Tag}}/sops-v{{.Version}}.linux",
		BinaryName:          "sops",
	},
	{
		Owner:               "mozilla",
		Repo:                "sops",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/mozilla/sops/releases/download/{{.Tag}}/sops-v{{.Version}}.darwin",
		BinaryName:          "sops",
	},
	{
		Owner:               "mozilla",
		Repo:                "sops",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/mozilla/sops/releases/download/{{.Tag}}/sops-v{{.Version}}.exe",
		BinaryName:          "sops.exe",
	},
	// open-policy-agent/opa
	{
		Owner:               "open-policy-agent",
		Repo:                "opa",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/open-policy-agent/opa/releases/download/{{.Tag}}/opa_linux_amd64",
		BinaryName:          "opa",
	},
	// protocolbuffers/protobuf
	{
		Owner:               "protocolbuffers",
		Repo:                "protobuf",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/protocolbuffers/protobuf/releases/download/{{.Tag}}/protoc-{{.Version}}-linux-x86_64.zip",
		BinaryName:          "protoc",
	},
	{
		Owner:               "protocolbuffers",
		Repo:                "protobuf",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/protocolbuffers/protobuf/releases/download/{{.Tag}}/protoc-{{.Version}}-osx-x86_64.zip",
		BinaryName:          "protoc",
	},
	{
		Owner:               "protocolbuffers",
		Repo:                "protobuf",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/protocolbuffers/protobuf/releases/download/{{.Tag}}/protoc-{{.Version}}-win64.zip",
		BinaryName:          "protoc.exe",
	},
	// starship/starship
	{
		Owner:               "starship",
		Repo:                "starship",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/starship/starship/releases/download/{{.Tag}}/starship-x86_64-unknown-linux-gnu.tar.gz",
		BinaryName:          "starship",
	},
	{
		Owner:               "starship",
		Repo:                "starship",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/starship/starship/releases/download/{{.Tag}}/starship-x86_64-apple-darwin.tar.gz",
		BinaryName:          "starship",
	},
	{
		Owner:               "starship",
		Repo:                "starship",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/starship/starship/releases/download/{{.Tag}}/starship-x86_64-pc-windows-msvc.zip",
		BinaryName:          "starship.exe",
	},
	// viaduct-ai/kustomize-sops
	{
		Owner:               "viaduct-ai",
		Repo:                "kustomize-sops",
		Goos:                "linux",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/viaduct-ai/kustomize-sops/releases/download/{{.Tag}}/ksops_{{.Version}}_Linux_x86_64.tar.gz",
		BinaryName:          "ksops",
	},
	{
		Owner:               "viaduct-ai",
		Repo:                "kustomize-sops",
		Goos:                "darwin",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/viaduct-ai/kustomize-sops/releases/download/{{.Tag}}/ksops_{{.Version}}_Darwin_x86_64.tar.gz",
		BinaryName:          "ksops",
	},
	{
		Owner:               "viaduct-ai",
		Repo:                "kustomize-sops",
		Goos:                "windows",
		Goarch:              "amd64",
		DownloadURLTemplate: "https://github.com/viaduct-ai/kustomize-sops/releases/download/{{.Tag}}/ksops_{{.Version}}_Windows_x86_64.tar.gz",
		BinaryName:          "ksops.exe",
	},
}
