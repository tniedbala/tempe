package nodes

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	opt "github.com/tniedbala/tempe-go/tempe/options"
)

type Var struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

func NewVar(name, expression string) Var {
	return Var{
		Name:       strings.TrimSpace(name),
		Expression: strings.TrimSpace(expression),
	}
}

type assignmentMarshaller struct {
	Whitespace WhitespaceSyntax `json:"whitespace"`
	Vars       []Var            `json:"vars"`
}

type Assignment struct {
	whitespace WhitespaceSyntax
	vars       []Var
}

func NewAssignment() *Assignment {
	return &Assignment{
		whitespace: NewWhitespaceSyntax(opt.Assignment),
		vars:       []Var{},
	}
}

func (n *Assignment) SetWhitespace(openStmt, closeStmt string) {
	n.whitespace.Leading = NewWhitespace(opt.Leading, openStmt)
	n.whitespace.Trailing = NewWhitespace(opt.Trailing, closeStmt)
}

func (n *Assignment) Children() []api.TemplateNode {
	return []api.TemplateNode{}
}

func (n *Assignment) Names() []string {
	names := make([]string, len(n.vars))
	for i, v := range n.vars {
		names[i] = v.Name
	}
	return names
}

func (n *Assignment) Append(name, expression string) {
	n.vars = append(n.vars, NewVar(name, expression))
}

func (n *Assignment) Render(options api.Options, env api.Env, w io.StringWriter) error {
	for _, v := range n.vars {
		value, err := env.Eval(v.Expression)
		if err != nil {
			return err
		}
		env.Set(v.Name, value)
	}
	return n.whitespace.Render(options, w)
}

func (n Assignment) Format() (string, string) {
	leading := replaceWhitespace(n.whitespace.Leading.Text)
	trailing := replaceWhitespace(n.whitespace.Trailing.Text)
	names := strings.Join(n.Names(), " = ; ")
	return "Assignment", fmt.Sprintf(`"%s{%% %s = ; %%}%s"`, leading, names, trailing)
}

func (n Assignment) String() string {
	return "Assignment{}"
}

func (n *Assignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonNodeSpec("Assignment", assignmentMarshaller{
		Whitespace: n.whitespace,
		Vars:       n.vars,
	}))
}
