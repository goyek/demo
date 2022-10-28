package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/workflow"
)

var goLint = goyek.Define(goyek.Task{
	Name:  "go-lint",
	Usage: "golangci-lint run --fix",
	Action: func(tf *goyek.TF) {
		if !workflow.Exec(tf, "go install github.com/golangci/golangci-lint/cmd/golangci-lint") {
			return
		}
		workflow.Exec(tf, "golangci-lint run --fix")
	},
})
