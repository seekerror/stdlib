// Package redgreen contains a generic implementation of a red-green binary search tree.
package redgreen

import (
	"github.com/seekerror/stdlib/pkg/container"
	"github.com/seekerror/stdlib/pkg/lang"
	"github.com/seekerror/stdlib/pkg/util/mathx"
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
	root, nil *node[K, V]
	cmp       lang.CompareFn[K]
}

// New returns is a self-balancing red-green binary search tree. Not thread-safe.
func New[K constraints.Ordered, V any]() *SearchTree[K, V] {
	return NewT[K, V](lang.Compare[K])
}

// NewT returns is a self-balancing red-green binary search tree. Not thread-safe.
func NewT[K, V any](cmp lang.CompareFn[K]) *SearchTree[K, V] {
	n := &node[K, V]{color: black}
	return &SearchTree[K, V]{
		nil:  n,
		root: n,
		cmp:  cmp,
	}
}

func (t *SearchTree[K, V]) List() lang.Iterator[container.KV[K, V]] {
	return &iterator[K, V]{t: t, next: t.min(t.root)}
}

func (t *SearchTree[K, V]) IsEmpty() bool {
	return t.root == t.nil
}

func (t *SearchTree[K, V]) Find(key K) (V, bool) {
	if n := t.find(t.root, key); n != t.nil {
		return n.value, true
	}
	var v V
	return v, false
}

func (t *SearchTree[K, V]) find(n *node[K, V], key K) *node[K, V] {
	for n != t.nil {
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
	return t.insert(&node[K, V]{key: key, value: value})
}

func (t *SearchTree[K, V]) insert(z *node[K, V]) (V, bool) {
	x := t.root // node being compared to z
	y := t.nil  // y will be parent of z.

	// (1) descend until reaching the node or nil

	for x != t.nil {
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
	if y == t.nil {
		t.root = z // tree was empty
	} else if c := t.cmp(y.key, z.key); c < 0 {
		y.right = z
	} else {
		y.left = z
	}
	z.left = t.nil
	z.right = t.nil
	z.color = red

	// (3) fixup red-black structure

	for z.parent.color == red { // red implies parent not root or nil
		if z.parent == z.parent.parent.left { // is z's parent a left child?
			y := z.parent.parent.right        // y is z's uncle
			if y != t.nil && y.color == red { // are z's parent and uncle both red?
				// case 1
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					// case 2
					z = z.parent
					t.leftRotate(z)
				}
				// case 3
				z.parent.color = black
				z.parent.parent.color = red
				t.rightRotate(z.parent.parent)
			}
		} else { // else: right child
			y := z.parent.parent.left         // y is z's uncle
			if y != t.nil && y.color == red { // are z's parent and uncle both red?
				// case 1
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					// case 2
					z = z.parent
					t.rightRotate(z)
				}
				// case 3
				z.parent.color = black
				z.parent.parent.color = red
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = black

	var v V
	return v, false
}

func (t *SearchTree[K, V]) Remove(key K) (V, bool) {
	if z := t.find(t.root, key); z != t.nil {
		v := z.value
		t.remove(z)
		return v, true
	}
	var v V
	return v, false
}

func (t *SearchTree[K, V]) remove(z *node[K, V]) {
	var x *node[K, V]
	y := z
	original := y.color

	// (1) delete node

	if z.left == t.nil {
		x = z.right
		t.transplant(z, z.right) // replace z by its right child
	} else if z.right == t.nil {
		x = z.left
		t.transplant(z, z.left) // replace z by its left child
	} else {
		y = t.min(z.right) // y is z's successor
		original = y.color
		x = y.right
		if y != z.right { // is y father down the tree?
			t.transplant(y, y.right) // replace y by its right child
			y.right = z.right        // z's right child becomes y's right child
			y.right.parent = y
		} else {
			x.parent = y // in case x is t.nil
		}
		t.transplant(z, y) // replace z by its successor y
		y.left = z.left    // and give z's left child to y, which had no left child
		y.left.parent = y
		y.color = z.color
	}

	if original == red {
		return
	}

	// (2) fixup red-black structure

	for x != t.root && x.color == black {
		if x == x.parent.left { // is x a left child?
			w := x.parent.right // w is x's sibling
			if w.color == red {
				// case 1
				w.color = black
				x.parent.color = red
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				// case 2
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					// case 3
					w.left.color = black
					w.color = red
					t.rightRotate(w)
					w = x.parent.right
				}
				// case 4
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left // w is x's sibling
			if w.color == red {
				// case 1
				w.color = black
				x.parent.color = red
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				// case 2
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					// case 3
					w.right.color = black
					w.color = red
					t.leftRotate(w)
					w = x.parent.left
				}
				// case 4
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				t.rightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = black
}

// Height returns the height of the tree, i.e., the number of nodes on the longest path.
func (t *SearchTree[K, V]) Height() int {
	return t.height(t.root)
}

func (t *SearchTree[K, V]) height(n *node[K, V]) int {
	if n == t.nil {
		return 0
	}
	return 1 + mathx.Max(t.height(n.left), t.height(n.left))
}

// transplant replaces the subtree rooted at u with the subtree rooted at v (which may be nil).
func (t *SearchTree[K, V]) transplant(u, v *node[K, V]) {
	if u.parent == t.nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// leftRotate rotates x to the left, making x.right the root:
//         x                y
//       a   y     ->     x   c
//          b c          a b
func (t *SearchTree[K, V]) leftRotate(x *node[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != t.nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil {
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
	if y.right != t.nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// min returns the minimum node, rooted at x.
func (t *SearchTree[K, V]) min(x *node[K, V]) *node[K, V] {
	if x != t.nil {
		for x.left != t.nil {
			x = x.left
		}
	}
	return x
}

// max returns the maximum node, rooted at x.
func (t *SearchTree[K, V]) max(x *node[K, V]) *node[K, V] {
	if x != t.nil {
		for x.right != t.nil {
			x = x.right
		}
	}
	return x
}

func (t *SearchTree[K, V]) predecessor(x *node[K, V]) *node[K, V] {
	if x.left != t.nil {
		return t.max(x.left) // right-most node in left subtree
	}
	// else: find the lowest ancestor of x whose right child is an ancestor of x
	y := x.parent
	for y != t.nil && x == y.left {
		x = y
		y = y.parent
	}
	return y
}

func (t *SearchTree[K, V]) successor(x *node[K, V]) *node[K, V] {
	if x.right != t.nil {
		return t.min(x.right) // left-most node in right subtree
	}
	// else: find the lowest ancestor of x whose left child is an ancestor of x
	y := x.parent
	for y != t.nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (t *SearchTree[K, V]) String() string {
	return lang.IteratorToString(t.List())
}

type iterator[K, V any] struct {
	t    *SearchTree[K, V]
	next *node[K, V]
}

func (it *iterator[K, V]) Next() (container.KV[K, V], bool) {
	if it.next == it.t.nil {
		return container.KV[K, V]{}, false
	}
	cur := it.next
	it.next = it.t.successor(cur)

	return container.KV[K, V]{K: cur.key, V: cur.value}, true
}
