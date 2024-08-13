package twoptrs

import (
	"fmt"
	"testing"
)

func TestDeduplication(t *testing.T) {
	nums := []int{3, 3, 2, 1, 4, 2, 1}
	newLength := Deduplication(nums)
	fmt.Println(newLength)        // 输出: 4
	fmt.Println(nums[:newLength]) // 输出: [1 2 3 4]
}
