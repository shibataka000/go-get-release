package github

import (
	"fmt"
	"testing"
)

func TestIsSpecialAsset(t *testing.T) {
	tests := []struct {
		owner          string
		repo           string
		isSpecialAsset bool
	}{
		{
			owner:          "hashicorp",
			repo:           "terraform",
			isSpecialAsset: true,
		},
		{
			owner:          "shibataka000",
			repo:           "go-get-release",
			isSpecialAsset: false,
		},
	}

	for _, tt := range tests {
		description := fmt.Sprintf("%s/%s", tt.owner, tt.repo)
		t.Run(description, func(t *testing.T) {
			actual := isSpecialAsset(tt.owner, tt.repo)
			if tt.isSpecialAsset != actual {
				t.Errorf("Expected is %t but actual is %t", tt.isSpecialAsset, actual)
				return
			}
		})
	}
}
