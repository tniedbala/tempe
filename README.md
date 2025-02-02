# tempe
Template engine framework for use with Go-compatible scripting languages
such as [cel](https://github.com/google/cel-go), [expr](https://github.com/expr-lang/expr)
and [starlark](https://github.com/google/starlark-go).

## Current Status
This library is still in very early stages of development and interfaces are subject to change (see [TODO](./docs/TODO.md)). 

## Docs
- [Template Syntax](./docs/template-syntax.md)
- [Implementation Guide](./docs/implementation-guid.md)
- [Development Environment](./docs/development-environment.md)

## Example Usage
```go
package main 

import (
    "fmt"
    "github.com/tniedbala/tempe-go/tempe"
)

func main() {
	engine := tempe.DefaultEngine()
	template, err := engine.NewTemplate("hello {{ name }}!")
	if err != nil {
		panic(err)
	}

	params := tempe.Params{"name": "world"}
	output, err := template.Render(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
```