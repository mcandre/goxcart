package goxcart

import (
	"errors"
)

// PortConfig parameterizes builds.
type PortConfig struct {
	// OutputDirectory denotes the root of the port artifact tree.
	//
	// Example: bin
	OutputDirectory string

	// Banner provides a label for the artifact tree (required).
	//
	// Example: hextime-0.0.1
	Banner string

	// Repository denotes the Go repository namespace of a project.
	// When empty, the current working directory is used instead.
	//
	// Example: github.com/mcandre/go-hextime
	Repository string

	// Image denotes the Docker image used for building ports.
	//
	// Example: mcandre/docker-gox
	Image string

	// RemoveContainer controls whether Docker containers are automatically removed upon termination.
	//
	// Example: true
	RemoveContainer bool

	// CommandPaths denotes where to search for Go applications to build.
	// CommandPaths is required, otherwise Port() triggers a validation error.
	//
	// Example: ./cmd/...
	CommandPaths string

	// Verbose enables additional logging when enabled.
	//
	// Example: true
	Verbose bool
}

// NewPortConfig constructs a PortConfig.
//
// Default image: mcandre/docker-gox
// Default remove container: true
// Default verbose: false
func NewPortConfig() PortConfig {
	var config PortConfig
	config.Image = "mcandre/docker-gox"
	config.RemoveContainer = true
	return config
}

// Validate checks a PortConfig for some parameter mistakes.
func (o PortConfig) Validate() error {
	if o.Banner == "" {
		return errors.New("Blank banner")
	}

	if o.CommandPaths == "" {
		return errors.New("Blank command path")
	}

	return nil
}
