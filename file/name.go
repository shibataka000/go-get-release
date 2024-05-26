package file

import "path/filepath"

// Name is file name.
type Name string

// Ext return file name extension.
func (n Name) Ext() string {
	return filepath.Ext(string(n))
}
