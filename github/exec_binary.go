package github

type ExecBinary struct {
	name string
}

func newExecBinary(name string) ExecBinary {
	return ExecBinary{
		name: name,
	}
}
