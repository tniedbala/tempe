package nodes

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	opt "github.com/tniedbala/tempe-go/tempe/options"
)

type forLoopWhitespace struct {
	For    WhitespaceSyntax `json:"for"`
	EndFor WhitespaceSyntax `json:"endfor"`
}

type forLoopMarshaller struct {
	HasIndex   bool                     `json:"has_index"`
	IndexName  string                   `json:"index_name"`
	VarName    string                   `json:"var_name"`
	Expression string                   `json:"expression"`
	Whitespace forLoopWhitespace        `json:"whitespace"`
	Body       *TemplateNodesCollection `json:"body"`
}

type ForLoop struct {
	hasIndex   bool
	indexName  string
	varName    string
	expression string
	body       *TemplateNodesCollection
	whitespace forLoopWhitespace
}

func NewForLoop(indexName, varName, expression string, body *TemplateNodesCollection) *ForLoop {
	indexName = strings.TrimSpace(indexName)
	return &ForLoop{
		hasIndex:   indexName != "",
		indexName:  indexName,
		varName:    strings.TrimSpace(varName),
		expression: strings.TrimSpace(expression),
		body:       body,
		whitespace: forLoopWhitespace{
			For:    NewWhitespaceSyntax(opt.For),
			EndFor: NewWhitespaceSyntax(opt.EndFor),
		},
	}
}

func (n *ForLoop) SetWhitespace(elem opt.SyntaxElement, openStmt, closeStmt string) {
	switch elem {
	case opt.For:
		n.whitespace.For.Leading = NewWhitespace(opt.Leading, openStmt)
		n.whitespace.For.Trailing = NewWhitespace(opt.Trailing, closeStmt)
	case opt.EndFor:
		n.whitespace.EndFor.Leading = NewWhitespace(opt.Leading, openStmt)
		n.whitespace.EndFor.Trailing = NewWhitespace(opt.Trailing, closeStmt)
	}
}

func (n *ForLoop) InnerWhitespace() string {
	text := ""
	ws := n.whitespace.EndFor.Leading
	if ws.Flag != Trim {
		text = ws.Text
	}
	return text
}

func (n *ForLoop) Children() []api.TemplateNode {
	return n.body.Children()
}

func (n *ForLoop) Render(options api.Options, env api.Env, w io.StringWriter) error {
	if err := n.whitespace.For.Render(options, w); err != nil {
		return err
	}
	value, err := env.Eval(n.expression)
	if err != nil {
		return err
	}
	iter, err := value.Iter()
	if err != nil {
		return err
	}
	localEnv, err := env.Copy()
	if err != nil {
		return err
	}
	innerWhitespace := n.InnerWhitespace()
	for index, value := range iter {
		if n.hasIndex {
			localEnv.Set(n.indexName, index)
		}
		localEnv.Set(n.varName, value)
		if err := n.body.Render(options, localEnv, w); err != nil {
			return err
		}
		if _, err = w.WriteString(innerWhitespace); err != nil {
			return err
		}
	}
	if err := n.whitespace.EndFor.Render(options, w); err != nil {
		return err
	}
	return env.Update(localEnv)
}

func (n ForLoop) Format() (string, string) {
	var forIn string
	if n.hasIndex {
		forIn = fmt.Sprintf("for %s, %s in", n.indexName, n.varName)
	} else {
		forIn = fmt.Sprintf("for %s in", n.varName)
	}
	leftUpper, rightUpper := n.whitespace.For.Leading.Text, n.whitespace.For.Trailing.Text
	leftUpper = replaceWhitespace(leftUpper)
	rightUpper = replaceWhitespace(rightUpper)
	return "ForLoop", fmt.Sprintf(`"%s{%% %s ... %%}%s"`, leftUpper, forIn, rightUpper)
}

func (n ForLoop) String() string {
	return "ForLoop{}"
}

func (n *ForLoop) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonNodeSpec("ForLoop", forLoopMarshaller{
		HasIndex:   n.hasIndex,
		IndexName:  n.indexName,
		VarName:    n.varName,
		Expression: n.expression,
		Whitespace: n.whitespace,
		Body:       n.body,
	}))
}
