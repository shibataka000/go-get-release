package github

import (
	"text/template"

	"github.com/shibataka000/go-get-release/mime"
)

// externalAssets is a map whose key is repository and whose value is a template list of GitHub release asset on server outside GitHub.
var externalAssets = map[Repository]AssetTemplateList{
	newRepository("hashicorp", "terraform"): {
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip"), mime.Zip),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip"), mime.Zip),
		newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip"), mime.Zip),
	},
}

// newTemplate allocates a new template and parses text as a template body for it.
// This gets into a panic if the error is non-nil.
func newTemplate(text string) *template.Template {
	return template.Must(template.New("").Parse(text))
}
