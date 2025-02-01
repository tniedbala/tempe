package tempe_starlark

import (
	"fmt"

	tempe "github.com/tniedbala/tempe-go/tempe"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

type StarlarkEngine struct {
	tempe.TemplateEngine
}

func NewEngine(opts *syntax.FileOptions, thread *starlark.Thread) StarlarkEngine {
	env, value := NewEnv(opts, thread), StarlarkValue{}
	return StarlarkEngine{
		tempe.NewTemplateEngine(env, value),
	}
}

func DefaultEngine() StarlarkEngine {
	opts := syntax.LegacyFileOptions()
	thread := &starlark.Thread{
		Name: "Starlark",
		Print: func(_ *starlark.Thread, msg string) {
			fmt.Println(msg)
		},
	}
	return NewEngine(opts, thread)
}

func (e StarlarkEngine) String() string {
	return "StarlarkEngine{}"
}
