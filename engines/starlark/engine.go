package tempo_starlark

import (
	"fmt"
	"io"

	"github.com/tniedbala/tempe-go/tempe"
	"github.com/tniedbala/tempe-go/tempe/api"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

type StarlarkEngine struct {
	opts   *syntax.FileOptions
	thread *starlark.Thread
}

func NewEngine() *StarlarkEngine {
	opts := syntax.LegacyFileOptions()
	thread := &starlark.Thread{
		Name: "Starlark",
		Print: func(_ *starlark.Thread, msg string) {
			fmt.Println(msg)
		},
	}
	return &StarlarkEngine{opts, thread}
}

func (e *StarlarkEngine) Read(reader io.Reader) (api.Template, error) {
	return tempe.Read(e, reader)
}

func (e *StarlarkEngine) ReadFile(path string) (api.Template, error) {
	return tempe.ReadFile(e, path)
}

func (e *StarlarkEngine) NewTemplate(src string) (api.Template, error) {
	return tempe.NewTemplate(e, src)
}

func (e *StarlarkEngine) DefaultEnv() api.Env {
	return NewEnv(e, starlark.StringDict{})
}

func (e *StarlarkEngine) NewEnv(env map[string]any) (api.Env, error) {
	stringDict := starlark.StringDict{}
	for k, v := range env {
		value, ok := v.(starlark.Value)
		if !ok {
			return nil, fmt.Errorf("error converting %v to starlark.value", v)
		}
		stringDict[k] = value
	}
	return NewEnv(e, stringDict), nil
}

func (e *StarlarkEngine) NewValue(value any) (api.Value, error) {
	v, ok := value.(starlark.Value)
	if !ok {
		return nil, fmt.Errorf("error converting %v to starlark.value", v)
	}
	return StarlarkValue{v}, nil
}

func (e *StarlarkEngine) String() string {
	return "StarlarkEngine{}"
}
