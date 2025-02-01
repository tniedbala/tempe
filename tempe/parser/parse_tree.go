package parser

import (
	"iter"

	"github.com/tniedbala/tempe-go/pkg/api"
)

type ParseTree struct {
	parent *ParseTree
	node   api.TemplateNode
	depth  int
}

func NewParseTree(node api.TemplateNode, parent *ParseTree) *ParseTree {
	depth := 0
	if parent != nil {
		depth = parent.depth + 1
	}
	return &ParseTree{parent, node, depth}
}

func (p *ParseTree) Depth() int {
	return p.depth
}

func (p *ParseTree) Length() int {
	return len(p.node.Children())
}

func (p *ParseTree) Parent() api.ParseTree {
	return p.parent
}

func (p *ParseTree) Node() api.TemplateNode {
	return p.node
}

func (p *ParseTree) Children() iter.Seq2[int, api.ParseTree] {
	return func(yield func(int, api.ParseTree) bool) {
		children := p.node.Children()
		for i, node := range children {
			tree := NewParseTree(node, p)
			if !yield(i, tree) {
				return
			}
		}
	}
}

func (p *ParseTree) PrettyPrint() (string, error) {
	formatter := NewParseTreeFormatter(nil, p, false)
	return formatter.PrettyPrint()
}