package github

import "text/template"

// externalAssets is a map whose key is repository and whose value is a template list of GitHub release asset on server outside GitHub.
var externalAssets = map[Repository]AssetTemplateList{
	newRepository("hashicorp", "terraform"): {
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip"), ""),
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip"), ""),
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip"), ""),
	},
}

// newTemplate allocates a new template with given name and parses text as a template body for it.
func newTemplate(name string, text string) *template.Template {
	return template.Must(template.New(name).Parse(text))
}
