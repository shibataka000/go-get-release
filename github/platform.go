package github

// OS.
type OS string

// Arch.
type Arch string

// Platform is pair of os and arch.
type Platform struct {
	OS   OS
	Arch Arch
}

// UnknownOS is unknown os.
var UnknownOS OS = "unknown"

// UnknownArch is unknown arch.
var UnknownArch Arch = "unknown"

// NewPlatform return new platform instance.
func NewPlatform(os OS, arch Arch) Platform {
	return Platform{
		OS:   os,
		Arch: arch,
	}
}

// Equal return true if two platforms are same.
func (p Platform) Equal(other Platform) bool {
	return p.OS == other.OS && p.Arch == other.Arch
}
