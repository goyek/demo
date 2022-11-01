package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var goLint = goyek.Define(goyek.Task{
	Name:  "go-lint",
	Usage: "golangci-lint run --fix",
	Action: func(tf *goyek.TF) {
		if !cmd.Exec(tf, "go install github.com/golangci/golangci-lint/cmd/golangci-lint") {
			return
		}
		cmd.Exec(tf, "golangci-lint run --fix")
	},
})
