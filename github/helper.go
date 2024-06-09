package github

import "text/template"

// newTemplate allocates a new template with given name and parses text as a template body for it.
func newTemplate(name string, text string) *template.Template {
	return template.Must(template.New(name).Parse(text))
}
