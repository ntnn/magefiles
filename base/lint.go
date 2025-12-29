package base

import (
	"strings"

	"github.com/magefile/mage/sh"
)

// Check runs "go vet" with "./...".
func Check() error {
	return Vet("./...")
}

// Vet runs "go vet" with the given argument.
//
// Due to a limitation in mage the argument has to be quoted:
// https://github.com/magefile/mage/issues/340 is resolved
//
// E.g.:
//
//	mage vet "pkg1 pkg2"
func Vet(pkgs string) error {
	return sh.RunV(
		"go",
		append([]string{"vet"}, strings.Split(pkgs, " ")...)...,
	)
}
