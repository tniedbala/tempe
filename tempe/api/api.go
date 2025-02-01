package api

import (
	"io"
	"iter"
)

type Option func(opt Options) (any, error)

type Options interface {
	Set(opt Option) error
	Get(opt Option) (any, bool)
	String() string
}

type TemplateEngine interface {
	Options() Options
	GetOption(opt Option) (any, bool)
	SetOptions(opts ...Option) error
	Read(reader io.Reader) (Template, error)
	ReadFile(path string) (Template, error)
	NewTemplate(src string) (Template, error)
	NewEnv() (Env, error)
	NewValue(value any) (Value, error)
}

type Template interface {
	Engine() TemplateEngine
	Env() Env
	SetEnv(env map[string]any) error
	Render(params ...map[string]any) (string, error)
	Write(w io.StringWriter, params ...map[string]any) error
	Parse(src string) error
	ParseTree() ParseTree
}

type Env interface {
	New() (Env, error)
	Copy() (Env, error)
	Keys() []string
	Map() map[string]Value
	Get(name string) (Value, bool)
	Set(name string, value Value) error
	Update(env Env) error
	Eval(src string) (Value, error)
}

type Value interface {
	New(value any) (Value, error)
	Bool() bool
	String() string
	Iter() (iter.Seq2[Value, Value], error)
}

type ParseTree interface {
	Depth() int
	Length() int
	Parent() ParseTree
	Node() TemplateNode
	Children() iter.Seq2[int, ParseTree]
	PrettyPrint() (string, error)
}

type TemplateNode interface {
	Children() []TemplateNode
	Render(opts Options, env Env, w io.StringWriter) error
	Format() (string, string)
	String() string
}
