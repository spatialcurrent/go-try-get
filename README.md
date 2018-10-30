[![CircleCI](https://circleci.com/gh/spatialcurrent/go-try-get/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-try-get/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-try-get)](https://goreportcard.com/report/spatialcurrent/go-try-get)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-try-get?status.svg)](https://godoc.org/github.com/spatialcurrent/go-try-get) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-try-get/blob/master/LICENSE.md)

# go-try-get

# Description

**go-try-get** aka (GTG) is a package that wraps around the [reflect](https://godoc.org/reflect) package to provide a standard abstraction layer for getting values by name from objects (structs, maps, and single-value "property" functions).  GTG is used by:

- [go-dfl](https://github.com/spatialcurrent/go-dfl) and
- [railgun](https://github.com/spatialcurrent/railgun).

For example, with GTG, you can provide a single pathways for an API to process structs and maps, allowing automated structured input from systems, as well as user-based input unmarshaled from JSON.

# Usage

**Go**

You can import **go-try-get** as a library with:

```
import (
  "github.com/spatialcurrent/go-try-get/gtg"
)
...
object := map[string]string{"yo":"yo"}
...
if gtg.TryGetString(object, "yo", "what") == "yo" {
  ...
}
...
```

See [gtg](https://godoc.org/github.com/spatialcurrent/go-try-get/gtg) in GoDoc for information on how to use Go API.

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-try-get/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
