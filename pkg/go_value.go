package tempe

import (
	"fmt"
	"github.com/tniedbala/tempe/pkg/api"
	"reflect"
	"iter"
)

type GoValue struct {
	value reflect.Value
}

func NewGoValue(value any) GoValue {
	v := reflect.ValueOf(value)
	return GoValue{
		value: reflectValue(v),
	}
}

func (v GoValue) Bool() bool {
	if v.value.Kind() == reflect.Bool {
		return v.value.Bool()
	}
	if !v.value.IsValid() {
		return false
	}
	return !v.value.IsZero()
}

func (v GoValue) String() string {
	return fmt.Sprintf("%v", v.value.Interface())
}

func (v GoValue) Iter() (iter.Seq2[api.Value, api.Value], error) {
	switch v.value.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.iterValues(), nil
	default:
		return nil, fmt.Errorf("unable to iterate over <%s>", v.value.Type().Name())
	}
}

func (v GoValue) iterValues() iter.Seq2[api.Value, api.Value] {
	return func(yield func(api.Value, api.Value) bool) {
		for i, v := range v.value.Seq2() {
			index, value := NewGoValue(i), NewGoValue(v)
			if !yield(index, value) {
				return
			}
		}
	}
}

// get reflect.Value wrapping underlying value 
func reflectValue(value reflect.Value) reflect.Value {
	switch value.Kind() {
	case reflect.Interface:
		return reflectValue(value)
	case reflect.Pointer:
		return reflectValue(value)
	default:
		return value
	}
}