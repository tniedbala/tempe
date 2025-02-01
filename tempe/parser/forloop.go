package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/opt"
)

type ForLoop struct {
	Statement
	hasIndex   bool
	indexName  string
	varName    string
	expression string
	body       *TemplateNodesCollection
}

func NewForLoop(indexName, varName, expression string, body *TemplateNodesCollection) *ForLoop {
	indexName = strings.TrimSpace(indexName)
	return &ForLoop{
		Statement:  NewStatement(),
		hasIndex:   indexName != "",
		indexName:  indexName,
		varName:    strings.TrimSpace(varName),
		expression: strings.TrimSpace(expression),
		body:       body,
	}
}

func (n *ForLoop) InnerWhitespace() string {
	text := ""
	if ws, ok := n.whitespace[opt.LowerLeft]; ok {
		if ws.flag != Strip {
			text = ws.text
		}
	}
	return text
}

func (n *ForLoop) Children() []api.TemplateNode {
	return n.body.Children()
}

func (n *ForLoop) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	if err := n.RenderWhitespace(Upper, opts, w); err != nil {
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
		if err := n.body.Render(opts, localEnv, w); err != nil {
			return err
		}
		if _, err = w.WriteString(innerWhitespace); err != nil {
			return err
		}
	}
	if err := n.RenderWhitespace(Lower, opts, w); err != nil {
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
	leftUpper, rightUpper := n.whitespace[opt.UpperLeft].text, n.whitespace[opt.UpperRight].text
	leftUpper = replaceWhitespace(leftUpper)
	rightUpper = replaceWhitespace(rightUpper)
	return "ForLoop", fmt.Sprintf(`"%s{%% %s ... %%}%s"`, leftUpper, forIn, rightUpper)
}

func (n ForLoop) String() string {
	return "ForLoop{}"
}
