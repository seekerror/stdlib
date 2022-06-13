// Package mathx provides extensions and utilities to the math package.
package mathx

import (
	"github.com/seekerror/stdlib/pkg/lang"
	"golang.org/x/exp/constraints"
)

// Numbers is an infinite number generator, returning from, from+1, from+2, ...
func Numbers[T constraints.Integer](from T) lang.Iterator[T] {
	return &numbers[T]{n: from}
}

type numbers[T constraints.Integer] struct {
	n T
}

func (it *numbers[T]) Next() (T, bool) {
	ret := it.n
	it.n++
	return ret, true
}

// Min returns the smallest element in the list. If empty, returns the default value. Convenience function.
func Min[T constraints.Ordered](list ...T) T {
	return MinT[T](lang.Compare[T], list...)
}

// MinT returns the smallest element in the list, using the given ordering. If empty, returns
// the default value of T.
func MinT[T any](cmp lang.CompareFn[T], list ...T) T {
	var min T
	first := true
	for _, t := range list {
		if first || cmp(t, min) < 0 {
			min = t
		}
		first = false
	}
	return min
}

// Max returns the largest element in the list. If empty, returns the default value. Convenience function.
func Max[T constraints.Ordered](list ...T) T {
	return MaxT[T](lang.Compare[T], list...)
}

// MaxT returns the largest element in the list, using the given ordering. If empty, returns
// the default value of T.
func MaxT[T any](cmp lang.CompareFn[T], list ...T) T {
	var max T
	first := true
	for _, t := range list {
		if first || cmp(max, t) < 0 {
			max = t
		}
		first = false
	}
	return max
}
