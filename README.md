# tempe
Template engine framework for use with Go-compatible scripting languages
such as [cel](https://github.com/google/cel-go), [expr](https://github.com/expr-lang/expr)
and [starlark](https://github.com/google/starlark-go).

## Current Status
Very much a work in progress - see [TODO](#todo).

## Docs
- [Template Syntax](./docs/template-syntax.md)
- [Implementation Guide](./docs/implementation-guid.md)
- [Development Environment](./docs/development-environment.md)

## Example Usage
```go
import (
    "fmt"
    "github.com/tniedbala/tempe/pkg"
)

engine := tempe.DefaultEngine()
template, err := engine.NewTemplate("hello {{ name }}!")
if err != nil {
    panic(err)
}

params := map[string]any{"name": "world"}
output, err := template.Render(params)
if err != nil {
    panic(err)
}
fmt.Println(output)
```

## TODO
Core library:
- [ ] Add logic for whitespace control.
- [ ] Review logic around local environment creation & variable scope.
- [ ] Improve error handling.
- [ ] Create test suite.

Implementations:
- [ ] Complete default implementation.
- [ ] Complete [Cel implementation](./engines/cel).
- [ ] Complete [Expr implementation](./engines/expr).
- [ ] Complete [Starlark implementation](./engines/starlark).

Documentation:
- [ ] Add docstrings to go source.
- [ ] Complete [implementation guide](./docs/implementation-guid.md).
- [ ] Improve [template syntax](./docs/template-syntax.md) docs.