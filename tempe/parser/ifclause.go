package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/pkg/api"
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
		return "IfClause"
	}
}

type IfClause struct {
	kind       IfClauseKind
	expression string
	body       *TemplateNodesCollection
}

func NewIfClause(kind IfClauseKind, expression string, body *TemplateNodesCollection) *IfClause {
	return &IfClause{
		kind:       kind,
		expression: strings.TrimSpace(expression),
		body:       body,
	}
}

func (n *IfClause) Children() []api.TemplateNode {
	return []api.TemplateNode{n.body}
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

func (n *IfClause) Render(env api.Env, w io.StringWriter) error {
	localEnv, err := env.LocalEnv()
	if err != nil {
		return err
	}
	if err = n.body.Render(localEnv, w); err != nil {
		return nil
	}
	return env.Update(localEnv)
}

func (n *IfClause) String() string {
	if n.kind == Else {
		return fmt.Sprintf("IfClause{%s}", n.kind)
	}
	return fmt.Sprintf(`IfClause{%s: %s}`, n.kind, formatText(n.expression))
}
