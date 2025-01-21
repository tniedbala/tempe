package parser

import (
	"fmt"
	"io"
	"strings"
	"github.com/tniedbala/tempe/pkg/api"
)

type Expression struct {
	expression string
}

func NewExpression(expression string) *Expression {
	return &Expression{
		expression: strings.TrimSpace(expression),
	}
}

func (n *Expression) Children() []api.TemplateNode {
	return []api.TemplateNode{}
}

func (n *Expression) Render(env api.Env, w io.StringWriter) error {
	value, err := env.Eval(n.expression)
	if err != nil {
		return err
	}
	_, err = w.WriteString(value.String())
	return err
}

func (n Expression) String() string {
	return fmt.Sprintf(`Expression: "%s"`, formatText(n.expression))
}