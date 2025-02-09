package tempe_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/tniedbala/tempe-go/tempe/api"
	"gopkg.in/yaml.v3"
)

type Test struct {
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Params      map[string]any `yaml:"params"`
	Template    string         `yaml:"template"`
	Output      string         `yaml:"output"`
}

type TestFile struct {
	Name  string
	Tests []Test `yaml:"tests"`
}

func NewTestFile(name string, reader io.Reader) (*TestFile, error) {
	file := &TestFile{Name: name}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(reader); err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(buf.Bytes(), file); err != nil {
		return nil, fmt.Errorf("error parsing %s: %w", file.Name, err)
	}
	return file, nil
}

func (f *TestFile) Run(t *testing.T, engine api.TemplateEngine) {
	t.Run(f.Name, func(t *testing.T) {
		f.RunTests(t, engine)
	})
}

func (f *TestFile) RunTests(t *testing.T, engine api.TemplateEngine) {
	for _, test := range f.Tests {
		t.Run(test.Name, func(t *testing.T) {
			template, err := engine.NewTemplate(test.Template)
			if err != nil {
				t.Fatal(err)
				return
			}
			output, err := template.Render(test.Params)
			if err != nil {
				t.Fatal(err)
				return
			}
			if output != test.Output {
				t.Fail()
			}
		})
	}
}
