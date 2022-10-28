// Build is the build pipeline for this repository.
package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/workflow"
)

func main() {
	goyek.Undefine(workflow.TaskGoVet)
	workflow.StageTest.SetDeps(
		append(goyek.Deps{spell, goLint}, workflow.StageTest.Deps()...), // add as the first task during test stage
	)

	workflow.Main()
}
