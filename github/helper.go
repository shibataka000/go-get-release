package github

import "text/template"

// newTemplate allocates a new template and parses text as a template body for it.
// This gets into a panic if the error is non-nil.
func newTemplate(text string) *template.Template {
	return template.Must(template.New("").Parse(text))
}
