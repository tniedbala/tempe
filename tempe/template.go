package tempe

import (
	"io"
	"strings"

	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/parser"
)

type Params map[string]any

type Template struct {
	engine api.TemplateEngine
	env    *Env
	parser *parser.TemplateParser
	tree   *parser.ParseTree
}

func NewTemplate(engine api.TemplateEngine, src string) (*Template, error) {
	env, err := engine.NewEnv()
	if err != nil {
		return nil, err
	}
	template := &Template{
		engine: engine,
		env:    NewEnv(env),
		parser: parser.NewTemplateParser(),
	}
	err = template.Parse(src)
	return template, err
}

func (t *Template) Engine() api.TemplateEngine {
	return t.engine
}

func (t *Template) Env() api.Env {
	return t.env
}

func (t *Template) SetEnv(env map[string]any) error {
	e, err := t.engine.NewEnv()
	if err != nil {
		return err
	}
	if err = t.updateEnv(e, env); err != nil {
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

func (t *Template) Render(params ...map[string]any) (string, error) {
	var b strings.Builder
	if err := t.Write(&b, params...); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (t *Template) Write(w io.StringWriter, params ...map[string]any) error {
	localEnv, err := t.env.Copy()
	if err != nil {
		return err
	}
	if err = t.updateEnv(localEnv, params...); err != nil {
		return err
	}
	opts := t.engine.Options()
	return t.tree.Node().Render(opts, localEnv, w)
}

func (t *Template) Inspect() api.TemplateInspector {
	return NewTemplateInspector(t)
}

func (t Template) String() string {
	return "Template{}"
}

func (t *Template) updateEnv(env api.Env, params ...map[string]any) error {
	for _, p := range params {
		for key, val := range p {
			value, err := t.engine.NewValue(val)
			if err != nil {
				return err
			}
			if err = env.Set(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}
