package tempo_cel

import (
	"io"
	"github.com/tniedbala/tempe/pkg"
	"github.com/tniedbala/tempe/pkg/api"
)

type CelEngine struct {
	
}

func NewEngine() *CelEngine {
	return &CelEngine{}
}

func (e *CelEngine) Read(reader io.Reader) (api.Template, error) {
	return tempe.Read(e, reader)
}

func (e *CelEngine) ReadFile(path string) (api.Template, error) {
	return tempe.ReadFile(e, path)
}

func (e *CelEngine) NewTemplate(src string) (api.Template, error) {
	return tempe.NewTemplate(e, src)
}

func (e *CelEngine) DefaultEnv() api.Env {
	return NewEnv(e)
}

func (e *CelEngine) NewEnv(env map[string]any) (api.Env, error) {
	for _, v := range env {
		_, err := e.NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (e *CelEngine) NewValue(value any) (api.Value, error) {
	return nil, nil
}

func (e *CelEngine) String() string {
	return "CelEngine{}"
}
