package tempe

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/parser"
)

type Template struct {
	engine api.TemplateEngine
	env    *Env
	parser *parser.TemplateParser
	tree   *parser.ParseTree
}

// TODO: this is incorrect, need to review how to read from reader of unknown size
func Read(engine api.TemplateEngine, reader io.Reader) (*Template, error) {
	var b []byte
	if _, err := reader.Read(b); err != nil {
		return nil, err
	}
	fmt.Println(b)
	return NewTemplate(engine, string(b))
}

func ReadFile(engine api.TemplateEngine, path string) (*Template, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewTemplate(engine, string(b))
}

func NewTemplate(engine api.TemplateEngine, src string) (*Template, error) {
	template := &Template{
		engine: engine,
		env:    NewEnv(engine.DefaultEnv()),
		parser: parser.NewTemplateParser(),
	}
	err := template.Parse(src)
	return template, err
}

func (t *Template) Engine() api.TemplateEngine {
	return t.engine
}

func (t *Template) Env() api.Env {
	return t.env
}

func (t *Template) SetEnv(env map[string]any) error {
	e, err := t.engine.NewEnv(env)
	if err != nil {
		return err
	}
	t.env = NewEnv(e)
	return nil
}

func (t *Template) Parse(src string) error {
	tree, err := t.parser.Parse(src)
	if err != nil {
		return err
	}
	t.tree = tree
	return nil
}

func (t *Template) ParseTree() api.ParseTree {
	return t.tree
}

func (t *Template) Render(params map[string]any) (string, error) {
	var b bytes.Buffer
	if err := t.Write(params, &b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (t *Template) Write(params map[string]any, w io.StringWriter) error {
	localEnv, err := t.env.LocalEnv()
	if err != nil {
		return err
	}
	for key, val := range params {
		value, err := t.engine.NewValue(val)
		if err != nil {
			return err
		}
		localEnv.Set(key, value)
	}
	return t.tree.Node().Render(localEnv, w)
}

func (t Template) String() string {
	return "Template{}"
}
