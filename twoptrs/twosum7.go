package twoptrs

import (
	"math"
)

func TwoSum7(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	target = int(math.Abs(float64(target)))
	j := 1
	for i := 0; i < len(nums); i++ {
		j = max(j, i+1)
		// 找到第一个 nums[j]-nums[i] >= target 的 j
		for j < len(nums) && nums[j]-nums[i] < target {
			j++
		}
		if j >= len(nums) {
			break
		}
		if nums[j]-nums[i] == target {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{-1, -1}
}
