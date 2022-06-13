// Package lang provides extensions and utilities to builtin language features.
package lang

import "golang.org/x/exp/constraints"

// EqualsFn defines equality for arbitrary types. This function is not a method to avoid a boxing
// penalty for every natively-ordered object in a data structure.
type EqualsFn[T any] func(a, b T) bool

func Equals[T comparable](a, b T) bool {
	return a == b
}

// CompareFn defines ordering for arbitrary types, returning -1 if a < b, 0 if a == b and 1 if a > b. This
// function is not a method to avoid a boxing penalty for every natively-ordered object in a data structure.
type CompareFn[T any] func(a, b T) int

func Compare[T constraints.Ordered](a, b T) int {
	switch {
	case a < b:
		return -1
	case b < a:
		return 1
	default:
		return 0
	}
}

func Reverse[T any](fn CompareFn[T]) CompareFn[T] {
	return func(a, b T) int {
		return fn(b, a)
	}
}
