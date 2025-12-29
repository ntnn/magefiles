package base

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// PreGenerateDeletePatterns are used in Generate to automatically
// delete generated files before generation using find.
// These can be set when importing the package, e.g.
//
// package main
//
// import (
//
//	//mage:import
//	"github.com/ntnn/magefiles/base"
//
// )
//
//	func init() {
//		base.PreGenerateDeletePatterns = []string{
//			"*_gen.go",
//		}
//	}
var PreGenerateDeletePatterns = []string{
	"*.gen.go",
	"*.gen_test.go",
	"*_gen.go",
	"*_gen_test.go",
}

// preGenerateDelete is automatically run as a dependency of Generate
// and removes files matching patterns from PreGenerateDeletePatterns
// using `find`.
func preGenerateDelete() error {
	for _, pattern := range PreGenerateDeletePatterns {
		if err := sh.RunV("find", ".", "-type", "f", "-name", pattern, "-delete"); err != nil {
			return err
		}
	}
	return nil
}

// PreGenerate allows to add more dependencies to Generate.
var PreGenerate = []any{}

// Generate runs go generate.
func Generate() error {
	mg.Deps(preGenerateDelete)
	mg.Deps(PreGenerate...)
	return sh.RunV("go", "generate", "./...")
}

// Test runs `go test` with coverage and parallelism enabled.
// Parallelism is set to the number of available CPUs.
func Test() error {
	return sh.RunV("go", "test", "-cover", "./...")
}

// All runs generate, check and test in order.
func All() {
	mg.SerialDeps(Generate, Check, Test)
}
