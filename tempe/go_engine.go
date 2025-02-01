package tempe

import (
	"io"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type GoEngine struct {}

func DefaultEngine() GoEngine {
	return NewGoEngine()
}

func NewGoEngine() GoEngine {
	return GoEngine{}
}

func (e GoEngine) Read(reader io.Reader) (api.Template, error) {
	return Read(e, reader)
}

func (e GoEngine) ReadFile(path string) (api.Template, error) {
	return ReadFile(e, path)
}

func (e GoEngine) NewTemplate(src string) (api.Template, error) {
	return NewTemplate(e, src)
}

func (e GoEngine) DefaultEnv() api.Env {
	return NewGoEnv(nil)
}

func (e GoEngine) NewEnv(env map[string]any) (api.Env, error) {
	return NewGoEnv(env), nil
}

func (e GoEngine) NewValue(value any) (api.Value, error) {
	return NewGoValue(value), nil
}

func (e GoEngine) String() string {
	return "GoEngine{}"
}
