package nodes

import (
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/opt"
)

type IfClause struct {
	Statement
	kind       IfClauseKind
	expression string
	body       *TemplateNodesCollection
}

func NewIfClause(kind IfClauseKind, expression string, body *TemplateNodesCollection) *IfClause {
	return &IfClause{
		Statement:  NewStatement(),
		kind:       kind,
		expression: strings.TrimSpace(expression),
		body:       body,
	}
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

func (n *IfClause) renderWhitespace(opts api.Options, w io.StringWriter) error {
	if n.kind == If {
		return n.RenderWhitespace(Upper, opts, w)
	}
	return n.RenderWhitespace(Inner, opts, w)
}

func (n *IfClause) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	if err := n.renderWhitespace(opts, w); err != nil {
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
	var leftWs, rightWs string
	if n.kind == If {
		leftWs, rightWs = n.whitespace[opt.UpperLeft].text, n.whitespace[opt.UpperRight].text
	} else {
		leftWs, rightWs = n.whitespace[opt.InnerLeft].text, n.whitespace[opt.InnerRight].text
	}
	leftWs = replaceWhitespace(leftWs)
	rightWs = replaceWhitespace(rightWs)
	if n.kind == Else {
		return label, fmt.Sprintf("%s{%% %s %%}%s", leftWs, n.kind, rightWs)
	}
	return label, fmt.Sprintf(`"%s{%% %s %s %%}%s"`, leftWs, n.kind, formatText(n.expression), rightWs)
}

func (n *IfClause) String() string {
	return fmt.Sprintf("IfClause{kind: %s}", n.kind)
}
