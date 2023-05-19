package pkg

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ulikunitz/xz"
	"golang.org/x/exp/slices"
)

// File.
type File struct {
	Name FileName
	Body []byte
}

// FileName is file name.
type FileName string

// NewFile return new file instance.
func NewFile(name FileName, body []byte) File {
	return File{
		Name: name,
		Body: body,
	}
}

// Extract compressed file.
func (f File) Extract() (File, error) {
	fileName := f.Name.Normalize()
	src := bytes.NewReader(f.Body)
	dst := new(bytes.Buffer)
	var err error

	switch fileName.Ext() {
	case ".gz":
		err = extractGzip(dst, src)
	case ".xz":
		err = extractXz(dst, src)
	default:
		err = fmt.Errorf("unsupported file format: %s", fileName.Ext())
	}

	if err != nil {
		return File{}, err
	}

	return NewFile(fileName.TrimExt(), dst.Bytes()), nil
}

// FindFile find file in archived file and return it.
func (f File) FindFile(target FileName) (File, error) {
	fileName := f.Name.Normalize()
	src := bytes.NewReader(f.Body)
	dst := new(bytes.Buffer)
	var err error

	switch fileName.Ext() {
	case ".tar":
		err = copyFileInTar(dst, src, target)
	case ".zip":
		err = copyFileInZip(dst, src, target)
	default:
		err = fmt.Errorf("unsupported file format: %s", fileName.Ext())
	}

	if err != nil {
		return File{}, err
	}

	return NewFile(target, dst.Bytes()), nil
}

// extractGzip extract src as gzip file and copy it to dst.
func extractGzip(dst io.Writer, src io.Reader) error {
	gzipSrc, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer gzipSrc.Close()

	_, err = io.Copy(dst, gzipSrc)
	return err
}

// extractXz extract src as xz file and copy it to dst.
func extractXz(dst io.Writer, src io.Reader) error {
	xzSrc, err := xz.NewReader(src)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, xzSrc)
	return err
}

// copyFileInTar find a file in tarball and copy it to dst.
func copyFileInTar(dst io.Writer, src io.Reader, target FileName) error {
	tarSrc := tar.NewReader(src)

	for {
		header, err := tarSrc.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			// do nothing
		case tar.TypeReg:
			if filepath.Base(header.Name) == target.String() {
				_, err := io.Copy(dst, tarSrc)
				return err
			}
		default:
			return fmt.Errorf("unexpected typeflag: %v", header.Typeflag)
		}
	}

	return fmt.Errorf("file '%s' was not found in tarball", target)
}

// copyFileInZip find a file in zip file and copy it to dst.
func copyFileInZip(dst io.Writer, src io.Reader, target FileName) error {
	tempFile, err := os.CreateTemp("", "*.zip")
	if err != nil {
		return err
	}
	_, err = io.Copy(tempFile, src)
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	rc, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return err
	}
	defer rc.Close()

	for _, f := range rc.File {
		if !f.FileInfo().IsDir() && filepath.Base(f.Name) == target.String() {
			fileIn, err := f.Open()
			if err != nil {
				return err
			}
			defer fileIn.Close()

			_, err = io.Copy(dst, fileIn)
			return err
		}
	}

	return fmt.Errorf("file '%s' was not found in zip file", target)
}

// NewFileName return new FileName instance.
func NewFileName(name string) FileName {
	return FileName(name)
}

// String return string typed file name.
func (f FileName) String() string {
	return string(f)
}

// Ext return file name extension.
func (f FileName) Ext() string {
	return filepath.Ext(f.String())
}

// TrimExt trim file name extension and return new FileName insntace.
func (f FileName) TrimExt() FileName {
	name := strings.TrimSuffix(f.String(), f.Ext())
	return NewFileName(name)
}

// Normalize file name.
// File name extension '.tgz' and '.txz' will be replaced to '.tar.gz' and '.tar.xz'.
func (f FileName) Normalize() FileName {
	switch f.Ext() {
	case ".tgz":
		return f.TrimExt() + ".tar.gz"
	case ".txz":
		return f.TrimExt() + ".tar.xz"
	default:
		return f
	}
}

