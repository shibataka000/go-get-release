package mime

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMimeIs(t *testing.T) {
	tests := []struct {
		name          string
		mime          MIME
		isArchived    bool
		isCompressed  bool
		isOctedStream bool
	}{
		{
			name:          string(Gz),
			mime:          Gz,
			isArchived:    false,
			isCompressed:  true,
			isOctedStream: false,
		},
		{
			name:          string(Tar),
			mime:          Tar,
			isArchived:    true,
			isCompressed:  false,
			isOctedStream: false,
		},
		{
			name:          string(Xz),
			mime:          Xz,
			isArchived:    false,
			isCompressed:  true,
			isOctedStream: false,
		},
		{
			name:          string(Zip),
			mime:          Zip,
			isArchived:    true,
			isCompressed:  true,
			isOctedStream: false,
		},
		{
			name:          string(OctedStream),
			mime:          OctedStream,
			isArchived:    false,
			isCompressed:  false,
			isOctedStream: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.isArchived, tt.mime.IsArchived())
			require.Equal(tt.isCompressed, tt.mime.IsCompressed())
			require.Equal(tt.isOctedStream, tt.mime.IsOctetStream())
		})
	}
}
