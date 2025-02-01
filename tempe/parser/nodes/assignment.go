package nodes

import (
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/opt"
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
	Statement
	vars []Var
}

func NewAssignment() *Assignment {
	return &Assignment{
		Statement: NewStatement(),
		vars:      []Var{},
	}
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

func (n *Assignment) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	if err := n.RenderWhitespace(Upper, opts, w); err != nil {
		return err
	}
	for _, v := range n.vars {
		value, err := env.Eval(v.expression)
		if err != nil {
			return err
		}
		env.Set(v.name, value)
	}
	if err := n.RenderWhitespace(Lower, opts, w); err != nil {
		return err
	}
	return nil
}

func (n Assignment) Format() (string, string) {
	leftUpper, rightUpper := n.whitespace[opt.UpperLeft].text, n.whitespace[opt.UpperRight].text
	leftUpper = replaceWhitespace(leftUpper)
	rightUpper = replaceWhitespace(rightUpper)
	names := strings.Join(n.Names(), " = ; ")
	return "Assignment", fmt.Sprintf(`"%s{%% %s = ; %%}%s"`, leftUpper, names, rightUpper)
}

func (n Assignment) String() string {
	return "Assignment{}"
}
