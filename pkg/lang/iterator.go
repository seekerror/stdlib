package lang

import (
	"fmt"
	"strings"
)

// Iterator allows a caller-controlled iteration over a set of elements. It is expected to be lazily-computed,
// such that the cost is proportional to the number of elements consumed.
type Iterator[T any] interface {
	Next() (T, bool)
}

// Map transforms values of an iterator lazily.
func Map[T, U any](it Iterator[T], fn func(t T) U) Iterator[U] {
	return &mapped[T, U]{it: it, fn: fn}
}

type mapped[T, U any] struct {
	it Iterator[T]
	fn func(t T) U
}

func (it *mapped[T, U]) Next() (U, bool) {
	if t, ok := it.it.Next(); ok {
		return it.fn(t), true
	}
	var u U
	return u, false
}

// Head limits an iterator to N values lazily.
func Head[T any](it Iterator[T], n int) Iterator[T] {
	return &head[T]{it: it, n: n}
}

type head[T any] struct {
	it Iterator[T]
	n  int
}

func (it *head[T]) Next() (T, bool) {
	if it.n > 0 {
		if t, ok := it.it.Next(); ok {
			it.n--
			return t, true
		}
		it.n = 0
	}
	var t T
	return t, false
}

// ToList materializes all iterator values into a list.
func ToList[T any](it Iterator[T]) []T {
	var ret []T
	for {
		if t, ok := it.Next(); ok {
			ret = append(ret, t)
		} else {
			break
		}
	}
	return ret
}

// Equivalent compares all elements.
func Equivalent[T comparable](a, b Iterator[T]) bool {
	return EquivalentT(a, b, Equals[T])
}

// EquivalentT compares all elements using the given equality function.
func EquivalentT[T any](a, b Iterator[T], eq EqualsFn[T]) bool {
	for {
		e, ok := a.Next()
		e2, ok2 := b.Next()

		if ok != ok2 {
			return false
		}
		if !ok {
			return true
		}
		if !eq(e, e2) {
			return false
		}
	}
}

// Sprint prints all iterator values. Debugging convenience.
func Sprint[T any](it Iterator[T]) string {
	ret := ToList(Map(it, func(t T) string {
		return fmt.Sprint(t)
	}))
	return fmt.Sprintf("[%v]", strings.Join(ret, ", "))
}
