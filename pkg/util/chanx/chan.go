// Package chanx contains generic channel utilities.
package chanx

import (
	"github.com/seekerror/stdlib/pkg/lang"
)

// NewFixed returns a closed chan with the given values.
func NewFixed[T any](list ...T) <-chan T {
	ret := make(chan T, len(list))
	for _, m := range list {
		ret <- m
	}
	close(ret)
	return ret
}

// ToList materializes all chan values into a list. Blocking.
func ToList[T any](ch <-chan T) []T {
	var ret []T
	for m := range ch {
		ret = append(ret, m)
	}
	return ret
}

// ToIterator transforms a chan into an iterator. Values are retrieved lazily.
func ToIterator[T any](ch <-chan T) lang.Iterator[T] {
	return &chanIterator[T]{ch: ch}
}

type chanIterator[T any] struct {
	ch <-chan T
}

func (i *chanIterator[T]) Next() (T, bool) {
	v, ok := <-i.ch
	return v, ok
}

// Map transforms values of a chan. Will leak a go routine if input is not closed.
func Map[T, U any](ch <-chan T, fn func(t T) U) <-chan U {
	out := make(chan U, 1)
	go func() {
		defer close(out)

		for t := range ch {
			out <- fn(t)
		}
	}()

	return out
}

// MapIf transforms selected values of a chan. Will leak a go routine if input is not closed.
func MapIf[T, U any](ch <-chan T, fn func(t T) (U, bool)) <-chan U {
	out := make(chan U, 1)
	go func() {
		defer close(out)

		for t := range ch {
			if o, ok := fn(t); ok {
				out <- o
			}
		}
	}()

	return out
}

// Breaker adds a quit chan breaker to a chan.
func Breaker[T any](in <-chan T, quit <-chan struct{}) <-chan T {
	out := make(chan T, 1)
	go func() {
		defer close(out)

		for {
			select {
			case msg, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- msg:
					//ok
				case <-quit:
					return
				}
			case <-quit:
				return
			}
		}
	}()

	return out
}
