package github

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/shibataka000/go-get-release/mime"
	"github.com/shibataka000/go-get-release/platform"
)

// AssetTestCaseEntry is a test case about a GitHub release asset.
type AssetTestCaseEntry struct {
	repo          Repository
	release       Release
	asset         Asset
	os            platform.OS
	arch          platform.Arch
	mime          mime.MIME
	hasExecBinary bool
}

type AssetTestCaseTable []AssetTestCaseEntry

func (table AssetTestCaseTable) repositories() []Repository {
	repos := []Repository{}
	for _, t := range table {
		if !slices.Contains(repos, t.repo) {
			repos = append(repos, t.repo)
		}
	}
	return repos
}

func (table AssetTestCaseTable) releases(repo Repository) []Release {
	releases := []Release{}
	for _, t := range table {
		if t.repo == repo && !slices.Contains(releases, t.release) {
			releases = append(releases, t.release)
		}
	}
	return releases
}

func (table AssetTestCaseTable) assets(repo Repository, release Release) AssetList {
	assets := AssetList{}
	for _, t := range table {
		if t.repo == repo && t.release == release && !slices.Contains(assets, t.asset) {
			assets = append(assets, t.asset)
		}
	}
	return assets
}

// readAssetTestData return a list of test data about a GitHub release asset.
func readAssetTestData(t *testing.T) (AssetTestCaseTable, error) {
	t.Helper()

	path := filepath.Join(".", "testdata", "assets.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)

	tests := []AssetTestCaseEntry{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if strings.HasPrefix(record[0], "#") {
			continue
		}

		if len(record) != 8 {
			return nil, &InvalidAssetTestCaseError{record}
		}

		repo := newRepository(record[0], record[1])
		release := newRelease(record[2])
		downloadURL, err := url.Parse(record[3])
		if err != nil {
			return nil, err
		}
		asset := newAsset(downloadURL)
		os := platform.OS(record[4])
		arch := platform.Arch(record[5])
		mime := mime.MIME(record[6])
		hasExecBinary, err := strconv.ParseBool(record[7])
		if err != nil {
			return nil, err
		}

		tests = append(tests, AssetTestCaseEntry{
			repo:          repo,
			release:       release,
			asset:         asset,
			os:            os,
			arch:          arch,
			mime:          mime,
			hasExecBinary: hasExecBinary,
		})
	}
	return tests, nil
}

// InvalidAssetTestCaseError is error raised when asset test case is invalid.
type InvalidAssetTestCaseError struct {
	record []string
}

// Error returns an error message.
func (e *InvalidAssetTestCaseError) Error() string {
	return fmt.Sprintf("Asset test case is invalid: %v", e.record)
}
