package github

import "regexp"

// externalAssets is a map whose key is repository and whose value is a list of GitHub release asset template on external server.
var externalAssets = map[Repository]AssetTemplateList{
	newRepository("hashicorp", "terraform"): {
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_arm64.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_386.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_amd64.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_freebsd_arm.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_386.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_arm.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_arm64.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_openbsd_386.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_openbsd_amd64.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_solaris_amd64.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_386.zip")),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip")),
	},
	newRepository("helm", "helm"): {
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-darwin-amd64.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-darwin-arm64.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-linux-386.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-linux-amd64.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-linux-arm.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-linux-arm64.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-linux-ppc64le.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-linux-riscv64.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-linux-s390x.tar.gz")),
		newAssetTemplate(newTemplate("https://get.helm.sh/helm-v3.15.2-windows-amd64.zip")),
	},
}

var ignoredAssets = AssetRegexpList{
	newAssetRegexp(regexp.MustCompile("abcd")),
}
