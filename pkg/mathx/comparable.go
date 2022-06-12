package mathx

// EqualsFn defines equality for arbitrary types. This function is not a method to avoid a boxing
// penalty for every natively-ordered object in a data structure.
type EqualsFn[T any] func(a, b T) bool

func Equals[T comparable](a, b T) bool {
	return a == b
}
