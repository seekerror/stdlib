package container

import "github.com/seekerror/stdlib/pkg/lang"

// Container is an abstract container of elements.
type Container[T any] interface {
	// List returns an iterator over all elements.
	List() lang.Iterator[T]
	// IsEmpty returns true if the container is empty.
	IsEmpty() bool
}
