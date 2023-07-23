// Package slicex extends "golang.org/x/exp/slices".
package slicex

// Map transforms the elements of a slice.
func Map[E, T any](list []E, fn func(E) T) []T {
	ret := make([]T, len(list))
	for _, e := range list {
		ret = append(ret, fn(e))
	}
	return ret
}
