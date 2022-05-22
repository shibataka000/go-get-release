package github

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

var assetNameMap = map[string]map[string]string{
	"aquasecurity/tfsec": {
		"default":       "tfsec-{{.Goos}}-{{.Goarch}}",
		"windows/amd64": "tfsec-{{.Goos}}-{{.Goarch}}.exe",
	},
	"argoproj/argo-cd": {
		"default":       "argocd-{{.Goos}}-{{.Goarch}}",
		"windows/amd64": "argocd-{{.Goos}}-{{.Goarch}}.exe",
	},
	"istio/istio": {
		"default":       "istioctl-{{.Tag}}-{{.Goos}}-{{.Goarch}}.tar.gz",
		"darwin/amd64":  "istioctl-{{.Tag}}-osx.tar.gz",
		"windows/amd64": "istioctl-{{.Tag}}-win.zip",
	},
	"open-policy-agent/opa": {
		"linux/amd64": "opa_{{.Goos}}_{{.Goarch}}",
	},
	"starship/starship": {
		"linux/amd64": "starship-x86_64-unknown-linux-gnu.tar.gz",
	},
	"gravitational/teleport": {
		"default":       "teleport-{{.Tag}}-{{.Goos}}-{{.Goarch}}-bin.tar.gz",
		"windows/amd64": "teleport-{{.Tag}}-{{.Goos}}-{{.Goarch}}-bin.zip",
	},
	"viaduct-ai/kustomize-sops": {
		"linux/amd64":   "ksops_{{.Version}}_Linux_x86_64.tar.gz",
		"darwin/amd64":  "ksops_{{.Version}}_Darwin_x86_64.tar.gz",
		"windows/amd64": "ksops_{{.Version}}_Windows_x86_64.tar.gz",
	},
}

var binaryNameMap = map[string]string{
	"argoproj/argo-workflows":          "argo",
	"argoproj/argo-cd":                 "argocd",
	"argoproj/argo-rollouts":           "kubectl-argo-rollouts",
	"aws/amazon-ec2-instance-selector": "ec2-instance-selector",
	"CircleCI-Public/circleci-cli":     "circleci",
	"cli/cli":                          "gh",
	"docker/compose":                   "docker-compose",
	"docker/machine":                   "docker-machine",
	"docker/scan-cli-plugin":           "docker-scan",
	"fluxcd/flux2":                     "flux",
	"istio/istio":                      "istioctl",
	"protocolbuffers/protobuf":         "protoc",
	"gravitational/teleport":           "tsh",
	"viaduct-ai/kustomize-sops":        "ksops",
}
