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

// findKeyWhichHasLongestMatchValue returns key in map which has longest matched value with target.
// If no value was matched with target, this returns defaultKey.
func findKeyWhichHasLongestMatchValue[E ~string](m map[E][]string, target string) E {
	var (
		matchKey   E
		matchValue string
	)
	for key, values := range m {
		for _, value := range values {
			if strings.Contains(target, value) && len(matchValue) < len(value) {
				matchKey = key
				matchValue = value
			}
		}
	}
	return matchKey
}
