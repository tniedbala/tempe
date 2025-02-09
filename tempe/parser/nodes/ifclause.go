package nodes

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	opt "github.com/tniedbala/tempe-go/tempe/options"
)

type IfClauseKind int

const (
	If IfClauseKind = iota
	ElseIf
	Else
)

func (k IfClauseKind) String() string {
	switch k {
	case If:
		return "if"
	case ElseIf:
		return "elseif"
	case Else:
		return "else"
	default:
		return "undefined"
	}
}

var ifClauseMap = map[IfClauseKind]opt.SyntaxElement{
	If:     opt.If,
	ElseIf: opt.ElseIf,
	Else:   opt.Else,
}

type ifClauseMarshaller struct {
	Clause     string                   `json:"clause"`
	Expression string                   `json:"expression"`
	Whitespace WhitespaceSyntax         `json:"whitespace"`
	Body       *TemplateNodesCollection `json:"body"`
}

type IfClause struct {
	whitespace WhitespaceSyntax
	kind       IfClauseKind
	expression string
	body       *TemplateNodesCollection
}

func NewIfClause(kind IfClauseKind, expression string, body *TemplateNodesCollection) *IfClause {
	elem := ifClauseMap[kind]
	return &IfClause{
		whitespace: NewWhitespaceSyntax(elem),
		kind:       kind,
		expression: strings.TrimSpace(expression),
		body:       body,
	}
}

func (n *IfClause) SetWhitespace(openStmt, closeStmt string) {
	n.whitespace.Leading = NewWhitespace(opt.Leading, openStmt)
	n.whitespace.Trailing = NewWhitespace(opt.Trailing, closeStmt)
}

func (n *IfClause) Children() []api.TemplateNode {
	return n.body.Children()
}

func (n *IfClause) Eval(env api.Env) (bool, error) {
	if n.kind == Else {
		return true, nil
	}
	value, err := env.Eval(n.expression)
	if err != nil {
		return false, err
	}
	return value.Bool(), nil
}

func (n *IfClause) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	if err := n.whitespace.Render(opts, w); err != nil {
		return err
	}
	localEnv, err := env.Copy()
	if err != nil {
		return err
	}
	if err = n.body.Render(opts, localEnv, w); err != nil {
		return nil
	}
	return env.Update(localEnv)
}

func (n *IfClause) Format() (string, string) {
	label := "IfClause"
	leading := replaceWhitespace(n.whitespace.Leading.Text)
	trailing := replaceWhitespace(n.whitespace.Trailing.Text)
	if n.kind == Else {
		return label, fmt.Sprintf("%s{%% %s %%}%s", leading, n.kind, trailing)
	}
	return label, fmt.Sprintf(`"%s{%% %s %s %%}%s"`, leading, n.kind, formatText(n.expression), trailing)
}

func (n *IfClause) String() string {
	return fmt.Sprintf("IfClause{kind: %s}", n.kind)
}

func (n *IfClause) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonNodeSpec("IfClause", ifClauseMarshaller{
		Clause:     n.kind.String(),
		Expression: n.expression,
		Whitespace: n.whitespace,
		Body:       n.body,
	}))
}
