# goyek demo

[![go.mod](https://img.shields.io/github/go-mod/go-version/goyek/demo)](go.mod)
[![LICENSE](https://img.shields.io/github/license/goyek/demo)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/goyek/demo/build.yml?branch=main)](https://github.com/goyek/demo/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/goyek/demo)](https://goreportcard.com/report/github.com/goyek/demo)
[![Codecov](https://codecov.io/gh/goyek/demo/branch/main/graph/badge.svg)](https://codecov.io/gh/goyek/demo)

## Description

This repository demonstrates how [`goyek`](https://github.com/goyek/goyek)
can be used in big organizations to reuse common patterns used in build pipelines.

It uses [`github.com/goyek/workflow`](https://github.com/goyek/workflow)
which is used as an example reusable, yet customizable, build framework.

`demo` customizes it so that:

- it uses [`golangci-lint`](https://github.com/golangci/golangci-lint)
  instead of `go vet`,
- it adds [`misspell`](https://github.com/client9/misspell)
  to the `test` stage.

## Build

```sh
cd build
go run ./build
```

Using convenient Bash script:

```sh
./goyek.sh
```

Using convenient PowerShell script:

```pwsh
.\goyek.ps1
```

Example output:

![example output](https://user-images.githubusercontent.com/5067549/199257929-03d7f960-6fd7-48d1-bcd2-704e95a39411.png)