// IsExecBinary return true if file is executable binary.
func (f FileName) IsExecBinary() bool {
	exts := []string{"", ".exe", ".linux", ".darwin", ".linux-amd64", ".darwin-amd64"}
	return slices.Contains(exts, f.Normalize().Ext())
}

// IsCompressed return true if file is compressed.
func (f FileName) IsCompressed() bool {
	exts := []string{".gz", ".xz", ".zip"}
	return slices.Contains(exts, f.Normalize().Ext())
}

// IsCompressed return true if file is archived.
func (f FileName) IsArchived() bool {
	exts := []string{".zip"}
	return slices.Contains(exts, f.Normalize().Ext()) || f.IsTarBall()
}

// IsTarBall return true if file is tarball.
func (f FileName) IsTarBall() bool {
	exts := []string{".tar"}
	return slices.Contains(exts, f.Normalize().Ext()) || slices.Contains(exts, f.Normalize().TrimExt().Ext())
}

// Platform return platform guessed by file name.
func (f FileName) Platform() (Platform, error) {
	os, err := f.os()
	if err != nil {
		return Platform{}, err
	}
	arch, err := f.arch()
	if err != nil {
		return Platform{}, err
	}
	return NewPlatform(os, arch), nil

}

// os return os guessed by file name.
func (f FileName) os() (string, error) {
	// These are listed by following command.
	// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\1/g" | sort | uniq`
	platforms := map[string][]string{
		"aix":       {"aix"},
		"android":   {"android"},
		"darwin":    {"darwin", "macos", "osx"},
		"dragonfly": {"dragonfly"},
		"freebsd":   {"freebsd"},
		"illumos":   {"illumos"},
		"ios":       {"ios"},
		"js":        {"js"},
		"linux":     {"linux"},
		"netbsd":    {"netbsd"},
		"openbsd":   {"openbsd"},
		"plan9":     {"plan9"},
		"solaris":   {"solaris"},
		"windows":   {"windows", "win", ".exe"},
	}
	lowner := strings.ToLower(f.String())
	return findKeyWhichHasLongestMatchValue(platforms, lowner)
}

// arch return arch guessed by file name.
// 'amd64' is returned by default.
func (f FileName) arch() (string, error) {
	// These are listed by following command.
	// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\2/g" | sort | uniq`
	platforms := map[string][]string{
		"386":      {"386", "x86_32", "32bit", "win32"},
		"amd64":    {"amd64", "x86_64", "64bit", "win64"},
		"arm":      {"arm"},
		"arm64":    {"arm64", "aarch64", "aarch_64"},
		"mips":     {"mips"},
		"mips64":   {"mips64"},
		"mips64le": {"mips64le"},
		"mipsle":   {"mipsle"},
		"ppc64":    {"ppc64"},
		"ppc64le":  {"ppc64le", "ppcle_64"},
		"riscv64":  {"riscv64"},
		"s390x":    {"s390x", "s390"},
		"wasm":     {"wasm"},
	}
	lowner := strings.ToLower(f.String())
	arch, err := findKeyWhichHasLongestMatchValue(platforms, lowner)
	if err != nil {
		return "amd64", nil
	}
	return arch, nil
}

// findKeyWhichHasLongestMatchValue return key in map which has longest matched value.
func findKeyWhichHasLongestMatchValue(m map[string][]string, value string) (string, error) {
	values := []string{}
	for _, vs := range m {
		values = append(values, vs...)
	}
	sort.Slice(values, func(i, j int) bool { return len(values[i]) > len(values[j]) })

	longestMatchValue := ""
	found := false
	for _, v := range values {
		if strings.Contains(value, v) {
			longestMatchValue = v
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("no value was matched")
	}

	for k, vs := range m {
		if slices.Contains(vs, longestMatchValue) {
			return k, nil
		}
	}
	return "", fmt.Errorf("no key was found")
}
