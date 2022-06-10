package dynamic

import (
	"github.com/samber/lo"
)

var (
	ServiceRegisters = NewScheme[string]()
)

func TotalServices[T]() int {
	return len[T](ServiceRegisters.types)
}

func ServicesNames[T]() []string {
	return lo.Keys[string, T](ServiceRegisters.types)
}
