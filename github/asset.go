package github

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/cheggaaa/pb/v3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/go-github/v62/github"
	"github.com/ulikunitz/xz"
	"golang.org/x/oauth2"
)

// Asset represents a GitHub release asset.
type Asset struct {
	downloadURL *url.URL
}

// newAsset returns a new GitHub release asset object.
func newAsset(downloadURL *url.URL) Asset {
	return Asset{
		downloadURL: downloadURL,
	}
}

// newAssetFromString returns a new GitHub release asset object.
func newAssetFromString(downloadURL string) (Asset, error) {
	url, err := url.Parse(downloadURL)
	if err != nil {
		return Asset{}, err
	}
	return newAsset(url), nil
}

// name returns a name of GitHub release asset.
func (a Asset) name() string {
	return path.Base(a.downloadURL.String())
}

// AssetList represents a list of GitHub release assets.
type AssetList []Asset

// find a GitHub release asset which matches any of given patterns.
func (al AssetList) find(patterns PatternList) (Asset, error) {
	for _, p := range patterns {
		for _, a := range al {
			if p.match(a) {
				return a, nil
			}
		}
	}
	return Asset{}, ErrAssetNotFound
}

type AssetContent []byte

func (a AssetContent) execBinary() (ExecBinaryContent, error) {
	var b bytes.Buffer

	if _, err := b.Write(a); err != nil {
		return nil, err
	}

	for {
		var r io.Reader
		var err error

		mime := mimetype.Detect(b.Bytes())

		switch mime.String() {
		case "application/x-tar":
			r, err = newTarReader(&b)
		case "application/zip":
			r, err = newZipReader(bytes.NewReader(b.Bytes()), int64(b.Len()))
		case "application/gzip":
			r, err = gzip.NewReader(&b)
		case "application/x-xz":
			r, err = xz.NewReader(&b)
		case "application/octet-stream":
			return ExecBinaryContent(b.Bytes()), nil
		default:
			return nil, ErrUnsupportedMIME
		}
		if err != nil {
			return nil, err
		}

		b.Reset()
		if _, err := b.ReadFrom(r); err != nil {
			return nil, err
		}
	}
}

// AssetRepository is a repository for a GitHub release asset.
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository returns a new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	return &AssetRepository{
		client: github.NewClient(httpClient),
	}
}

// list returns a list of GitHub release assets.
func (r *AssetRepository) list(ctx context.Context, repo Repository, release Release) (AssetList, error) {
	assets := AssetList{}

	// Get GitHub release ID.
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.owner, repo.name, release.tag)
	if err != nil {
		return nil, err
	}
	releaseID := *githubRelease.ID

	// List GitHub release assets.
	for page := 1; page != 0; {
		githubAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.owner, repo.name, releaseID, &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, githubAsset := range githubAssets {
			downloadURL := githubAsset.GetBrowserDownloadURL()
			asset, err := newAssetFromString(downloadURL)
			if err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		}
		page = resp.NextPage
	}

	return assets, nil
}

func (r *AssetRepository) download(asset Asset, progressBar io.Writer) (AssetContent, error) {
	resp, err := http.Get(asset.downloadURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	src := pb.Full.Start64(resp.ContentLength).SetWriter(progressBar).NewProxyReader(resp.Body)
	dst := new(bytes.Buffer)

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	return AssetContent(dst.Bytes()), nil
}
