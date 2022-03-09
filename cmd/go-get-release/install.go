package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/cheggaaa/pb"
	"github.com/shibataka000/go-get-release/internal/archive"
	"github.com/shibataka000/go-get-release/internal/github"
)

func install(name, token, goos, goarch, dir string, showPrompt bool) error {
	repo, release, asset, err := findAsset(name, token, goos, goarch)
	if err != nil {
		return err
	}

	if showPrompt {
		fmt.Printf("repo:\t%s/%s\ntag:\t%s\nasset:\t%s\n\n", repo.Owner(), repo.Name(), release.Tag(), asset.Name())
		if !prompter.YN("Are you sure to install release binary from above repository?", true) {
			return nil
		}
		fmt.Println()
	}

	tempDir, err := ioutil.TempDir("", "go-get-release-")
	if err != nil {
		return err
	}

	downloadFilePath := filepath.Join(tempDir, asset.Name())
	err = downloadFile(asset.DownloadURL(), downloadFilePath, showPrompt)
	if err != nil {
		return err
	}

	binaryName, err := asset.BinaryName()
	if err != nil {
		return err
	}

	var downloadBinaryPath string
	if asset.IsArchived() || asset.IsCompressed() {
		err = archive.Extract(downloadFilePath, tempDir)
		if err != nil {
			return err
		}
		var targetFileName string
		if asset.IsArchived() {
			targetFileName = binaryName
		} else {
			filename := filepath.Base(downloadFilePath)
			targetFileName = strings.TrimSuffix(filename, filepath.Ext(filename))
		}
		downloadBinaryPath, err = findFile(tempDir, targetFileName)
		if err != nil {
			return err
		}
	} else {
		downloadBinaryPath = downloadFilePath
	}

	installBinaryPath := filepath.Join(dir, binaryName)
	err = os.Rename(downloadBinaryPath, installBinaryPath)
	if err != nil {
		return err
	}
	err = os.Chmod(installBinaryPath, 0775)
	if err != nil {
		return err
	}

	err = os.RemoveAll(tempDir)
	if err != nil {
		return err
	}
	return nil
}

func findAsset(name, token, goos, goarch string) (github.Repository, github.Release, github.Asset, error) {
	owner, repoName, tag, err := parse(name)
	if err != nil {
		return nil, nil, nil, err
	}

	client, err := github.NewClient(token)
	if err != nil {
		return nil, nil, nil, err
	}

	var repo github.Repository
	if owner != "" {
		repo, err = client.Repository(owner, repoName)
	} else {
		repo, err = client.FindRepository(repoName)
	}
	if err != nil {
		return nil, nil, nil, err
	}

	var release github.Release
	if tag != "" {
		release, err = repo.Release(tag)
	} else {
		release, err = repo.LatestRelease()
	}
	if err != nil {
		return nil, nil, nil, err
	}

	asset, err := release.AssetByPlatform(goos, goarch)
	if err != nil {
		return nil, nil, nil, err
	}

	return repo, release, asset, nil
}

func parse(name string) (string, string, string, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^=]+))?`)
	if re.MatchString(name) {
		match := re.FindStringSubmatch(name)
		return match[2], match[3], match[5], nil
	}
	return "", "", "", fmt.Errorf("parsing package name failed: %s\npackage name should be \"owner/repo=tag\" format", name)
}

func downloadFile(url, filePath string, showProgress bool) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	r := resp.Body
	if showProgress {
		bar := pb.Full.Start64(resp.ContentLength)
		r = bar.NewProxyReader(resp.Body)
	}
	_, err = io.Copy(out, r)
	return err
}

func findFile(dirPath, fileName string) (string, error) {
	filePath := ""
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			filePath = path
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if filePath == "" {
		return "", fmt.Errorf("%s is not found in %s", fileName, dirPath)
	}
	return filePath, nil
}
