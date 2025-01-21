package parser

import (
	"fmt"
	"io"
	"strings"
	"github.com/tniedbala/tempe/pkg/api"
)

type Var struct {
	name       string
	expression string
}

func NewVar(name, expression string) Var {
	return Var{
		name:       strings.TrimSpace(name),
		expression: strings.TrimSpace(expression),
	}
}

type Assignment struct {
	vars []Var
}

func NewAssignment() *Assignment {
	return &Assignment{vars: []Var{}}
}

func (n *Assignment) Children() []api.TemplateNode {
	return []api.TemplateNode{}
}

func (n *Assignment) Names() []string {
	names := make([]string, len(n.vars))
	for i, v := range n.vars {
		names[i] = v.name
	}
	return names
}

func (n *Assignment) Append(name, expression string) {
	n.vars = append(n.vars, NewVar(name, expression))
}

func (n *Assignment) Render(env api.Env, w io.StringWriter) error {
	for _, v := range n.vars {
		value, err := env.Eval(v.expression)
		if err != nil {
			return err
		}
		env.Set(v.name, value)
	}
	return nil
}

func (n Assignment) String() string {
	names := strings.Join(n.Names(), "; ")
	return fmt.Sprintf(`Assignment: "{%% %s; %%}"`, names)
}
