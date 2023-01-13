package release

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReleaseSemVer(t *testing.T) {
	tests := []struct {
		id     int64
		tag    string
		semver string
		err    error
	}{
		{
			id:     0,
			tag:    "v1.2.3",
			semver: "1.2.3",
		},
		{
			id:     0,
			tag:    "1.2.3",
			semver: "1.2.3",
		},
		{
			id:     0,
			tag:    "x.y.z",
			semver: "",
			err:    fmt.Errorf("x.y.z is not valid semver"),
		},
	}

	for _, tt := range tests {
		name := tt.tag
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			release := New(tt.id, tt.tag)
			semver, err := release.SemVer()
			if tt.err == nil {
				assert.NoError(err)
				assert.Equal(tt.semver, semver)
			} else {
				assert.EqualError(err, tt.err.Error())
			}
		})
	}

}
