// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/goxcart"
	"github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

var integrationTestSourceDir = "test"

// integrationTestArtifactsPath describes where integration tests build artifacts.
var integrationTestArtifactsPath = "bin"

// Default references the default build task.
var Default = Test

// UnitTest executes the unit test suite.
func UnitTest() error { return mageextras.UnitTest() }

// IntegrationTest executes the integration test suite.
func IntegrationTest() error {
	mg.Deps(Install)

	cmd := exec.Command(
		"goxcart",
		"-verbose",
		"-output",
		integrationTestArtifactsPath,
		"-banner",
		"hello-0.0.1",
		"-commands",
		"./...",
	)
	cmd.Dir = integrationTestSourceDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Test executes unit and integration tests.
func Test() error { mg.Deps(UnitTest); mg.Deps(IntegrationTest); return nil }

// GoVet runs go vet with shadow checks enabled.
func GoVet() error { return mageextras.GoVetShadow() }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("goxcart-%s", goxcart.Version)

// Port archives build artifacts.
func Port() error { mg.Deps(Gox); return mageextras.Archive(portBasename, artifactsPath) }

// Gox cross-compiles Go binaries.
func Gox() error {
	return mageextras.Gox(
		artifactsPath,
		strings.Join(
			[]string{
				portBasename,
				"{{.OS}}",
				"{{.Arch}}",
				"{{.Dir}}",
			},
			mageextras.PathSeparatorString,
		),
	)
}

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("goxcart") }

// CleanIntegrationTestArtifacts deletes leftover integration test files.
func CleanIntegrationTestArtifacts() error {
	return os.RemoveAll(
		path.Join(integrationTestSourceDir, integrationTestArtifactsPath),
	)
}

// Clean deletes artifacts.
func Clean() error { mg.Deps(CleanIntegrationTestArtifacts); return os.RemoveAll(artifactsPath) }
