package goxcart

// PlatformGroup denotes a group of specialized gox build configurations.
type PlatformGroup struct {
	// ImageVariant denotes an optional Docker image tag to apply.
	//
	// Example: musl
	ImageTag string

	// OSVariant provides an optional artifact label to apply.
	//
	// Example: musl
	OSVariant string

	// OSs denotes which operating systems are supported by this group.
	// Empty indicates all operating systems known to the Docker image.
	//
	// Example: linux, darwin, windows
	OSs []string

	// Archs denotes which architectures are supported by this group.
	// Empty indicates all architectures known to the Docker image.
	//
	// Example: amd64, 386
	Archs []string

	// LinkerFlags denotes any additional linker flags.
	//
	// Example: '-linkmode external'
	LinkerFlags string
}

// PlatformGroups lists available specialized gox build configurations.
var PlatformGroups = []PlatformGroup{
	// linux-musl
	{
		ImageTag:  "musl",
		OSVariant: "musl",
		OSs:       []string{"linux"},
		Archs:     []string{"amd64"},
	},
	// linux-gnu
	{
		ImageTag:  "glibc",
		OSVariant: "glibc",
		OSs:       []string{"linux"},
		Archs:     []string{"386", "amd64", "arm", "arm64", "mips", "mips64", "mips64le", "mipsle", "ppc64", "ppc64le", "s390x"},
	},
	// dragonfly
	{
		ImageTag: "glibc",
		OSs:      []string{"dragonfly"},
		Archs:    []string{"amd64"},
	},
	// nacl
	{
		ImageTag: "glibc",
		OSs:      []string{"nacl"},
		Archs:    []string{"arm", "amd64p32"},
	},
	// solaris
	{
		ImageTag: "glibc",
		OSs:      []string{"solaris"},
		Archs:    []string{"amd64"},
	},
	// assorted *NIX
	{
		ImageTag: "glibc",
		OSs:      []string{"freebsd", "netbsd", "plan9"},
		Archs:    []string{"386", "amd64", "arm"},
	},
	// PC
	{
		ImageTag: "glibc",
		OSs:      []string{"darwin", "openbsd", "windows"},
		Archs:    []string{"386", "amd64"},
	},
}
