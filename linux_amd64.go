package main

var (
	// defaultPatterns are recommended patterns for linux/amd64.
	defaultPatterns = map[string]string{
		// These are recommended patterns for general repository.
		`(?i)^.+/(?P<name>[^\.]+)([\-\._]v?\d+\.\d+\.\d+)?[\-\._]linux([\-\._](amd64|x86_64|x64|64bit))?(\.tar\.gz|\.tar\.xz|\.zip|\.gz|\.tgz)?$`: "{{.name}}",
		`https://github\.com/.+/releases/download/.+/(?P<name>.+)-x86_64-unknown-linux-gnu\.tar\.gz$`:                                             "{{.name}}",

		// These are recommended patterns for specific repository whose release assets are hosted on GitHub.
		// These should start with literals containing host and repository name to avoid conflict with other patterns.
		`https://github\.com/istio/istio/releases/download/.+/istioctl-\d+\.\d+\.\d+-linux-amd64\.tar\.gz$`:    "istioctl",
		`https://github\.com/protocolbuffers/protobuf/releases/download/.+/protoc-\d+\.\d+-linux-x86_64\.zip$`: "protoc",

		// These are recommended patterns for specific repository whose release assets are hosted on server other than GitHub.
		// These should start with literals containing host to avoid conflict with other patterns.
		`https://dl\.k8s\.io/release/.+/bin/linux/amd64/kubectl`:              "kubectl",
		`https://cdn\.teleport\.dev/teleport-v.+-linux-amd64-bin\.tar\.gz`:    "tsh",
		`https://releases\.hashicorp\.com/tfctl/.+/tfctl_.+_linux_amd64\.zip`: "tfctl",
	}
)
