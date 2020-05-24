package github

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

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
}

func isSpecialAsset(owner, repo string) bool {
	key := fmt.Sprintf("%s/%s", owner, repo)
	_, ok := specialAssetMap[key]
	return ok
}

func getSpecialAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	key := fmt.Sprintf("%s/%s", repo.owner, repo.name)
	assetMap, ok := specialAssetMap[key]
	if !ok {
		return nil, fmt.Errorf("%s is not found in specialAssetMap", key)
	}

	var assetTemplate *asset
	key = fmt.Sprintf("%s/%s", goos, goarch)
	if value, ok := assetMap[key]; ok {
		assetTemplate = value
	}
	if value, ok := assetMap["default"]; assetTemplate == nil && ok {
		assetTemplate = value
	}
	if assetTemplate == nil {
		return nil, fmt.Errorf("Unsupported GOOS and GOARCH in this repository: %s", key)
	}

	version := strings.TrimLeft(release.tag, "v")

	tmpl, err := template.New("downloadURL").Parse(assetTemplate.downloadURL)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, struct {
		Owner   string
		Repo    string
		Tag     string
		Version string
		Goos    string
		Goarch  string
	}{
		Owner:   repo.owner,
		Repo:    repo.name,
		Tag:     release.tag,
		Version: version,
		Goos:    goos,
		Goarch:  goarch,
	})
	if err != nil {
		return nil, err
	}
	downloadURL := buf.String()

	binaryName := assetTemplate.binaryName

	return &asset{
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil
}
