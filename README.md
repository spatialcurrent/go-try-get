[![CircleCI](https://circleci.com/gh/spatialcurrent/go-try-get/tree/main.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-try-get/tree/main)
[![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-try-get?style=flat-square)](https://goreportcard.com/report/github.com/spatialcurrent/go-try-get)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/spatialcurrent/go-try-get)](https://pkg.go.dev/github.com/spatialcurrent/go-try-get)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-try-get/blob/main/LICENSE)

# go-try-get

# Description

**go-try-get** aka (GTG) is a package that wraps around the [reflect](https://pkg.go.dev/reflect) package to provide a standard abstraction layer for getting values by name from objects (structs, maps, and single-value "property" functions).  GTG is used by:

- [go-dfl](https://github.com/spatialcurrent/go-dfl) and
- [railgun](https://github.com/spatialcurrent/railgun).

For example, with GTG, you can provide a single pathways for an API to process structs and maps, allowing automated structured input from systems, as well as user-based input unmarshaled from JSON.

# Usage

**Go**

You can import **go-try-get** as a library with:

```
import (
  "github.com/spatialcurrent/go-try-get/pkg/gtg"
)
```

See [gtg](https://pkg.go.dev/github.com/spatialcurrent/go-try-get/pkg/gtg) in the docs for information on how to use Go API.

# Testing

To run Go tests using `make test` or (`bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-try-get/blob/main/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
