package tempe_test

import (
	"embed"
	"io"
	"io/fs"
	"testing"

	"github.com/tniedbala/tempe-go/tempe/api"
)

//go:embed tests
var testFS embed.FS

// TestDir is an fs.FS containing standard test .yaml files.
var TestDir fs.FS

func init() {
	var err error 
	TestDir, err = fs.Sub(testFS, "tests")
	if err != nil {
		panic("test directory not found")
	}
}

type TestSuite struct {
	Engine api.TemplateEngine
}

func NewTestSuite(engine api.TemplateEngine) TestSuite {
	return TestSuite{Engine: engine}
}

func (s TestSuite) TestSpec(t *testing.T, filenames ...string) {
	if len(filenames) == 0 {
		s.TestDir(t, TestDir, "*.yaml")
	}
	for _, filename := range filenames {
		file, err := TestDir.Open(filename)
		if err != nil {
			t.Fatal(err)
			return
		}
		s.TestFile(t, filename, file)
		if err = file.Close(); err != nil {
			t.Error(err)
		}
	}
}

func (s TestSuite) TestFile(t *testing.T, name string, reader io.Reader) {
	file, err := NewTestFile(name, reader)
	if err != nil {
		t.Fatal(err)
		return
	}
	file.Run(t, s.Engine)
}

func (s TestSuite) TestDir(t *testing.T, dir fs.FS, glob string) {
	filenames, err := fs.Glob(dir, glob)
	if err != nil {
		t.Fatal(err)
	}
	for _, filename := range filenames {
		file, err := dir.Open(filename)
		if err != nil {
			t.Fatal(err)
			continue
		}
		s.TestFile(t, filename, file)
		if err = file.Close(); err != nil {
			t.Error(err)
		}
	}
}
