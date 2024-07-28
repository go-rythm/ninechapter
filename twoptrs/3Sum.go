package twoptrs

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		findTwoSum(nums, i, &res)
	}
	return res
}

func findTwoSum(nums []int, idx int, res *[][]int) {
	left := idx + 1
	right := len(nums) - 1
	target := -nums[idx]
	for left < right {
		tmp := nums[left] + nums[right]
		if tmp < target {
			left++
		} else if tmp > target {
			right--
		} else {
			*res = append(*res, []int{nums[idx], nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
		}
	}
}
