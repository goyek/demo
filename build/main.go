// Build is the build pipeline for this repository.
package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/workflow"
)

func main() {
	goyek.Undefine(workflow.TaskGoVet) // remove go vet (we prefer golangci-lint which is a superset)
	workflow.StageTest.SetDeps(
		append(goyek.Deps{spell, goLint}, workflow.StageTest.Deps()...), // add new tasks at the beginning of the test stage
	)

	workflow.Main()
}
