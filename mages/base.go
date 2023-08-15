package mages

import (
	"fmt"
	"runtime"
	"strings"

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
//	"github.com/ntnn/magefiles/mages"
//
// )
//
//	func init() {
//		mages.PreGenerateDeletePatterns = []string{
//			"_gen.go",
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

func Generate() error {
	mg.Deps(append(PreGenerate, preGenerateDelete)...)
	return sh.RunV("go", "generate", "./...")
}

// Check runs golangci-lint.
func Check() error {
	return sh.RunV("go", "run", "github.com/golangci/golangci-lint/cmd/golangci-lint", "run", "./...")
}

// Lint runs golangci-lint with the given argument.
// Due to a limitation in mage the argument has to be quoted:
// https://github.com/magefile/mage/issues/340 is resolved
//
// E.g.:
//
//	go run ./mage.go default:lint 'completion zsh'
func Lint(subcmd string) error {
	return sh.RunV(
		"go",
		append(
			[]string{"run", "github.com/golangci/golangci-lint/cmd/golangci-lint"},
			strings.Split(subcmd, " ")...,
		)...,
	)
}

// Test runs `go test` with coverage and parallelism enabled.
// Parallelism is set to the number of available CPUs.
func Test() error {
	return sh.RunV("go", "test", "-parallel", fmt.Sprintf("%d", runtime.NumCPU()), "-cover", "./...")
}

// All runs generate, check and test in order.
func All() {
	mg.SerialDeps(Generate, Check, Test)
}
