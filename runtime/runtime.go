package runtime

import (
	"strings"
)

// GOOS is the running program's operating system target.
type GOOS string

// GOARCH is the running program's architecture target.
type GOARCH string

var goosKeywords = map[GOOS][]string{
	"aix":       {"aix"},
	"android":   {"android"},
	"darwin":    {"darwin", "macos", "osx"},
	"dragonfly": {"dragonfly"},
	"freebsd":   {"freebsd"},
	"illumos":   {"illumos"},
	"ios":       {"ios"},
	"js":        {"js"},
	"linux":     {"linux"},
	"netbsd":    {"netbsd"},
	"openbsd":   {"openbsd"},
	"plan9":     {"plan9"},
	"solaris":   {"solaris"},
	"windows":   {"windows", "win", ".exe"},
}

var goarchKeywords = map[GOARCH][]string{
	"386":      {"386", "x86_32", "32bit", "win32"},
	"686":      {"686"},
	"amd64":    {"amd64", "x86_64", "64bit", "win64"},
	"arm":      {"arm"},
	"arm64":    {"arm64", "aarch64", "aarch_64"},
	"mips":     {"mips"},
	"mips64":   {"mips64"},
	"mips64le": {"mips64le"},
	"mipsle":   {"mipsle"},
	"ppc64":    {"ppc64"},
	"ppc64le":  {"ppc64le", "ppcle_64"},
	"riscv64":  {"riscv64"},
	"s390x":    {"s390x", "s390"},
	"wasm":     {"wasm"},
}

func Guess(name string) (GOOS, GOARCH) {
	goos := findKeyWhichHasLongestMatchValue(goosKeywords, name, "unknown")
	goarch := findKeyWhichHasLongestMatchValue(goarchKeywords, name, "amd64")
	return goos, goarch
}

// findKeyWhichHasLongestMatchValue return key in map which has longest matched value.
func findKeyWhichHasLongestMatchValue[E ~string](m map[E][]string, target string, defaultKey E) E {
	var matchKey E = ""
	var matchValue string = ""
	for key, values := range m {
		for _, value := range values {
			if strings.Contains(target, value) && len(matchValue) < len(value) {
				matchKey = key
				matchValue = value
			}
		}
	}
	if matchKey == "" {
		return defaultKey
	}
	return matchKey
}
