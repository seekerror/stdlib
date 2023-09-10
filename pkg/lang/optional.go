package lang

import "fmt"

// Optional represents an optional value of type T. Default value is None.
type Optional[T any] struct {
	t  T
	ok bool
}

func Some[T any](t T) Optional[T] {
	return Optional[T]{t: t, ok: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

// V returns the value. False if not present.
func (o Optional[T]) V() (T, bool) {
	return o.t, o.ok
}

func (o Optional[T]) String() string {
	if !o.ok {
		return "none"
	}
	return fmt.Sprintf("some(%v)", o.t)
}
