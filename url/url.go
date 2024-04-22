package url

import (
	"bytes"
	"path"
	"text/template"
)

// URL.
type URL string

// Template is URL template.
type Template string

// Base returns the last element of URL.
func (url URL) Base() string {
	return path.Base(string(url))
}

// Execute applies a URL template to the specified data object, and return it as URL.
func (url Template) Execute(data any) (URL, error) {
	tmpl, err := template.New("url").Parse(string(url))
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return URL(buf.String()), nil
}
