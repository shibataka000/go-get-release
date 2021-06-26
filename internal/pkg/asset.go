package pkg

var assetMap = map[string]map[string]string{
	"aquasecurity/trivy": {
		"linux/amd64":  "trivy_{{.Version}}_Linux-64bit.tar.gz",
		"darwin/amd64": "trivy_{{.Version}}_macOS-64bit.tar.gz",
	},
	"argoproj/argo-cd": {
		"default":       "argocd-{{.Goos}}-{{.Goarch}}",
		"windows/amd64": "argocd-{{.Goos}}-{{.Goarch}}.exe",
	},
	"aws/amazon-ec2-instance-selector": {
		"default": "ec2-instance-selector-{{.Goos}}-{{.Goarch}}",
	},
	"buildpacks/pack": {
		"linux/amd64":   "pack-{{.Tag}}-linux.tgz",
		"darwin/amd64":  "pack-{{.Tag}}-macos.tgz",
		"windows/amd64": "pack-{{.Tag}}-windows.zip",
	},
	"docker/compose": {
		"linux/amd64":   "docker-compose-Linux-x86_64",
		"darwin/amd64":  "docker-compose-Darwin-x86_64",
		"windows/amd64": "docker-compose-Windows-x86_64.exe",
	},
	"goodwithtech/dockle": {
		"linux/amd64":   "dockle_{{.Version}}_Linux-64bit.tar.gz",
		"darwin/amd64":  "dockle_{{.Version}}_macOS-64bit.tar.gz",
		"windows/amd64": "dockle_{{.Version}}_Windows-64bit.zip",
	},
	"istio/istio": {
		"linux/amd64":   "istioctl-{{.Tag}}-{{.Goos}}-{{.Goarch}}.tar.gz",
		"darwin/amd64":  "istioctl-{{.Tag}}-osx.tar.gz",
		"windows/amd64": "istioctl-{{.Tag}}-win.zip",
	},
	"mikefarah/yq": {
		"default":       "yq_{{.Goos}}_{{.Goarch}}",
		"windows/amd64": "yq_{{.Goos}}_{{.Goarch}}.exe",
	},
	"mozilla/sops": {
		"default":       "sops-{{.Tag}}.{{.Goos}}",
		"windows/amd64": "sops-{{.Tag}}.exe",
	},
	"open-policy-agent/opa": {
		"default":       "opa_{{.Goos}}_{{.Goarch}}",
		"windows/amd64": "opa_{{.Goos}}_{{.Goarch}}.exe",
	},
	"protocolbuffers/protobuf": {
		"linux/amd64":   "protoc-{{.Version}}-linux-x86_64.zip",
		"darwin/amd64":  "protoc-{{.Version}}-osx-x86_64.zip",
		"windows/amd64": "protoc-{{.Version}}-win64.zip",
	},
	"starship/starship": {
		"linux/amd64":   "starship-x86_64-unknown-linux-gnu.tar.gz",
		"darwin/amd64":  "starship-x86_64-apple-darwin.tar.gz",
		"windows/amd64": "starship-x86_64-pc-windows-msvc.zip",
	},
	"viaduct-ai/kustomize-sops": {
		"linux/amd64":   "ksops_{{.Version}}_Linux_x86_64.tar.gz",
		"darwin/amd64":  "ksops_{{.Version}}_Darwin_x86_64.tar.gz",
		"windows/amd64": "ksops_{{.Version}}_Windows_x86_64.tar.gz",
	},
}

var binaryMap = map[string]string{
	"argoproj/argo-workflows":          "argo",
	"argoproj/argo-cd":                 "argocd",
	"argoproj/argo-rollouts":           "kubectl-argo-rollouts",
	"aws/amazon-ec2-instance-selector": "ec2-instance-selector",
	"CircleCI-Public/circleci-cli":     "circleci",
	"docker/compose":                   "docker-compose",
	"docker/machine":                   "docker-machine",
	"fluxcd/flux2":                     "flux",
	"istio/istio":                      "istioctl",
	"protocolbuffers/protobuf":         "protoc",
	"viaduct-ai/kustomize-sops":        "ksops",
}
