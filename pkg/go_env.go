package tempe

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
	"github.com/tniedbala/tempe/pkg/api"
)

var regexpVarName *regexp.Regexp

func init() {
	regexpVarName = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
}

type GoEnv struct {
	values map[string]GoValue
}

func NewGoEnv(values map[string]any) *GoEnv {
	m := map[string]GoValue{}
	for key, value:= range values {
		m[key] = NewGoValue(value)
	}
	return &GoEnv{m}
}

func (e *GoEnv) LocalEnv() (api.Env, error) {
	localEnv := NewGoEnv(nil)
	return localEnv, nil
}

func (e *GoEnv) Keys() []string {
	return slices.Collect(maps.Keys(e.values))
}

func (e *GoEnv) Map() map[string]api.Value {
	m := map[string]api.Value{}
	for key, value := range e.values {
		m[key] = value
	}
	return m
}

func (e *GoEnv) Get(name string) (api.Value, bool) {
	v, ok := e.values[name]
	return v, ok
}

func (e *GoEnv) Set(name string, value api.Value) error {
	v, ok := value.(GoValue)
	if !ok {
		// TODO: improve error message
		return fmt.Errorf("error converting %v to GoValue", v)
	}
	e.values[name] = v
	return nil
}

func (e *GoEnv) Update(env api.Env) error {
	for name, value := range env.Map() {
		if err := e.Set(name, value); err != nil {
			return err
		}
	}
	return nil
}

func (e *GoEnv) Eval(src string) (api.Value, error) {
	if !regexpVarName.Match([]byte(src)) {
		return nil, fmt.Errorf(`syntax error: GoEngine can only evaluate variable names; got "%s"`, src)
	}
	value, ok := e.values[src]
	if !ok {
		return nil, fmt.Errorf(`variable not found: "%s"`, src)
	}
	return value, nil
}

func (e GoEnv) String() string {
	return fmt.Sprintf("GoEnv{%s}", strings.Join(e.Keys(), ", "))
}
