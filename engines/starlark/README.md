# tempe_starlark
Tempe implementation using [starlark-go](https://github.com/google/starlark-go).

## Example Usage
```go
import (
    "fmt"
    "github.com/tniedbala/tempe/engines/starlark"
)

engine := tempe_starlark.NewEngine()
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