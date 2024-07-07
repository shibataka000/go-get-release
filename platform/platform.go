package platform

import (
	"strings"
)

// OS is operating system.
type OS string

// Arch is architecture.
type Arch string

// Detect os/arch from name.
// If os/arch can't be detected, this returns empty string.
func Detect(name string) (OS, Arch) {
	name = strings.ToLower(name)
	os := findKeyWhichHasLongestMatchValue(osKeywords, name)
	arch := findKeyWhichHasLongestMatchValue(archKeywords, name)
	return os, arch
}
