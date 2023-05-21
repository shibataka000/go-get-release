package pkg

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

// Package.
type Package struct {
	Repository Repository
	Release    Release
	Asset      Asset
	ExecBinary ExecBinary
}

// Repository.
type Repository GitHubRepository

// Release.
type Release struct {
	Tag string
}

// Asset.
type Asset GitHubAsset

// ExecBinary.
type ExecBinary struct {
	Name FileName
}

// New package instance.
func New(repo Repository, release Release, asset Asset, execBinary ExecBinary) Package {
	return Package{
		Repository: repo,
		Release:    release,
		Asset:      asset,
		ExecBinary: execBinary,
	}
}

// NewRepository return new repository instance.
func NewRepository(owner string, name string) Repository {
	return Repository{
		Owner: owner,
		Name:  name,
	}
}

// NewRelease return new release instance.
func NewRelease(tag string) Release {
	return Release{
		Tag: tag,
	}
}

// NewAsset return new asset instance.
func NewAsset(downloadURL URL) Asset {
	return Asset{
		DownloadURL: downloadURL,
	}
}

// NewExecBinary return new executable binary instance.
func NewExecBinary(name FileName) ExecBinary {
	return ExecBinary{
		Name: name,
	}
}

// StringToPrompt return string to prompt.
func (p Package) StringToPrompt() string {
	return fmt.Sprintf("Repo:\t%s/%s\nTag:\t%s\nAsset:\t%s\nBinary:\t%s", p.Repository.Owner, p.Repository.Name, p.Release.Tag, p.Asset.DownloadURL.FileName().String(), p.ExecBinary.Name)
}

// SemVer return semver formatted release tag.
// For example, if release tag is "v1.2.3", this return "1.2.3".
func (r Release) SemVer() (string, error) {
	if !semver.IsValid(r.Tag) && !semver.IsValid(fmt.Sprintf("v%s", r.Tag)) {
		return "", fmt.Errorf("%s is not valid semver", r.Tag)
	}
	return strings.TrimLeft(r.Tag, "v"), nil
}
