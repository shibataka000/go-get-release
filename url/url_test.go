package url

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLBase(t *testing.T) {
	tests := []struct {
		name string
		url  URL
		base string
	}{
		{
			name: "https://example.com/a/b/c",
			url:  "https://example.com/a/b/c",
			base: "c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.base, tt.url.Base())
		})
	}
}

func TestTemplateExecute(t *testing.T) {
	type Data struct {
		A string
		B string
	}

	tests := []struct {
		name string
		tmpl Template
		data Data
		url  URL
	}{
		{
			name: "https://example.com/aaa/bbb",
			tmpl: "https://example.com/{{.A}}/{{.B}}",
			data: Data{
				A: "aaa",
				B: "bbb",
			},
			url: "https://example.com/aaa/bbb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			url, err := tt.tmpl.Execute(tt.data)
			require.NoError(err)
			require.Equal(tt.url, url)
		})
	}
}
