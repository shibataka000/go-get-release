package archive

import (
	"path/filepath"
	"strings"
)

// IsArchived return true if file is archived
func IsArchived(filename string) bool {
	return hasExt(filename, []string{".tar", ".tgz", ".tbz2", ".txz", ".zip", ".lzh", ".rar", ".7z"}) || hasExt(trimExt(filename), []string{".tar"})
}

// IsCompressed return true if file is compressed
func IsCompressed(filename string) bool {
	return hasExt(filename, []string{".gz", ".tgz", ".bz2", ".tbz2", ".xz", ".txz", ".zip", ".lzh", ".rar", ".7z"})
}

// hasExt return true if 'name' have specific extension which is in 'exts'
func hasExt(name string, exts []string) bool {
	for _, ext := range exts {
		if filepath.Ext(name) == ext {
			return true
		}
	}
	return false
}

// trimExt return filepath trimmed extension
func trimExt(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path))
}
