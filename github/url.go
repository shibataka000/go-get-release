package github

import (
	"bytes"
	"path"
	"text/template"
)

// URL.
type URL string

// URLTemplate is URL template.
type URLTemplate string

// String return string typed URL.
func (url URL) String() string {
	return string(url)
}

// FileName return file name of downloaded file from this URL.
func (url URL) FileName() FileName {
	base := path.Base(url.String())
	return FileName(base)
}

// String return string typed URL template.
func (url URLTemplate) String() string {
	return string(url)
}

// RenderWithRelease render URL with release.
func (url URLTemplate) RenderWithRelease(release Release) (URL, error) {
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
	return URL(buf.String()), nil
}
