package mathx

import "golang.org/x/exp/constraints"

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

// Min returns the smallest element in the list. If empty, returns the default value. Convenience function.
func Min[T constraints.Ordered](list ...T) T {
	return MinT[T](Compare[T], list...)
}

// MinT returns the smallest element in the list, using less. If empty, returns the default value of T.
func MinT[T any](cmp CompareFn[T], list ...T) T {
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
	return MaxT[T](Compare[T], list...)
}

// MaxT returns the largest element in the list, using less. If empty, returns the default value of T.
func MaxT[T any](cmp CompareFn[T], list ...T) T {
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
