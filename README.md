# goyek demo

[![go.mod](https://img.shields.io/github/go-mod/go-version/goyek/demo)](go.mod)
[![LICENSE](https://img.shields.io/github/license/goyek/demo)](LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/goyek/demo/build)](https://github.com/goyek/demo/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/goyek/demo)](https://goreportcard.com/report/github.com/goyek/demo)
[![Codecov](https://codecov.io/gh/goyek/demo/branch/main/graph/badge.svg)](https://codecov.io/gh/goyek/demo)

⭐ `Star` this repository if you find it valuable and worth maintaining.

## Description

This repository demonstrates how [`goyek`](https://github.com/goyek/goyek)
can be used in big organization to reuse common pattern used in build pipelines.

It uses [`github.com/goyek/workflow`](https://github.com/goyek/workflow)
which is used as a reusable, yet customizable, build framework.

It customizes `workflow` to:

- use `golangci-lint` instead of `go vet`,
- run `misspell` during `test` stage.

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
