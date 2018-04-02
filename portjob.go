package goxcart

// PortJob configures a worker build.
type PortJob struct {
	// PlatformGroup denotes a specialized gox build parameter set.
	PlatformGroup PlatformGroup

	// CurrentWorkingDirectory denotes the host shell working directory.
	CurrentWorkingDirectory string

	// Environment denotes the host shell environment.
	Environment []string

	// SourcePath denotes the Go source location within Docker.
	SourcePath string
}
