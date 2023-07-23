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

// MapIf transforms selected elements of a slice.
func MapIf[E, T any](list []E, fn func(E) (T, bool)) []T {
	var ret []T
	for _, e := range list {
		if o, ok := fn(e); ok {
			ret = append(ret, o)
		}
	}
	return ret
}
