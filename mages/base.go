package mages

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func Generate() error {
	return sh.RunV("go", "generate", "./...")
}

func Check() error {
	return sh.RunV("go", "run", "github.com/golangci/golangci-lint/cmd/golangci-lint", "run", "./...")
}

func Lint(subcmd string) error {
	// TODO until https://github.com/magefile/mage/issues/340 is resolved
	return sh.RunV(
		"go",
		append(
			[]string{"run", "github.com/golangci/golangci-lint/cmd/golangci-lint"},
			strings.Split(subcmd, " ")...,
		)...,
	)
}

func Test() error {
	return sh.RunV("go", "test", "-parallel", fmt.Sprintf("%d", runtime.NumCPU()), "-cover", "./...")
}

func All() {
	mg.SerialDeps(Generate, Check, Test)
}
