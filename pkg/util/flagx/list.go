package flagx

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
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

// ParseList parses a comma-separated list of T values. The representation of T
// is assumed not to use comma.
func ParseList[T any](list string, parser ParseFn[T]) ([]T, error) {
	var ret []T
	for _, str := range strings.Split(list, ",") {
		a, err := parser(str)
		if err != nil {
			return nil, err
		}
		ret = append(ret, a)
	}
	return ret, nil
}

// List is a list of T values.
func List[T any](name string, value []T, usage string, parser ParseFn[T]) *[]T {
	var val []T
	flag.Var(newList(parser, value, &val), name, usage)
	return &val
}

type list[T any] struct {
	val    *[]T
	parser ParseFn[T]
}

func newList[T any](parser ParseFn[T], val []T, p *[]T) *list[T] {
	*p = val
	return &list[T]{
		val:    p,
		parser: parser,
	}
}

func (d *list[T]) String() string {
	if d == nil {
		return ""
	}
	return fmt.Sprintf("%v", *d.val)
}

func (d *list[T]) Set(value string) error {
	l, err := ParseList(value, d.parser)
	if err != nil {
		return err
	}
	*d.val = l
	return nil
}

// Int64List is a list of int64s.
func Int64List(name string, value []int64, usage string) *[]int64 {
	return List[int64](name, value, usage, func(str string) (int64, error) {
		return strconv.ParseInt(str, 0, 64)
	})
}

// StringList is a list of strings.
func StringList(name string, value []string, usage string) *[]string {
	return List[string](name, value, usage, func(str string) (string, error) {
		return str, nil
	})
}

// DurationList is a list of time.Duration.
func DurationList(name string, value []time.Duration, usage string) *[]time.Duration {
	return List[time.Duration](name, value, usage, time.ParseDuration)
}
