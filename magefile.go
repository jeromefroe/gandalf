//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Compiles the project binary.
func Build() error {
	return sh.Run("go", "build", "./...")
}

// Runs unit tests.
func Test() error {
	return sh.RunV("go", "test", "-v", "./...")
}

// Removes old build artifacts.
func Clean() error {
	return sh.Run("go", "clean")
}
