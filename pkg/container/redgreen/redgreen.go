// Package redgreen contains a generic implementation of a red-green binary search tree.
package redgreen

import (
	"github.com/seekerror/stdlib/pkg/container"
	"github.com/seekerror/stdlib/pkg/mathx"
	"golang.org/x/exp/constraints"
)

// color is {red, black}
type color bool

const (
	red   color = true
	black color = false
)

// node is the internal red-green tree node.
type node[K, V any] struct {
	parent, left, right *node[K, V]

	key   K
	value V

	color color
}

// SearchTree is a red-green binary search tree, based on CLRS 4th edition. Not thread-safe.
type SearchTree[K, V any] struct {
	root *node[K, V]
	cmp  mathx.CompareFn[K]
}

// New returns is a self-balancing red-green binary search tree. Not thread-safe.
func New[K constraints.Ordered, V any]() *SearchTree[K, V] {
	return &SearchTree[K, V]{
		cmp: mathx.Compare[K],
	}
}

// NewT returns is a self-balancing red-green binary search tree. Not thread-safe.
func NewT[K, V any](cmp mathx.CompareFn[K]) *SearchTree[K, V] {
	return &SearchTree[K, V]{
		cmp: cmp,
	}
}

func (t *SearchTree[K, V]) List() container.Iterator[container.KV[K, V]] {
	if t.root == nil {
		return &iterator[K, V]{}
	}
	return &iterator[K, V]{next: min(t.root)}
}

func (t *SearchTree[K, V]) Find(key K) (V, bool) {
	if n := t.find(t.root, key); n != nil {
		return n.value, true
	}
	var v V
	return v, false
}

func (t *SearchTree[K, V]) find(n *node[K, V], key K) *node[K, V] {
	for n != nil {
		c := t.cmp(n.key, key)
		if c == 0 {
			break
		}
		if c < 0 {
			n = n.right
		} else {
			n = n.left
		}
	}
	return n
}

func (t *SearchTree[K, V]) Insert(key K, value V) (V, bool) {
	return t.insert(&node[K, V]{key: key, value: value, color: red})
}

func (t *SearchTree[K, V]) insert(z *node[K, V]) (V, bool) {
	x := t.root       // node being compared to z
	var y *node[K, V] // y will be parent of z.

	// (1) descend until reaching the node or nil

	for x != nil {
		y = x
		c := t.cmp(x.key, z.key)
		if c == 0 {
			ret := x.value
			x.value = z.value
			return ret, true
		}
		if c < 0 {
			x = x.right
		} else {
			x = x.left
		}
	}

	// (2) found the location for new value -- insert z with parent y

	z.parent = y
	if y == nil {
		t.root = z // tree was empty
	} else if c := t.cmp(y.key, z.key); c < 0 {
		y.right = z
	} else {
		y.left = z
	}

	// (3) fixup red-black structure

	var v V
	return v, false
}

func (t *SearchTree[K, V]) Remove(key K) bool {
	return false
}

// transplant replaces the subtree rooted at u with the subtree rooted at v (which may be nil).
func (t *SearchTree[K, V]) transplant(u, v *node[K, V]) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// leftRotate rotates x to the left, making x.right the root:
//         x                y
//       a   y     ->     x   c
//          b c          a b
func (t *SearchTree[K, V]) leftRotate(x *node[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// rightRotate rotates x to the right, making x.left the root:
//         x               y
//       y   c   ->      a   x
//      a b                 b c
func (t *SearchTree[K, V]) rightRotate(x *node[K, V]) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (t *SearchTree[K, V]) String() string {
	return container.ToString(t.List())
}

type iterator[K, V any] struct {
	next *node[K, V]
}

func (it *iterator[K, V]) Next() (container.KV[K, V], bool) {
	if it.next == nil {
		return container.KV[K, V]{}, false
	}
	cur := it.next
	it.next = successor(cur)

	return container.KV[K, V]{K: cur.key, V: cur.value}, true
}

// min returns the minimum node, rooted at x.
func min[K, V any](x *node[K, V]) *node[K, V] {
	for x.left != nil {
		x = x.left
	}
	return x
}

// max returns the maximum node, rooted at x.
func max[K, V any](x *node[K, V]) *node[K, V] {
	for x.right != nil {
		x = x.right
	}
	return x
}

func predecessor[K, V any](x *node[K, V]) *node[K, V] {
	if x.left != nil {
		return max(x.left) // right-most node in left subtree
	}
	// else: find the lowest ancestor of x whose right child is an ancestor of x
	y := x.parent
	for y != nil && x == y.left {
		x = y
		y = y.parent
	}
	return y
}

func successor[K, V any](x *node[K, V]) *node[K, V] {
	if x.right != nil {
		return min(x.right) // left-most node in right subtree
	}
	// else: find the lowest ancestor of x whose left child is an ancestor of x
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}
