// Package slicex extends "golang.org/x/exp/slices".
package slicex

// Map transforms the elements of a slice.
func Map[E, T any](list []E, fn func(E) T) []T {
	ret := make([]T, len(list))
	for i, e := range list {
		ret[i] = fn(e)
	}
	return ret
}
