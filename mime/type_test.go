package mime

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		name     string
		mimetype Type
	}{
		{
			name:     "a.gz",
			mimetype: Gzip,
		},
		{
			name:     "a.tar.gz",
			mimetype: Gzip,
		},
		{
			name:     "a.tgz",
			mimetype: CompressedTar,
		},
		{
			name:     "a",
			mimetype: OctedStream,
		},
		{
			name:     "a.linux-amd64",
			mimetype: OctedStream,
		},
		{
			name:     "a.tar",
			mimetype: Tar,
		},
		{
			name:     "a.xz",
			mimetype: Xz,
		},
		{
			name:     "a.tar.xz",
			mimetype: Xz,
		},
		{
			name:     "a.zip",
			mimetype: Zip,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.mimetype, Detect(tt.name))
		})
	}
}

func TestTypeIsCompressed(t *testing.T) {
	tests := []struct {
		name         string
		mimetype     Type
		isCompressed bool
	}{
		{
			name:         string(Gzip),
			mimetype:     Gzip,
			isCompressed: true,
		},
		{
			name:         string(OctedStream),
			mimetype:     OctedStream,
			isCompressed: false,
		},
		{
			name:         string(Tar),
			mimetype:     Tar,
			isCompressed: false,
		},
		{
			name:         string(Xz),
			mimetype:     Xz,
			isCompressed: true,
		},
		{
			name:         string(Zip),
			mimetype:     Zip,
			isCompressed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.isCompressed, tt.mimetype.IsCompressed())
		})
	}
}

func TestTypeIsOctedStream(t *testing.T) {
	tests := []struct {
		name          string
		mimetype      Type
		isOctedStream bool
	}{
		{
			name:          string(OctedStream),
			mimetype:      OctedStream,
			isOctedStream: true,
		},
		{
			name:          string(Gzip),
			mimetype:      Gzip,
			isOctedStream: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.isOctedStream, tt.mimetype.IsOctetStream())
		})
	}
}
