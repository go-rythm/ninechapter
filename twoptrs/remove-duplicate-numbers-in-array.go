package twoptrs

import "sort"

func Deduplication(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	i, j := 0, 1
	for i = 0; i < len(nums); i++ {
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		if j >= len(nums) {
			break
		}
		nums[i+1] = nums[j]
	}
	return i + 1
}
