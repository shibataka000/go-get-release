package github

import "text/template"

func newTemplate(name string, text string) *template.Template {
	return template.Must(template.New(name).Parse(text))
}
