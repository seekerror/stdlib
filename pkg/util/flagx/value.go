package flagx

import (
	"flag"
	"fmt"
)

// ParseFn is a parser of T values.
type ParseFn[T any] func(str string) (T, error)

// Shim converts a (T, bool) parsing function to a ParseFn.
func Shim[T any](fn func(str string) (T, bool)) ParseFn[T] {
	return func(str string) (T, error) {
		if v, ok := fn(str); ok {
			return v, nil
		}
		var t T
		return t, fmt.Errorf("failed to parse '%v'", str)
	}
}

// Value is a T value.
func Value[T any](name string, value T, usage string, parser ParseFn[T]) *T {
	var val T
	flag.Var(newValue(parser, value, &val), name, usage)
	return &val
}

type value[T any] struct {
	val    *T
	parser ParseFn[T]
}

func newValue[T any](parser ParseFn[T], val T, p *T) *value[T] {
	*p = val
	return &value[T]{
		val:    p,
		parser: parser,
	}
}

func (d *value[T]) String() string {
	if d == nil {
		return ""
	}
	return fmt.Sprintf("%v", *d.val)
}

func (d *value[T]) Set(value string) error {
	l, err := d.parser(value)
	if err != nil {
		return err
	}
	*d.val = l
	return nil
}
