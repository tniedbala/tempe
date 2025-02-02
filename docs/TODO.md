# TODO
Core library:
- [x] Add logic for whitespace control.
- [ ] Add comment syntax.
- [ ] Add break/continue syntax for for-loops.
- [ ] Improve errors/error handling.
- [ ] Review logic around local environment creation & variable scope.
- [ ] Clean up/minimize interfaces.

Implementations:
- [ ] Complete default implementation.
- [ ] Complete [Cel implementation](./engines/cel).
- [ ] Complete [Expr implementation](./engines/expr).
- [ ] Complete [Starlark implementation](./engines/starlark).

Documentation:
- [ ] Add docstrings to go source.
- [ ] Complete [implementation guide](./docs/implementation-guid.md).
- [ ] Improve [template syntax](./docs/template-syntax.md) docs.
- [ ] Add docs for whitespace control.

Testing
- [x] Create test suite.
- [ ] Add additional default tests.
- [ ] Create github CI workflow to run tests.