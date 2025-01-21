package api

import (
	"io"
	"iter"
)

type TemplateEngine interface {
	Read(reader io.Reader) (Template, error)
	ReadFile(path string) (Template, error)
	NewTemplate(src string) (Template, error)
	DefaultEnv() Env
	NewEnv(env map[string]any) (Env, error)
	NewValue(value any) (Value, error)
}

type Template interface {
	Engine() TemplateEngine
	Env() Env
	SetEnv(env map[string]any) error
	Parse(src string) error
	ParseTree() ParseTree
	Render(params map[string]any) (string, error)
	Write(params map[string]any, w io.StringWriter) error
}

type Env interface {
	LocalEnv() (Env, error)
	Keys() []string
	Map() map[string]Value
	Get(name string) (Value, bool)
	Set(name string, value Value) error
	Update(env Env) error
	Eval(src string) (Value, error)
}

type Value interface {
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
	Render(env Env, w io.StringWriter) error
	String() string
}
