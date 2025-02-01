package tempo_starlark

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	"go.starlark.net/starlark"
)

type StarlarkEnv struct {
	engine     *StarlarkEngine
	stringDict starlark.StringDict
}

func NewEnv(engine *StarlarkEngine, stringDict starlark.StringDict) *StarlarkEnv {
	return &StarlarkEnv{engine, stringDict}
}

func (e *StarlarkEnv) LocalEnv() (api.Env, error) {
	localEnv := &StarlarkEnv{
		engine:     e.engine,
		stringDict: starlark.StringDict{},
	}
	return localEnv, nil
}

func (e *StarlarkEnv) Keys() []string {
	return slices.Collect(maps.Keys(e.stringDict))
}

func (e *StarlarkEnv) Map() map[string]api.Value {
	m := map[string]api.Value{}
	for k, v := range e.stringDict {
		m[k] = StarlarkValue{v}
	}
	return m
}

func (e *StarlarkEnv) Get(name string) (api.Value, bool) {
	v, ok := e.stringDict[name]
	if !ok {
		return nil, false
	}
	return StarlarkValue{v}, true
}

func (e *StarlarkEnv) Set(name string, value api.Value) error {
	v, ok := value.(StarlarkValue)
	if !ok {
		// TODO: improve error message
		return fmt.Errorf("error converting %v to starlark.value", v)
	}
	e.stringDict[name] = v.value
	return nil
}

func (e *StarlarkEnv) Update(env api.Env) error {
	for name, value := range env.Map() {
		if err := e.Set(name, value); err != nil {
			return err
		}
	}
	return nil
}

func (e *StarlarkEnv) Eval(src string) (api.Value, error) {
	value, err := starlark.EvalOptions(e.engine.opts, e.engine.thread, "starlark", src, e.stringDict)
	return StarlarkValue{value}, err
}

func (e StarlarkEnv) String() string {
	return fmt.Sprintf("StarlarkEnv{%s}", strings.Join(e.stringDict.Keys(), ", "))
}
