package twoptrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum7(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	target := 4
	result := TwoSum7(nums, target)
	assert.Equal(t, []int{1, 5}, result)
}
