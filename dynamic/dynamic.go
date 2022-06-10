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
		return nil, fmt.Errorf("unrecognized type name: %s", name)
	}
	return reflect.New(reflect.TypeOf(t)).Interface().(T), nil
}

func (s *Scheme[T]) TotalServices() int {
	return len[T](s.types)
}

func (s *Scheme[T]) ServicesNames() []string {
	return lo.Keys[string, T](s.types)
}

func NewScheme[T]() *Scheme[T] {
	return &Scheme[T]{
		types: map[string]T{},
	}
}
