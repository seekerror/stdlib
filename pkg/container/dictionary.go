package container

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

/*
// Direction is a traversal direction.
type Direction int

const (
	Ascending Direction = iota
	Descending
)

// SearchTree is an associative mapping with an ordered domain.
type SearchTree[K, V any] interface {
	Dictionary[K, V]

	// Min returns the smallest key, if not empty.
	Min() (K, bool)
	// Max returns the largest key, if not empty.
	Max() (K, bool)

	// Traverse returns an iterator over all elements in key-order.
	Traverse(direction Direction) Iterator[KV[K, V]]
}
*/

// Dictionary is an associative mapping, typically implemented as a hashtable or binary search tree.
type Dictionary[K, V any] interface {
	// List returns an iterator over all elements in an implementation-defined order.
	List() Iterator[KV[K, V]]

	// Find returns the value associated with the key, if present.
	Find(k K) (V, bool)
	// Insert sets the value of the key.
	Insert(k K, v V) (V, bool)
	// Remove removes the key. Returns true iff an element was removed.
	Remove(k K) bool
}

// KV is a key-value pair.
type KV[K, V any] struct {
	K K
	V V
}

func (kv KV[K, V]) String() string {
	return fmt.Sprintf("%v:%v", kv.K, kv.V)
}

// InsertAll inserts all map elements in the dictionary.
func InsertAll[K constraints.Ordered, V any](m map[K]V, dict Dictionary[K, V]) {
	for k, v := range m {
		dict.Insert(k, v)
	}
}
