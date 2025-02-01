package parser

import (
	"io"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type Text struct {
	text string
}

func NewText(text string) *Text {
	return &Text{text}
}

func (n *Text) Children() []api.TemplateNode {
	return []api.TemplateNode{}
}

func (n *Text) Render(env api.Env, w io.StringWriter) error {
	_, err := w.WriteString(n.text)
	return err
}

func (n Text) String() string {
	return "Text: ..."
}
