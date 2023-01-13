package platform

// Platform is pair of GOOS and GOARCH.
type Platform struct {
	os   string
	arch string
}

// New platform instance.
func New(os string, arch string) Platform {
	return Platform{
		os:   os,
		arch: arch,
	}
}

// OS return GOOS.
func (p Platform) OS() string {
	return p.os
}

// Arch return GOARCH.
func (p Platform) Arch() string {
	return p.arch
}

// Equal return true if p and q are same platform.
func (p Platform) Equal(q Platform) bool {
	return p.OS() == q.OS() && p.Arch() == q.Arch()
}
