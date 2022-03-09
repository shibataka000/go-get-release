package archive

import (
	"testing"
)

func TestIsArchivedOrCompressed(t *testing.T) {
	tests := []struct {
		filename   string
		archived   bool
		compressed bool
	}{
		// binary
		{
			filename:   "aaa",
			archived:   false,
			compressed: false,
		},
		// archiving only
		{
			filename:   "aaa.tar",
			archived:   true,
			compressed: false,
		},
		// compression only
		{
			filename:   "aaa.gz",
			archived:   false,
			compressed: true,
		},
		{
			filename:   "aaa.bz2",
			archived:   false,
			compressed: true,
		},
		{
			filename:   "aaa.xz",
			archived:   false,
			compressed: true,
		},
		// archiving and compression (tar)
		{
			filename:   "aaa.tar.gz",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.tar.bz2",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.tar.xz",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.tgz",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.tbz2",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.txz",
			archived:   true,
			compressed: true,
		},
		// archiving and compression
		{
			filename:   "aaa.zip",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.lzh",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.rar",
			archived:   true,
			compressed: true,
		},
		{
			filename:   "aaa.7z",
			archived:   true,
			compressed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if IsArchived(tt.filename) != tt.archived {
				t.Fatalf("IsArchived: expected is %v but actual is %v", tt.archived, IsArchived(tt.filename))
			}
			if IsCompressed(tt.filename) != tt.compressed {
				t.Fatalf("IsCompressed: expected is %v but actual is %v", tt.compressed, IsCompressed(tt.filename))
			}
		})
	}
}
