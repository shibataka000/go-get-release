package distribution

// osKeywords is a map whose key is os and whose values are its keywords.
// These are listed by following command.
// `go tool dist list -json | jq -r .[].GOOS | sort | uniq`
var osKeywords = map[OS][]string{
	"aix":       {"aix"},
	"android":   {"android"},
	"darwin":    {"darwin", "macos", "osx"},
	"dragonfly": {"dragonfly"},
	"freebsd":   {"freebsd"},
	"illumos":   {"illumos"},
	"ios":       {"ios"},
	"js":        {"js"},
	"linux":     {"linux", "alpine"},
	"netbsd":    {"netbsd"},
	"openbsd":   {"openbsd"},
	"plan9":     {"plan9"},
	"solaris":   {"solaris"},
	"wasip1":    {"wasip1"},
	"windows":   {"windows", "win", ".exe"},
}

// archKeywords is a map whose key is arch and whose values are its keywords.
// These are listed by following command.
// `go tool dist list -json | jq -r .[].GOARCH | sort | uniq`
var archKeywords = map[Arch][]string{
	"386":      {"386", "x86_32", "32bit", "win32", "i686"},
	"amd64":    {"amd64", "x86_64", "64bit", "win64"},
	"arm":      {"arm"},
	"arm64":    {"arm64", "aarch64", "aarch_64"},
	"loong64":  {"loong64"},
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
