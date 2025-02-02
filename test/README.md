# Tempe Test Suite
Module containing standard tests and additional testing utilities for tempe implementations.

## Usage:

### Install 
```shell
go get https://github.com/tniedbala/tempe-go/tempe
go get https://github.com/tniedbala/tempe-go/test
```

### Minimum Requirements
A number of standard tests are saved to [./tests](./tests), which can be run against 
any implementation of `TemplateEngine`. Note that the target implementation being tested
must be able to properly infer the following types provided as `any` values to test params:
- Primitives: `bool`, `int`, `float64`, `string`
- Collections: `[]any`, `map[string]any`, where `any` is any primitive.

### Standard Tests
Standard tests can be run using `TestSuite.TestSpec()` to ensure compatibilty with tempe minimum spec. 
All default tests can be run at once, or tests can be run for an individual file within [./tests](./tests):

```go
package default_test

import (
    "testing"
    "github.com/tniedbala/tempe-go/tempe"
    tempe_test "github.com/tniedbala/tempe-go/test"
)

var testSuite tempe_test.TestSuite

func init() {
    engine := tempe.DefaultEngine()
    return tempe_test.NewTestSuite(engine)
}

func TestAll(t *testing.T) {
    testSuite.TestSpec(t)
}

func TestIfStatement(t *testing.T) {
    testSuite.TestSpec(t, "if-statement.yaml")
}
```

### Custom Tests
The test suite will allow you to run tests provided as yaml files (see files in 
[./tests](./tests) directory for test format). Assuming you have your own 
`.yaml` files to a local `./test` directory, you can test individual files or all 
files in the directory:

```go
package custom_test

import (
    "embed"
    "testing"
    "github.com/tniedbala/tempe-go/tempe"
    tempe_test "github.com/tniedbala/tempe-go/test"
)

//go:embed tests
var dirFs embed.FS

//go:embed tests/custom-tests.yaml
var fileFs embed.FS

var testSuite tempe_test.TestSuite

func init() {
    engine := tempe.DefaultEngine()
    return tempe_test.NewTestSuite(engine)
}

func TestFilesInDirectory(t *testing.T) {
    testSuite.TestDir(t, dirFs, "tests/*.yaml")
}

func TestIndividualFile(t *testing.T) {
    filename := "tests/custom-tests.yaml"
    file, err := fileFs.Open(filename)
    if err != nil {
        t.Fatal(err)
        return
    }
    testSuite.TestFile(t, filename, file)
}
```