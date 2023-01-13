package application

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
)

// Command to install executable binary from GitHub release asset.
type Command struct {
	owner       string
	repo        string
	tag         string
	os          string
	arch        string
	installDir  string
	interactive bool
}

// NewCommand return new command instance to install executable binary from GitHub release asset.
func NewCommand(owner string, repo string, tag string, os string, arch string, installDir string, interactive bool) Command {
	return Command{
		owner:       owner,
		repo:        repo,
		tag:         tag,
		os:          os,
		arch:        arch,
		installDir:  installDir,
		interactive: interactive,
	}
}

// NewCommandFromQuery return new command instance from query to search GitHub repository.
func NewCommandFromQuery(query string, os string, arch string, installDir string, interactive bool) (Command, error) {
	re := regexp.MustCompile(`([^/=]+/)?([^/=]+)(=[^/=]+)?`)
	submatch := re.FindStringSubmatch(query)
	if submatch == nil || len(submatch) != 4 {
		return Command{}, fmt.Errorf("%s is not valid query", query)
	}
	owner := strings.TrimSuffix(submatch[1], "/")
	repo := submatch[2]
	tag := strings.TrimPrefix(submatch[3], "=")
	return NewCommand(owner, repo, tag, os, arch, installDir, interactive), nil
}

// Owner of GitHub repository.
func (c Command) Owner() string {
	return c.owner
}

// Repo return GitHub repository name.
func (c Command) Repo() string {
	return c.repo
}

// Tag return tag of GitHub release.
func (c Command) Tag() string {
	return c.tag
}

// InstallDir to install executable binary to.
func (c Command) InstallDir() string {
	return c.installDir
}

// hasOwner return true if query contains GitHub repository owner.
func (c Command) hasOwner() bool {
	return c.owner != ""
}

// HasOwner return true if query contains GitHub release tag.
func (c Command) hasTag() bool {
	return c.tag != ""
}

// searchRepositoryQuery return query to search GitHub repository.
func (c Command) searchRepositoryQuery() string {
	return c.repo
}

// platform which install executable binary to.
func (c Command) platform() platform.Platform {
	return platform.New(c.os, c.arch)
}

// isInteractive return true if interactive mode is on.
// In interactive mode, some prompt are shown.
func (c Command) isInteractive() bool {
	return c.interactive
}
