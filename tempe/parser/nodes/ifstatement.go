package nodes

import (
	"encoding/json"
	"fmt"
	"io"
	// "strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	opt "github.com/tniedbala/tempe-go/tempe/options"
)

type ifStatementMarshaller struct {
	Whitespace WhitespaceSyntax         `json:"whitespace"`
	Clauses    *TemplateNodesCollection `json:"clauses"`
}

type IfStatement struct {
	whitespace WhitespaceSyntax
	clauses    *TemplateNodesCollection
}

func NewIfStatement(clauses *TemplateNodesCollection) *IfStatement {
	return &IfStatement{
		whitespace: NewWhitespaceSyntax(opt.EndIf),
		clauses:    clauses,
	}
}

func (n *IfStatement) SetWhitespace(openStmt, closeStmt string) {
	n.whitespace.Leading = NewWhitespace(opt.Leading, openStmt)
	n.whitespace.Trailing = NewWhitespace(opt.Trailing, closeStmt)
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
			// b := w.(*strings.Builder)
			// fmt.Println(b.String())
			break
		}
	}
	return n.whitespace.Render(opts, w)
}

func (n *IfStatement) Format() (string, string) {
	return "IfStatement", ""
}

func (n IfStatement) String() string {
	return "IfStatement{}"
}

func (n *IfStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonNodeSpec("IfStatement", ifStatementMarshaller{
		Whitespace: n.whitespace,
		Clauses:    n.clauses,
	}))
}
