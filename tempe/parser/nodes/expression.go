package nodes

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type expressionMarshaller struct {
	Expression string `json:"expression"`
}

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

func (n *Expression) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	value, err := env.Eval(n.expression)
	if err != nil {
		return err
	}
	_, err = w.WriteString(value.String())
	return err
}

func (n Expression) Format() (string, string) {
	return "Expression", fmt.Sprintf(`"{{ %s }}"`, formatText(n.expression))
}

func (n Expression) String() string {
	return "Expression{}"
}

func (n *Expression) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonNodeSpec("Expression", expressionMarshaller{
		Expression: n.expression,
	}))
}
