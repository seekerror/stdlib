package redgreen_test

import (
	"github.com/seekerror/stdlib/pkg/container/redgreen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree(t *testing.T) {
	rgt := redgreen.New[int, int]()

	_, ok := rgt.Find(1)
	assert.False(t, ok)

	for i := 1; i < 10; i++ {
		rgt.Insert(i, 10*i)
		println(rgt.String())
	}

	for i := 1; i < 10; i++ {
		v, ok := rgt.Find(i)
		assert.True(t, ok)
		assert.Equal(t, v, 10*i)
	}
}
