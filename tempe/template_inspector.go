package tempe

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type templateMarshaller struct {
	Nodes api.TemplateNode `json:"nodes"`
}

type TemplateInspector struct {
	template *Template
}

func NewTemplateInspector(template *Template) TemplateInspector {
	return TemplateInspector{template}
}

func (t TemplateInspector) ParseTree() api.ParseTree {
	return t.template.tree
}

func (t TemplateInspector) PrettyPrint() {
	var b strings.Builder
	if err := t.template.tree.RenderTree(&b); err != nil {
		fmt.Println(err)
	}
	fmt.Println(b.String())
}

func (t TemplateInspector) ToJSON() (string, error) {
	obj := map[string]any{
		"template": templateMarshaller{
			Nodes: t.template.tree.Node(),
		},
	}
	b, err := json.MarshalIndent(obj, "", "  ")
	return string(b), err
}
