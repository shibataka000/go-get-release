package github

import (
	"bytes"
	"compress/gzip"
	"context"
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

// newAsset returns a new GitHub release asset object.
func newAsset(id int64, name string) Asset {
	return Asset{
		id:   id,
		name: name,
	}
}

// AssetList represents a list of GitHub release assets.
type AssetList []Asset

// find a GitHub release asset which matches any of given patterns.
// If two or more assets match, this returns asset which matches prior pattern.
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

type AssetPattern Pattern

func (ap AssetPattern) match(asset Asset) bool {
	return ap.re.Match([]byte(asset.name))
}

type AssetPatternList []AssetPattern

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
			r, err = newExecBinaryReaderFromTar(&b)
		case "application/zip":
			r, err = newExecBinaryReaderFromZip(bytes.NewReader(b.Bytes()), int64(b.Len()))
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

// download a GitHub release asset.
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
