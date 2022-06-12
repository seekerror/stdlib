package container

import (
	"fmt"
	"strings"
)

// Container is an abstract container of elements.
type Container[T any] interface {
	// List returns an iterator over all elements.
	List() Iterator[T]
	// IsEmpty returns true if the container is empty.
	IsEmpty() bool
}

// Iterator allows a caller-controlled over a set of elements. It is expected to be lazily-computed,
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

// ToString prints all iterator values. Debugging convenience.
func ToString[T any](it Iterator[T]) string {
	ret := ToList(Map(it, func(t T) string {
		return fmt.Sprint(t)
	}))
	return fmt.Sprintf("[%v]", strings.Join(ret, ", "))
}
