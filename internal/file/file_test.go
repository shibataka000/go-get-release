package file

import (
	"testing"
)

func TestHasExt(t *testing.T) {
	tests := []struct {
		description string
		name        string
		exts        []string
		hasExt      bool
	}{
		{
			description: "a.exe",
			name:        "a.exe",
			exts:        []string{".exe"},
			hasExt:      true,
		},
		{
			description: "a.exe_does_not_zip",
			name:        "a.exe",
			exts:        []string{".zip"},
			hasExt:      false,
		},
		{
			description: "a",
			name:        "a",
			exts:        []string{""},
			hasExt:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := HasExt(tt.name, tt.exts)
			if actual != tt.hasExt {
				t.Fatalf("Expected is %v but actual is %v", tt.hasExt, actual)
			}
		})
	}
}

func TestIsArchive(t *testing.T) {
	tests := []struct {
		name       string
		isArchived bool
	}{
		{
			name:       "a",
			isArchived: false,
		},
		{
			name:       "a.zip",
			isArchived: true,
		},
		{
			name:       "a.tar.gz",
			isArchived: true,
		},
		{
			name:       "a-v1.0.0.tar.gz",
			isArchived: true,
		},
		{
			name:       "a.tar.gz.deb",
			isArchived: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsArchived(tt.name)
			if actual != tt.isArchived {
				t.Fatalf("Expected is %v but actual is %v", tt.isArchived, actual)
			}
		})
	}
}

func TestIsExecBinary(t *testing.T) {
	tests := []struct {
		name         string
		isExecBinary bool
	}{
		{
			name:         "a",
			isExecBinary: true,
		},
		{
			name:         "a.exe",
			isExecBinary: true,
		},
		{
			name:         "a.tar.gz",
			isExecBinary: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsExecBinary(tt.name)
			if actual != tt.isExecBinary {
				t.Fatalf("Expected is %v but actual is %v", tt.isExecBinary, actual)
			}
		})
	}
}
