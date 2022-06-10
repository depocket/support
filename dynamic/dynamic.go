package dynamic

import (
	"fmt"
	"github.com/samber/lo"
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
		var x T
		return x, fmt.Errorf("unrecognized type name: %s", name)
	}
	return reflect.New(reflect.TypeOf(t).Elem()).Interface().(T), nil
}

func (s *Scheme[T]) TotalServices() int {
	return len(s.types)
}

func (s *Scheme[T]) ServicesNames() []string {
	return lo.Keys[string, T](s.types)
}

func NewScheme[T any]() *Scheme[T] {
	return &Scheme[T]{
		types: map[string]T{},
	}
}
