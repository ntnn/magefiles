# magefiles

Common [mage](https://magefile.org/) targets I reuse across repositories.

## Usage

See the upstream documentation for details: https://magefile.org/importing/

```go
package main

import (
	// Import into the global namespace
	//mage:import
	_ "github.com/ntnn/magefiles/base"
	// Import into the "go" namespace
	//mage:import go
	_ "github.com/ntnn/magefiles/base"
)
```

Importing into the "go" namespace results in this:

```
$ mage -l
Targets:
  go:all               runs generate, check and test in order.
  go:check             runs "go vet" with "./...".
  go:generate          runs go generate.
  go:test              runs `go test` with coverage and parallelism enabled.
  go:vet               runs "go vet" with the given argument.
  [...]
```
