package twoptrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add(t *testing.T) {
	add(1)
	add(3)
	add(5)
	assert.Equal(t, find(4), true)  //返回true
	assert.Equal(t, find(7), false) //返回false
}
