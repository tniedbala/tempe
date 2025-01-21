# tempe
Template engine framework for use with Go-compatible scripting languages
such as [cel](https://github.com/google/cel-go), [expr](https://github.com/expr-lang/expr)
and [starlark](https://github.com/google/starlark-go).

## Current Status
This is very much a work in progress - there is quite a bit more development currently needed
to get this into a useable state.

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
