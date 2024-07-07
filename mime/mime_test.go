package mime

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMimeDetect(t *testing.T) {
	tests := []struct {
		name string
		mime MIME
	}{
		{
			name: "a.gz",
			mime: Gzip,
		},
		{
			name: "a.tar.gz",
			mime: Gzip,
		},
		{
			name: "a.tgz",
			mime: CompressedTar,
		},
		{
			name: "a",
			mime: OctedStream,
		},
		{
			name: "a.linux-amd64",
			mime: OctedStream,
		},
		{
			name: "a.tar",
			mime: Tar,
		},
		{
			name: "a.xz",
			mime: Xz,
		},
		{
			name: "a.tar.xz",
			mime: Xz,
		},
		{
			name: "a.zip",
			mime: Zip,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.mime, Detect(tt.name))
		})
	}
}

func TestMimeIsCompressed(t *testing.T) {
	tests := []struct {
		name         string
		mime         MIME
		isCompressed bool
	}{
		{
			name:         string(Gzip),
			mime:         Gzip,
			isCompressed: true,
		},
		{
			name:         string(OctedStream),
			mime:         OctedStream,
			isCompressed: false,
		},
		{
			name:         string(Tar),
			mime:         Tar,
			isCompressed: false,
		},
		{
			name:         string(Xz),
			mime:         Xz,
			isCompressed: true,
		},
		{
			name:         string(Zip),
			mime:         Zip,
			isCompressed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.isCompressed, tt.mime.IsCompressed())
		})
	}
}

func TestMimeIsOctedStream(t *testing.T) {
	tests := []struct {
		name          string
		mime          MIME
		isOctedStream bool
	}{
		{
			name:          string(OctedStream),
			mime:          OctedStream,
			isOctedStream: true,
		},
		{
			name:          string(Gzip),
			mime:          Gzip,
			isOctedStream: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.isOctedStream, tt.mime.IsOctetStream())
		})
	}
}
