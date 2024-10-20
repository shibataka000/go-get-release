package github

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/cheggaaa/pb/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/go-github/v62/github"
	"github.com/ulikunitz/xz"
)

// Asset represents a GitHub release asset.
type Asset struct {
	id   int64
	Name string
}

// newAsset returns a new [Asset] object.
func newAsset(id int64, name string) Asset {
	return Asset{
		id:   id,
		Name: name,
	}
}

// AssetList is a list of [Asset].
type AssetList []Asset

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// extract [ExecBinaryContent] from [AssetContent] and return it.
func (a AssetContent) extract() (ExecBinaryContent, error) {
	b := []byte(a)

	for !isOctetStream(b) {
		r, c, err := newReaderToExtract(b)
		if err != nil {
			return nil, err
		}
		b, err = io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		if c != nil {
			if err := c.Close(); err != nil {
				return nil, err
			}
		}
	}

	return ExecBinaryContent(b), nil
}

func isOctetStream(b []byte) bool {
	expect := []string{"application/octet-stream", "application/x-executable"}
	mime := mimetype.Detect(b)
	return slices.Contains(expect, mime.String())
}

func newReaderToExtract(b []byte) (io.Reader, io.Closer, error) {
	br := bytes.NewReader(b)
	mime := mimetype.Detect(b)

	switch mime.String() {
	case "application/gzip":
		r, err := gzip.NewReader(br)
		return r, nil, err
	case "application/x-xz":
		r, err := xz.NewReader(br)
		return r, nil, err
	case "application/x-tar":
		r, err := newExecBinaryReaderInTar(br)
		return r, nil, err
	case "application/zip":
		r, err := newExecBinaryReaderInZip(br, br.Size())
		return r, r, err
	default:
		return nil, nil, fmt.Errorf("%w: %s", ErrUnexpectedMIME, mime.String())
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
