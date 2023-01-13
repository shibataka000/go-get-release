package asset

// Content of executable binary in GitHub asset.
type ExecBinary struct {
	body []byte
}

// NewExecBinary return content of executable binary in GitHub asset.
func NewExecBinary(body []byte) ExecBinary {
	return ExecBinary{
		body: body,
	}
}

// Bytes return bytes of executable binary.
func (eb ExecBinary) Bytes() []byte {
	return eb.body
}
