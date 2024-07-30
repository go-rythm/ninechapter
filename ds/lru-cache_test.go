package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLRUCache(t *testing.T) {
	var res []int
	c := NewLRUCache(2)
	c.Set(2, 1)
	c.Set(1, 1)
	res = append(res, c.Get(2))
	c.Set(4, 1)
	res = append(res, c.Get(1))
	res = append(res, c.Get(2))
	assert.Equal(t, []int{1, -1, 1}, res)

	res = []int{}
	c = NewLRUCache(1)
	c.Set(2, 1)
	res = append(res, c.Get(2))
	c.Set(3, 2)
	res = append(res, c.Get(2))
	res = append(res, c.Get(3))
	assert.Equal(t, []int{1, -1, 2}, res)
}
