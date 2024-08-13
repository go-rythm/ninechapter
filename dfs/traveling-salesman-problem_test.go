package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCost(t *testing.T) {
	n := 3
	roads := [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}
	assert.Equal(t, 3, minCost3(n, roads))
}
