# TODO

Core library:
- [ ] Add logic for whitespace control.
    - [x] Add optional `+/-` operators to statement syntax to override default whitespace handling.
    - [x] Create engine-level options for handling whitespace.
    - [ ] Add unit tests for whitespace options and inline whitespace syntax.
    - [ ] Add docs for whitespace control.
- [ ] Add comment syntax.
    - [ ] Add comment syntax to template grammar.
    - [ ] Create parser node & update visitor logic to capture comments.
    - [ ] Add unit tests for comment whitespace handling.
- [ ] Add break/continue statements for for-loops.
    - [ ] Add break/continue syntax to template grammar.
    - [ ] Create parser nodes & update visitor logic to capture break/continue statements.
    - [ ] Add unit tests for comment whitespace handling.
- [ ] Improve errors/error handling.
- [ ] Review logic around local environment creation & variable scope.
- [ ] Clean up/minimize interfaces.

Implementations:
- [ ] Complete default implementation.
- [ ] Complete [Cel implementation](./engines/cel).
- [ ] Complete [Expr implementation](./engines/expr).
- [ ] Complete [Starlark implementation](./engines/starlark).
- [ ] Research go libraries for a JSONPath implementation.
- [ ] Consider whether individual implementatons should live under [/engines](../engines) or live in separate repos.

Documentation:
- [ ] Add docstrings to go source.
- [ ] Complete [implementation guide](./docs/implementation-guid.md).
- [ ] Improve [template syntax](./docs/template-syntax.md) docs.
- [ ] Add docs for working with [`TemplateInspector`](../tempe/template_inspector.go).
- [ ] Add docs for default implementation.

Testing
- [x] Create test suite.
- [ ] Consider whether [/test](../test) module should really be a package within [/tempe](../tempe).
- [ ] Add additional standard tests to test suite:
    - [ ] Whitespace control.
    - [ ] Template options.
    - [ ] Complex templates.
    - [ ] Invalid template syntax.
- [ ] Add additional functionality to test suite:
    - [ ] Functionality to pass template options from yaml test files.
    - [ ] Add field to indicate whether an error is expected.
- [ ] Add additional tests for `tempe` types outside of test suite scope.
- [ ] Create github CI workflow to run tests.