package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/cheggaaa/pb/v3"
	"github.com/shibataka000/go-get-release/pkg/archive"
	"github.com/shibataka000/go-get-release/pkg/github"
)

func install(name, token, goos, goarch, installDir string, showPrompt bool) error {
	ctx := context.Background()

	// Find asset
	client, err := github.NewClient(ctx, token)
	if err != nil {
		return err
	}
	splitted := strings.Split(name+"=latest", "=")
	asset, err := client.FindAsset(ctx, splitted[0], splitted[1], goos, goarch)
	if err != nil {
		return err
	}
	assetName := filepath.Base(asset.DownloadURL)

	// Confirmation
	if showPrompt {
		fmt.Printf("repo:\t%s/%s\ntag:\t%s\nasset:\t%s\n\n", asset.Owner, asset.Repo, asset.Tag, assetName)
		if !prompter.YN("Are you sure to install release binary from above repository?", true) {
			return nil
		}
		fmt.Println()
	}

	// Download file
	tempDir, err := os.MkdirTemp("", "go-get-release-")
	if err != nil {
		return err
	}
	downloadFilePath := filepath.Join(tempDir, assetName)
	err = downloadFile(asset.DownloadURL, downloadFilePath, showPrompt)
	if err != nil {
		return err
	}
	downloadBinaryPath := downloadFilePath

	// Extract
	if archive.IsArchived(assetName) || archive.IsCompressed(assetName) {
		downloadBinaryPath = filepath.Join(tempDir, asset.BinaryName)
		err := archive.Extract(downloadFilePath, downloadBinaryPath, asset.BinaryName)
		if err != nil {
			return err
		}
	}

	// Move release binary to install directory
	installBinaryPath := filepath.Join(installDir, asset.BinaryName)
	err = os.Rename(downloadBinaryPath, installBinaryPath)
	if err != nil {
		return err
	}
	err = os.Chmod(installBinaryPath, 0775)
	if err != nil {
		return err
	}

	// Clean up
	err = os.RemoveAll(tempDir)
	if err != nil {
		return err
	}

	return nil
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
