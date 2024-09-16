package github

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cheggaaa/pb/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/go-github/v62/github"
	"github.com/ulikunitz/xz"
)

// Asset represents a GitHub release asset.
type Asset struct {
	id   int64
	name string
}

// newAsset returns a new [Asset] object.
func newAsset(id int64, name string) Asset {
	return Asset{
		id:   id,
		name: name,
	}
}

// AssetList is a list of [Asset].
type AssetList []Asset

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// execBinaryContent extracts [ExecBinaryContent] from [AssetContent] and return it.
func (a AssetContent) execBinaryContent() (ExecBinaryContent, error) {
	var b bytes.Buffer

	if _, err := b.Write(a); err != nil {
		return nil, err
	}

	for {
		var r io.Reader
		var err error

		mime := mimetype.Detect(b.Bytes())

		switch mime.String() {
		case "application/octet-stream", "application/x-executable":
			return ExecBinaryContent(b.Bytes()), nil
		case "application/x-tar":
			r, err = newExecBinaryReaderInTar(&b)
		case "application/zip":
			r, err = newExecBinaryReaderInZip(&b)
		case "application/gzip":
			r, err = gzip.NewReader(&b)
		case "application/x-xz":
			r, err = xz.NewReader(&b)
		default:
			r, err = nil, fmt.Errorf("%w: %s", ErrUnexpectedMIME, mime.String())
		}

		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrExtractingExecBinaryContentFailure, err)
		}

		var bb bytes.Buffer
		if _, err := bb.ReadFrom(r); err != nil {
			return nil, err
		}
		b = bb
	}
}

// AssetRepository is a repository for [Asset] and [AssetContent].
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository returns a new [AssetRepository] object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	client := github.NewClient(http.DefaultClient)
	if token != "" {
		client = client.WithAuthToken(token)
	}
	return &AssetRepository{
		client: client,
	}
}

// list GitHub release assets and returns it.
func (r *AssetRepository) list(ctx context.Context, repo Repository, release Release) (AssetList, error) {
	assets := AssetList{}

	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.owner, repo.name, release.tag)
	if err != nil {
		return nil, err
	}

	for page := 1; page != 0; {
		githubAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.owner, repo.name, githubRelease.GetID(), &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, githubAsset := range githubAssets {
			assets = append(assets, newAsset(githubAsset.GetID(), githubAsset.GetName()))
		}
		page = resp.NextPage
	}

	return assets, nil
}

// download a GitHub release asset content and returns it.
// This writes progress bar to w.
func (r *AssetRepository) download(ctx context.Context, repo Repository, asset Asset, w io.Writer) (AssetContent, error) {
	githubAsset, _, err := r.client.Repositories.GetReleaseAsset(ctx, repo.owner, repo.name, asset.id)
	if err != nil {
		return nil, err
	}

	rc, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, repo.owner, repo.name, asset.id, http.DefaultClient)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	total := int64(githubAsset.GetSize())
	pr := pb.Full.Start64(total).SetWriter(w).NewProxyReader(rc)
	defer pr.Close()

	return io.ReadAll(pr)
}
