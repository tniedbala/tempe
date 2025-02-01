package tempo_cel

import (
	"fmt"
	"maps"
	"slices"

	"github.com/google/cel-go/cel"
	"github.com/tniedbala/tempe-go/tempe/api"
)

type CelEnv struct {
	engine *CelEngine
	values map[string]api.Value
	env    *cel.Env
}

func NewEnv(engine *CelEngine) (*CelEnv, error) {
	env, err := cel.NewEnv()
	if err != nil {
		return nil, err
	}
	e := &CelEnv{
		engine: engine,
		values: map[string]api.Value{},
		env:    env,
	}
	return e, nil
}

func (e *CelEnv) LocalEnv() (api.Env, error) {
	env, err := e.env.Extend()
	if err != nil {
		return nil, err
	}
	localEnv := &CelEnv{
		engine: e.engine,
		values: map[string]api.Value{},
		env:    env,
	}
	return localEnv, nil
}

func (e *CelEnv) Keys() []string {
	return slices.Collect(maps.Keys(e.values))
}

func (e *CelEnv) Map() map[string]api.Value {
	return e.values
}

func (e *CelEnv) Get(name string) (api.Value, bool) {
	value, ok := e.values[name]
	return value, ok
}

func (e *CelEnv) Set(name string, value api.Value) error {
	val, ok := value.(CelValue)
	if !ok {
		// TODO: improve error message
		return fmt.Errorf("error converting %v to CelValue", value)
	}
	t := val.value.Type().(*cel.Type)
	env, err := e.env.Extend(cel.Variable(name, t))
	if err != nil {
		return err
	}
	e.env = env
	e.values[name] = value
	return nil
}

func (e *CelEnv) Update(env api.Env) error {
	for name, value := range env.Map() {
		if err := e.Set(name, value); err != nil {
			return err
		}
	}
	return nil
}

func (e *CelEnv) Eval(src string) (api.Value, error) {
	ast, issues := e.env.Compile(src)
	// TODO: understand if there are cases where issues occur but there is no error
	if issues != nil && issues.Err() != nil {
		return nil, issues.Err()
	}
	prg, err := e.env.Program(ast)
	if err != nil {
		return nil, err
	}

	// TODO: understand usage of details returned from program.Eval()
	val, _, err := prg.Eval(map[string]any{
		"name": "world",
	})
	if err != nil {
		return nil, err
	}
	return CelValue{val}, nil
}

func (e CelEnv) String() string {
	return ""
}
