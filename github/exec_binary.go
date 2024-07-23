package github

import (
	"fmt"

	dist "github.com/shibataka000/go-get-release/distribution"
)

type ExecBinary struct {
	name string
}

func newExecBinary(name string) ExecBinary {
	return ExecBinary{
		name: name,
	}
}

func newExecBinaryWithOS(name string, os dist.OS) ExecBinary {
	if os == "windows" {
		n := fmt.Sprintf("%s.exe", name)
		return newExecBinary(n)
	}
	return newExecBinary(name)
}

type ExecBinaryTemplate struct {
	name string
}

func newExecBinaryTemplate(name string) ExecBinaryTemplate {
	return ExecBinaryTemplate{
		name: name,
	}
}

func (b ExecBinaryTemplate) execute(os dist.OS) ExecBinary {
	return newExecBinaryWithOS(b.name, os)
}
