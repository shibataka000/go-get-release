package pkg

import (
	"bytes"
	"path"
	"text/template"
)

// URL.
type URL string

// URLTemplate is URL template.
type URLTemplate string

// NewURL return URL instance.
func NewURL(url string) URL {
	return URL(url)
}

// String return string typed URL.
func (url URL) String() string {
	return string(url)
}

// FileName return file name downloaded from this URL.
func (url URL) FileName() FileName {
	base := path.Base(url.String())
	return NewFileName(base)
}

// NewURLTemplate return URL template instance.
func NewURLTemplate(tmpl string) URLTemplate {
	return URLTemplate(tmpl)
}

// String return string typed URL template.
func (url URLTemplate) String() string {
	return string(url)
}

// RenderWithRelease render URL with GitHub release.
func (url URLTemplate) RenderWithRelease(release GitHubRelease) (URL, error) {
	semver, err := release.SemVer()
	if err != nil {
		return "", err
	}
	param := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.Tag,
		SemVer: semver,
	}

	tmpl, err := template.New("").Parse(url.String())
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, param)
	if err != nil {
		return "", err
	}
	return NewURL(buf.String()), nil
}
