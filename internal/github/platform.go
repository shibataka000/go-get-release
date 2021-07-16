package github

// listGoos return list of GOOS, which are same as result of following command.
// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\1/g" | sort | uniq`
func listGoos() []string {
	return []string{
		"aix",
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"illumos",
		"ios",
		"js",
		"linux",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"windows",
	}
}

// listGoarch return list of GOARCH, which are same as result of following command.
// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\2/g" | sort | uniq`
func listGoarch() []string {
	return []string{
		"386",
		"amd64",
		"arm",
		"arm64",
		"mips",
		"mips64",
		"mips64le",
		"mipsle",
		"ppc64",
		"ppc64le",
		"riscv64",
		"s390x",
		"wasm",
	}
}
