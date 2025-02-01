package tempe

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/tniedbala/tempe-go/tempe/api"
	"github.com/tniedbala/tempe-go/tempe/opt"
	"github.com/tniedbala/tempe-go/tempe/parser"
)

type TemplateEngine struct {
	env     api.Env
	value   api.Value
	options api.Options
}

func NewTemplateEngine(env api.Env, value api.Value) TemplateEngine {
	return TemplateEngine{
		env:     env,
		value:   value,
		options: opt.DefaultOptions(),
	}
}

func (e TemplateEngine) Options() api.Options {
	return e.options
}

func (e TemplateEngine) GetOption(opt api.Option) (any, bool) {
	return e.options.Get(opt)
}

func (e TemplateEngine) SetOptions(opts ...api.Option) error {
	for _, o := range opts {
		if err := e.options.Set(o); err != nil {
			return err
		}
	}
	return nil
}

func (e TemplateEngine) Read(reader io.Reader) (api.Template, error) {
	var b bytes.Buffer
	if _, err := b.ReadFrom(reader); err != nil {
		return nil, err
	}
	fmt.Println(b)
	return NewTemplate(e, b.String())
}

func (e TemplateEngine) ReadFile(path string) (api.Template, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewTemplate(e, string(b))
}

func (e TemplateEngine) NewTemplate(src string) (api.Template, error) {
	env, err := e.env.New()
	if err != nil {
		return nil, err
	}
	template := &Template{
		engine: e,
		env:    NewEnv(env),
		parser: parser.NewTemplateParser(),
	}
	err = template.Parse(src)
	return template, err
}

func (e TemplateEngine) NewEnv() (api.Env, error) {
	return e.env.New()
}

func (e TemplateEngine) NewValue(value any) (api.Value, error) {
	return e.value.New(value)
}

func (e TemplateEngine) String() string {
	return "TemplateEngine{}"
}
