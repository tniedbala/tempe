# tempe_starlark
Tempe implementation using [starlark-go](https://github.com/google/starlark-go).

## Example Usage
```go
import (
    "fmt"
	"go.starlark.net/starlark"
	"github.com/tniedbala/tempe-go/tempe"
    "github.com/tniedbala/tempe-go/engines/starlark"
)

engine := tempe_starlark.DefaultEngine()
template, err := engine.NewTemplate("hello {{ name }}!")
if err != nil {
    panic(err)
}

params := tempe.Params{"name": starlark.String("world")}
output, err := template.Render(params)
if err != nil {
    panic(err)
}
fmt.Println(output)
```