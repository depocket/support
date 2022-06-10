package dynamic

import (
	"fmt"
	"reflect"
)

type Scheme[T any] struct {
	types map[string]T
}

func (s *Scheme[T]) RegisterService(name string, t T) {
	s.types[name] = t
}

func (s *Scheme[T]) New(name string) (T, error) {
	t, ok := s.types[name]
	if !ok {
		return nil, fmt.Errorf("unrecognized type name: %s", name)
	}
	return reflect.New(reflect.TypeOf(t)).Interface().(T), nil
}

func NewScheme[T]() *Scheme[T] {
	return &Scheme[T]{
		types: map[string]T{},
	}
}
