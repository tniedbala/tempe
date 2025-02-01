package nodes

import (
	"fmt"
	"io"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type IfStatement struct {
	Statement
	clauses *TemplateNodesCollection
}

func NewIfStatement(clauses *TemplateNodesCollection) *IfStatement {
	return &IfStatement{
		Statement: NewStatement(),
		clauses:   clauses,
	}
}

func (n *IfStatement) Children() []api.TemplateNode {
	return n.clauses.Children()
}

func (n *IfStatement) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	for node := range n.clauses.Iter() {
		clause, ok := node.(*IfClause)
		if !ok {
			// TODO: improve error message
			return fmt.Errorf("invalid node: %v", node)
		}
		ok, err := clause.Eval(env)
		if err != nil {
			return err
		}
		if ok {
			if err = clause.Render(opts, env, w); err != nil {
				return err
			}
		}
	}
	if err := n.RenderWhitespace(Lower, opts, w); err != nil {
		return err
	}
	return nil
}

func (n *IfStatement) Format() (string, string) {
	return "IfStatement", ""
}

func (n IfStatement) String() string {
	return "IfStatement{}"
}