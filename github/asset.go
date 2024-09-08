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
	"golang.org/x/oauth2"
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

func (a Asset) execBinary(assetPatterns AssetPatternList, execBinaryPatterns ExecBinaryPatternList) (ExecBinary, error) {
	return ExecBinary{}, nil
}

// AssetList is a list of [Asset].
type AssetList []Asset

// find a [Asset] whose name matches any of [AssetPattern].
// If two or more [Asset] match, this returns a [Asset] which matches prior [AssetPattern].
func (al AssetList) find(patterns []AssetPattern) (Asset, error) {
	for _, p := range patterns {
		for _, a := range al {
			if p.match(a) {
				return a, nil
			}
		}
	}
	return Asset{}, ErrAssetNotFound
}

// AssetPattern is regular expression which matches a [Asset]'s name.
type AssetPattern Pattern

// match returns true if [AssetPattern] matches a [Asset]'s name.
func (ap AssetPattern) match(asset Asset) bool {
	return ap.re.Match([]byte(asset.name))
}

// AssetPatternList is a list of [AssetPattern].
type AssetPatternList []AssetPattern

// compileAssetPatternList compiles exprs as a list of regular expression and return [AssetPatternList].
// exprs must be a list of regular expression.
func compileAssetPatternList(exprs []string) (AssetPatternList, error) {
	apl := AssetPatternList{}
	for _, expr := range exprs {
		p, err := compilePattern(expr)
		if err != nil {
			return nil, err
		}
		apl = append(apl, AssetPattern(p))
	}
	return apl, nil
}

// AssetContent represents a GitHub release asset content.
type AssetContent []byte

// execBinaryContent returns [ExecBinaryContent] in [AssetContent].
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
		case "application/octet-stream":
			return ExecBinaryContent(b.Bytes()), nil
		case "application/x-tar":
			r, err = newExecBinaryReaderFromTar(&b)
		case "application/zip":
			r, err = newExecBinaryReaderFromZip(&b)
		case "application/gzip":
			r, err = gzip.NewReader(&b)
		case "application/x-xz":
			r, err = xz.NewReader(&b)
		default:
			r, err = nil, fmt.Errorf("%w: %s", ErrUnsupportedMIME, mime.String())
		}

		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrGettingExecBinaryContentFailure, err)
		}

		b.Reset()
		if _, err := b.ReadFrom(r); err != nil {
			return nil, err
		}
	}
}

// AssetRepository is a repository for [Asset] and [AssetContent].
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository returns a new [AssetRepository] object.
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
