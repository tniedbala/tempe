package tempe

import (
	"fmt"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
)

// Wrapper struct for api.Env, which will be used to implement intended template variable scope.
type Env struct {
	parent *Env
	env    api.Env
}

// Create a new Env wrapper around env.
func NewEnv(env api.Env) *Env {
	return &Env{env: env}
}

// Create a new local environment (copy of the parent environment).
func (e *Env) LocalEnv() (api.Env, error) {
	localEnv, err := e.env.LocalEnv()
	if err != nil {
		return nil, err
	}
	if err := localEnv.Update(e.env); err != nil {
		return nil, err
	}
	return &Env{e, localEnv}, nil
}

func (e *Env) Keys() []string {
	return e.env.Keys()
}

func (e *Env) Map() map[string]api.Value {
	return e.env.Map()
}

func (e *Env) Get(name string) (api.Value, bool) {
	return e.env.Get(name)
}

func (e *Env) Set(name string, value api.Value) error {
	return e.env.Set(name, value)
}

// only update variables that exist in the current env
func (e *Env) Update(env api.Env) error {
	for _, name := range e.Keys() {
		value, ok := env.Get(name)
		if !ok {
			continue
		}
		if err := e.Set(name, value); err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) Eval(src string) (api.Value, error) {
	return e.env.Eval(src)
}

func (e Env) String() string {
	return fmt.Sprintf("Env{%s}", strings.Join(e.Keys(), ", "))
}
