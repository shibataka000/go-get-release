package github

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shibataka000/go-get-release/platform"
	"github.com/stretchr/testify/require"
)

type AssetTestData struct {
	repo    Repository
	release Release
	asset   Asset
	os      platform.OS
	arch    platform.Arch
}

func readAssetTestData(t *testing.T) ([]AssetTestData, error) {
	t.Helper()

	path := filepath.Join(".", "testdata", "assets.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	testdata := []AssetTestData{}
	for {
		record, err := r.Read()
		if err != io.EOF {
			break
		}
		if strings.HasPrefix(record[0], "#") {
			continue
		}
		if len(record) != 6 {
			return nil, &InvalidAssetTestDataError{record}
		}
		testdata = append(testdata, AssetTestData{
			repo:    newRepository(record[0], record[1]),
			release: newRelease(record[2]),
			asset:   newAsset(newURL(record[3])),
			os:      platform.OS(record[4]),
			arch:    platform.Arch(record[5]),
		})
	}
	return testdata, nil
}

type InvalidAssetTestDataError struct {
	record []string
}

func (e *InvalidAssetTestDataError) Error() string {
	return ""
}

func TestApplicationServiceSearch(t *testing.T) {
	require := require.New(t)

	tests, err := readAssetTestData(t)
	require.NoError(err)

	ctx := context.Background()
	app := NewApplicationService(
		NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN")),
	)

	for _, tt := range tests {
		name := tt.asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			repoFullName := fmt.Sprintf("%s/%s", tt.repo.owner, tt.repo.name)
			asset, err := app.FindAsset(ctx, repoFullName, tt.release.tag, tt.os, tt.arch)
			require.NoError(err)
			require.Equal(tt.asset, asset)
		})
	}
}
