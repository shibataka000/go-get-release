package archive

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
)

// Extract archived or compressed file and pick up specific file.
func Extract(src, dst, file string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	switch {
	// archived and compressed
	case strings.HasSuffix(src, ".zip"):
		return extractZip(src, out, file)
	case strings.HasSuffix(src, ".tar.gz") || strings.HasSuffix(src, ".tgz"):
		return extractTarGz(in, out, file)
	case strings.HasSuffix(src, ".tar.xz") || strings.HasSuffix(src, ".txz"):
		return extractTarXz(in, out, file)
	// compressed only
	case strings.HasSuffix(src, ".gz"):
		return extractGz(in, out, file)
	// unsupported format
	default:
		return fmt.Errorf("unsupported archiving or compression format: %s", src)
	}
}

// Extract zip file.
func extractZip(src string, out io.Writer, target string) error {
	rc, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer rc.Close()

	for _, f := range rc.File {
		if !f.FileInfo().IsDir() && filepath.Base(f.Name) == target {
			in, err := f.Open()
			if err != nil {
				return err
			}
			defer in.Close()

			_, err = io.Copy(out, in)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("file '%s' is not found in zip file", target)
}

// Extract tar file.
func extractTar(in io.Reader, out io.Writer, target string) error {
	tarIn := tar.NewReader(in)

	for {
		header, err := tarIn.Next()
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
			if filepath.Base(header.Name) == target {
				_, err := io.Copy(out, tarIn)
				if err != nil {
					return err
				}
				return nil
			}
		default:
			return fmt.Errorf("fail to extract tarball: %s %v", header.Name, header.Typeflag)
		}
	}

	return fmt.Errorf("file '%s' is not found in tarball", target)
}

// Extract .tar.gz file.
func extractTarGz(in io.Reader, out io.Writer, target string) error {
	gzipIn, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gzipIn.Close()

	return extractTar(gzipIn, out, target)
}

// Extract .tar.xz file
func extractTarXz(in io.Reader, out io.Writer, target string) error {
	xzIn, err := xz.NewReader(in)
	if err != nil {
		return err
	}

	return extractTar(xzIn, out, target)
}

// Extract gzip file, which is not archived.
func extractGz(in io.Reader, out io.Writer, target string) error {
	gzipIn, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gzipIn.Close()

	_, err = io.Copy(out, gzipIn)
	return err
}
