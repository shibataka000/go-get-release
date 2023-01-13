package http

import (
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/stretchr/testify/require"
)

func TestDownloadAssetContent(t *testing.T) {
	tests := []struct {
		meta            asset.Metadata
		content         asset.Content
		showProgressBar bool
	}{
		{
			meta:            asset.NewMetadata("https://raw.githubusercontent.com/shibataka000/go-get-release-test/master/README.md"),
			content:         asset.NewContent("README.md", []byte("# go-get-release-test")),
			showProgressBar: false,
		},
	}

	repository := New()

	for _, tt := range tests {
		name := tt.meta.DownloadURL()
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			content, err := repository.DownloadAssetContent(tt.meta, tt.showProgressBar)
			assert.NoError(err)
			assert.Equal(tt.content, content)
		})
	}
}
