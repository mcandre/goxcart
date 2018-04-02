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
		ImageTag:    "musl",
		OSVariant:   "musl",
		OSs:         []string{"linux"},
		Archs:       []string{"amd64"},
		LinkerFlags: "-linkmode external",
	},
	// linux-gnu
	{
		ImageTag:  "glibc",
		OSVariant: "glibc",
		OSs:       []string{"linux"},
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
	// plan9, solaris
	{
		ImageTag: "glibc",
		OSs:      []string{"plan9", "solaris"},
	},
	// android
	{
		ImageTag: "glibc",
		OSs:      []string{"android"},
		Archs:    []string{"386"},
	},
	// most platforms
	{
		ImageTag: "glibc",
		OSs:      []string{"!linux", "!dragonfly", "!nacl", "!plan9", "!solaris", "!android"},
	},
}
