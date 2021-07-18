package file

import "path/filepath"

// HasExt return true if 'name' have specific extension which is in 'exts'
func HasExt(name string, exts []string) bool {
	for _, ext := range exts {
		if filepath.Ext(name) == ext {
			return true
		}
	}
	return false
}

// IsArchived check file name and return true if it is archive file
func IsArchived(name string) bool {
	return HasExt(name, []string{".tar", ".gz", ".tgz", ".bz2", ".tbz", ".Z", ".zip", ".bz2", ".lzh", ".7z", ".gz", ".rar", ".cab", ".afz"})
}

// IsExecBinary check file name and return true if it is executable binary
func IsExecBinary(name string) bool {
	return HasExt(name, []string{"", ".exe"})
}
