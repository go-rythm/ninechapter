package bfs4

import (
	"sort"
)

// https://www.lintcode.com/problem/subsets

/**
 * @param nums: A set of numbers
 * @return: A list of lists
 *          we will sort your return value in output
 */
func Subsets(nums []int) [][]int {
	if nums == nil {
		return [][]int{}
	}

	sort.Ints(nums)
	queue := [][]int{{}}
	index := 0

	for index < len(queue) {
		subset := queue[index]
		index++
		for i := 0; i < len(nums); i++ {
			if len(subset) != 0 && subset[len(subset)-1] >= nums[i] {
				continue
			}
			newSubset := append([]int(nil), subset...)
			newSubset = append(newSubset, nums[i])
			queue = append(queue, newSubset)
		}
	}

	return queue
}

func Subsets2(nums []int) [][]int {
	if nums == nil {
		return [][]int{}
	}

	// 对数组进行排序
	sort.Ints(nums)
	queue := [][]int{{}}

	// 遍历每个数字
	for _, num := range nums {
		size := len(queue)
		// 遍历现有的每个子集
		for i := 0; i < size; i++ {
			// 创建新的子集
			newSubset := append([]int(nil), queue[i]...)
			newSubset = append(newSubset, num)
			queue = append(queue, newSubset)
		}
	}

	return queue
}
