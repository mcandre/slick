//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mcandre/mx"
)

// ArtifactsPath describes where artifacts are produced.
const ArtifactsPath = "bin"

// Default references the default build task.
var Default = Build

// Audit runs a security audit.
func Audit() error { return Govulncheck() }

// Build compiles Go projects.
func Build() error {
        dest := ArtifactsPath

        if d, ok := os.LookupEnv("DEST"); ok && d != "" {
                dest = d
        }

        if err := os.MkdirAll(dest, 0755); err != nil {
                return err
        }

        return sh.RunV("go", "build", "-o", dest, "./...")
}

// Clean deletes artifacts.
func Clean() error { return CleanBuild() }

// CleanBuild removes build artifacts.
func CleanBuild() error { return os.RemoveAll(ArtifactsPath) }

// Deadcode runs deadcode.
func Deadcode() error { return sh.RunV("deadcode", "./...") }

// Errcheck runs errcheck.
func Errcheck() error { return sh.RunV("errcheck", "-blank") }

// GoImports runs goimports.
func GoImports() error { return mx.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mx.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return sh.RunV("govulncheck", "-scan", "package", "./...") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// Nakedret runs nakedret.
func Nakedret() error { return mx.Nakedret("-l", "0") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mx.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return sh.RunV("staticcheck", "./...") }

// Test runs a test suite.
func Test() error { return mx.UnitTest() }

// Install builds and installs Go applications.
func Install() error { return mx.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mx.Uninstall("slick") }
