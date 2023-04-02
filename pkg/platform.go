package pkg

// Platform is pair of OS and Arch.
type Platform struct {
	OS   string
	Arch string
}

// NewPlatform return new platform instance.
func NewPlatform(os string, arch string) Platform {
	return Platform{
		OS:   os,
		Arch: arch,
	}
}

// Equals return true if p and q are same platform.
func (p Platform) Equals(q Platform) bool {
	return p.OS == q.OS && p.Arch == q.Arch
}
