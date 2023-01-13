package http

import (
	"bytes"
	"io"
	"net/http"

	"github.com/cheggaaa/pb/v3"
	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
)

// DownloadAssetContent download and return asset content.
func (r *Repository) DownloadAssetContent(meta asset.Metadata, showProgressBar bool) (asset.Content, error) {
	resp, err := http.Get(meta.DownloadURL())
	if err != nil {
		return asset.Content{}, err
	}
	defer resp.Body.Close()

	in := resp.Body
	if showProgressBar {
		bar := pb.Full.Start64(resp.ContentLength)
		in = bar.NewProxyReader(resp.Body)
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, in)
	if err != nil {
		return asset.Content{}, err
	}

	return asset.NewContent(meta.Name(), buf.Bytes()), nil
}
