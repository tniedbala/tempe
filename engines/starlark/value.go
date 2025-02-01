package tempo_starlark

import (
	"fmt"
	"iter"

	"github.com/tniedbala/tempe-go/pkg/api"
	"go.starlark.net/starlark"
)

type StarlarkValue struct {
	value starlark.Value
}

func NewValue(value starlark.Value) StarlarkValue {
	return StarlarkValue{value}
}

func (v StarlarkValue) Bool() bool {
	return bool(v.value.Truth())
}

func (v StarlarkValue) String() string {
	str, ok := starlark.AsString(v.value)
	if !ok {
		return v.value.String()
	}
	return str
}

func (v StarlarkValue) Iter() (iter.Seq2[api.Value, api.Value], error) {
	iterable, ok := v.value.(starlark.Iterable)
	if !ok {
		return nil, fmt.Errorf("attempted to iterate over <%s>", v.value.Type())
	}
	seq := func(yield func(api.Value, api.Value) bool) {
		iterator := iterable.Iterate()
		defer iterator.Done()
		var value starlark.Value
		i := 0
		for iterator.Next(&value) {
			index := StarlarkValue{starlark.MakeInt(i)}
			val := StarlarkValue{value}
			if !yield(index, val) {
				return
			}
			i++
		}
	}
	return seq, nil
}
