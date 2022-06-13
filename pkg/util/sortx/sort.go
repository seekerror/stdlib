// Package sortx provides extensions and utilities to the sort package.
package sortx

import (
	"github.com/seekerror/stdlib/pkg/lang"
	"golang.org/x/exp/constraints"
	"sort"
)

// Sort sorts the given list, using the default ordering. Convenience function.
func Sort[T constraints.Ordered](list []T) {
	SortT[T](list, lang.Compare[T])
}

// SortT sorts the given list, using the given ordering. Convenience function.
func SortT[T any](list []T, cmp lang.CompareFn[T]) {
	sort.Slice(list, func(i, j int) bool {
		return cmp(list[i], list[j]) < 0
	})
}
