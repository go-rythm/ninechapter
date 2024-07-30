package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataStream(t *testing.T) {
	ds := NewDataStream()
	ds.Add(1)
	ds.Add(2)
	assert.Equal(t, 1, ds.FirstUnique())
	ds.Add(1)
	assert.Equal(t, 2, ds.FirstUnique())
}
