package container

import (
	"fmt"
	"github.com/seekerror/stdlib/pkg/lang"
	"golang.org/x/exp/constraints"
)

// Dictionary is an associative mapping, typically implemented as a hashtable or binary search tree.
type Dictionary[K, V any] interface {
	// List returns an iterator over all elements in an implementation-defined order.
	List() lang.Iterator[KV[K, V]]

	// Find returns the value associated with the key, if present.
	Find(k K) (V, bool)
	// Insert sets the value of the key. Returns prior value, if present.
	Insert(k K, v V) (V, bool)
	// Remove removes the key. Returns removed value, if present.
	Remove(k K) (V, bool)
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
