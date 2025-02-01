package tempo_cel

import (
	"fmt"
	"iter"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"github.com/tniedbala/tempe-go/tempe/api"
)

type CelValue struct {
	value ref.Val
}

func NewCelValue(value any) (CelValue, error) {
	val, ok := value.(ref.Val)
	if !ok {
		return CelValue{}, fmt.Errorf("unable to convert value to ref.Val: %v", value)
	}
	return CelValue{val}, nil
}

// TODO: see if cel has rules for truthiness/falsyness
func (v CelValue) Bool() bool {
	if types.IsBool(v.value) {
		return v.value.Value().(bool)
	}
	if v.value.Value() == nil {
		return false
	}
	return false
}

func (v CelValue) String() string {
	// TODO: review cel.FormatCELType() function
	return fmt.Sprintf("%v", v.value.Value())
}

func (v CelValue) Iter() (iter.Seq2[api.Value, api.Value], error) {
	if !v.value.Type().HasTrait(traits.IterableType) {
		// TODO: improve error message
		return nil, fmt.Errorf("cannot iterate over <%s>", v.value.Type().TypeName())
	}
	iter := func(yield func(api.Value, api.Value) bool) {
		iterable := v.value.(traits.Iterable)
		iterator := iterable.Iterator()
		i := 0
		for iterator.HasNext().Value().(bool) {
			val := iterator.Next()
			index, value := CelValue{types.Int(i)}, CelValue{val}
			if !yield(index, value) {
				return
			}
			i++
		}
	}
	return iter, nil
}
