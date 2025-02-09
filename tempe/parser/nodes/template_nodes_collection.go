package nodes

import (
	"encoding/json"
	"io"
	"iter"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type TemplateNodesCollection struct {
	nodes []api.TemplateNode
}

func NewTemplateNodesCollection(nodes ...api.TemplateNode) *TemplateNodesCollection {
	return &TemplateNodesCollection{
		nodes: []api.TemplateNode{},
	}
}

func (n *TemplateNodesCollection) Children() []api.TemplateNode {
	return n.nodes
}

func (n *TemplateNodesCollection) Iter() iter.Seq[api.TemplateNode] {
	return func(yield func(api.TemplateNode) bool) {
		for _, node := range n.nodes {
			if !yield(node) {
				return
			}
		}
	}
}

func (n *TemplateNodesCollection) Append(node api.TemplateNode) {
	n.nodes = append(n.nodes, node)
}

func (n *TemplateNodesCollection) Render(opts api.Options, env api.Env, w io.StringWriter) error {
	for node := range n.Iter() {
		if err := node.Render(opts, env, w); err != nil {
			return err
		}
	}
	return nil
}

func (n *TemplateNodesCollection) Format() (string, string) {
	return "TemplateNodesCollection", ""
}

func (n *TemplateNodesCollection) String() string {
	return "TemplateNodesCollection{}"
}

func (n *TemplateNodesCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.nodes)
}