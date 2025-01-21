package parser

import (
	"fmt"
	"github.com/tniedbala/tempe/pkg/api"
	"io"
)

type IfStatement struct {
	clauses *TemplateNodesCollection
}

func NewIfStatement(clauses *TemplateNodesCollection) *IfStatement {
	return &IfStatement{clauses}
}

func (n *IfStatement) Children() []api.TemplateNode {
	return n.clauses.Children()
}

func (n *IfStatement) Render(env api.Env, w io.StringWriter) error {
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
			return clause.Render(env, w)
		}
	}
	return nil
}

func (n IfStatement) String() string {
	return "IfStatmement{}"
}
