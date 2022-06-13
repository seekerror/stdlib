package redgreen_test

import (
	"github.com/seekerror/stdlib/pkg/container/redgreen"
	"github.com/seekerror/stdlib/pkg/lang"
	"github.com/seekerror/stdlib/pkg/util/mathx"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSearchTree(t *testing.T) {
	// (1) Empty tree

	rgt := redgreen.New[int, int]()
	assert.Equal(t, 0, rgt.Height())
	assert.Equal(t, 0, len(lang.ToList(rgt.List())))

	_, ok := rgt.Find(1)
	assert.False(t, ok)

	const N = 1000
	const K = 50

	keys := lang.ToList(lang.Head(mathx.Numbers(0), N))
	for k := 0; k < K; k++ {
		mathx.Shuffle(keys)

		// (2) Add 100 new elements in random order.

		for _, key := range keys {
			rgt.Insert(key, k)
		}
		assert.True(t, rgt.Height() < int(2*math.Log2(N))+1)
		assert.Equal(t, N, len(lang.ToList(rgt.List())))

		// (3) Find them

		for i := 0; i < N; i++ {
			v, ok := rgt.Find(i)
			assert.True(t, ok)
			assert.Equal(t, k, v)
		}

		// (4) Remove half

		for _, rm := range keys[:N/2] {
			v, found := rgt.Remove(rm)
			assert.True(t, found)
			assert.Equal(t, k, v)
		}
		assert.Equal(t, N/2, len(lang.ToList(rgt.List())))
	}
}
