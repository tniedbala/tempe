package parser

import (
	"bytes"
	"io"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type ParseTreeFormatter struct {
	tree      api.ParseTree
	parent    *ParseTreeFormatter
	node      api.TemplateNode
	depth     int
	lastChild bool
	indent    string
}

func NewParseTreeFormatter(parent *ParseTreeFormatter, tree api.ParseTree, lastChild bool) *ParseTreeFormatter {
	formatter := &ParseTreeFormatter{
		parent:    parent,
		tree:      tree,
		node:      tree.Node(),
		depth:     tree.Depth(),
		lastChild: lastChild,
	}
	return formatter
}

func (p *ParseTreeFormatter) Branch() string {
	if p.lastChild {
		return "└── "
	} else {
		return "├── "
	}
}

func (p *ParseTreeFormatter) ParentIndent() string {
	if p.depth <= 1 {
		return ""
	}
	if p.parent == nil {
		return ""
	}
	return p.parent.indent
}

func (p *ParseTreeFormatter) Indent() string {
	if p.depth <= 1 {
		return ""
	}
	if p.parent != nil && p.parent.lastChild {
		return "    "
	}
	return "│   "
}

func (p *ParseTreeFormatter) Render(w io.StringWriter) error {
	var newline, branch, str string
	if p.depth == 0 {
		newline, branch, str = "", "", "Template:"
	} else {
		newline, branch, str = "\n", p.Branch(), p.node.String()
	}
	p.indent = p.ParentIndent() + p.Indent()
	if _, err := w.WriteString(newline + p.indent + branch + str); err != nil {
		return err
	}
	lastIndex := p.tree.Length() - 1
	for i, tree := range p.tree.Children() {
		formatter := NewParseTreeFormatter(p, tree, i==lastIndex)
		if err := formatter.Render(w); err != nil {
			return err
		}
	}
	return nil
}

func (p *ParseTreeFormatter) PrettyPrint() (string, error) {
	var b bytes.Buffer
	if err := p.Render(&b); err != nil {
		return "", err
	}
	return b.String(), nil
}
