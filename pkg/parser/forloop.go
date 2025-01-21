package parser

import (
	"fmt"
	"github.com/tniedbala/tempe/pkg/api"
	"io"
	"strings"
)

type ForLoop struct {
	hasIndex   bool
	indexName  string
	varName    string
	expression string
	body       *TemplateNodesCollection
}

func NewForLoop(indexName, varName, expression string, body *TemplateNodesCollection) *ForLoop {
	indexName = strings.TrimSpace(indexName)
	return &ForLoop{
		hasIndex:   indexName != "",
		indexName:  indexName,
		varName:    strings.TrimSpace(varName),
		expression: strings.TrimSpace(expression),
		body:       body,
	}
}

func (n *ForLoop) Children() []api.TemplateNode {
	return n.body.Children()
}

func (n *ForLoop) Render(env api.Env, w io.StringWriter) error {
	value, err := env.Eval(n.expression)
	if err != nil {
		return err
	}
	iter, err := value.Iter()
	if err != nil {
		return err
	}
	localEnv, err := env.LocalEnv()
	if err != nil {
		return err
	}
	for index, value := range iter {
		if n.hasIndex {
			localEnv.Set(n.indexName, index)
		}
		localEnv.Set(n.varName, value)
		if err := n.body.Render(localEnv, w); err != nil {
			return err
		}
	}
	return env.Update(localEnv)
}

func (n ForLoop) String() string {
	var forIn string
	if n.hasIndex {
		forIn = fmt.Sprintf("for %s, %s in", n.indexName, n.varName)
	} else {
		forIn = fmt.Sprintf("for %s in", n.varName)
	}
	return fmt.Sprintf(`ForLoop: "{%% %s ... %%}"`, forIn)
}
