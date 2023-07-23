package slicex_test

import (
	"github.com/seekerror/stdlib/pkg/util/slicex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	strs := slicex.Map([]int{1, 2}, strconv.Itoa)
	require.Len(t, strs, 2)
	assert.Equal(t, strs[0], "1")
	assert.Equal(t, strs[1], "2")
}
