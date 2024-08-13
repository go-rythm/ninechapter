package dfs2

import "sort"

// https://www.lintcode.com/problem/combination-sum/

/**
 * @param candidates: A list of integers
 * @param target: An integer
 * @return: A list of lists of integers
 *          we will sort your return value in output
 */
func CombinationSum(candidates []int, target int) [][]int {
	subsets := [][]int{}
	if len(candidates) == 0 {
		return subsets
	}
	dict := removeDuplicatesAndSort(candidates)
	dfsCombinationSum(dict, target, 0, nil, &subsets)
	return subsets
}

func dfsCombinationSum(dict []int, target int, idx int, subset []int, subsets *[][]int) {
	if target == 0 {
		*subsets = append(*subsets, append([]int{}, subset...))
	}
	for i := idx; i < len(dict); i++ {
		if target < dict[i] {
			break
		}
		subset = append(subset, dict[i])
		dfsCombinationSum(dict, target-dict[i], i, subset, subsets)
		subset = subset[:len(subset)-1]
	}
}

func removeDuplicatesAndSort(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	sort.Ints(nums)
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}

/**
 * @param nums: A list of integers.
 * @return: A list of permutations.
 *          we will sort your return value in output
 */
// Permute 生成所有可能的排列组合
func Permute(nums []int) [][]int {
	var results [][]int
	if nums == nil {
		return results
	}

	visited := make([]bool, len(nums))
	dfs3(nums, visited, []int{}, &results)
	return results
}

// dfs 递归生成排列组合
// 递归的定义
func dfs3(nums []int,
	visited []bool,
	permutation []int,
	results *[][]int) {

	// 递归的出口
	if len(permutation) == len(nums) {
		// 深拷贝当前排列并添加到结果集中
		temp := make([]int, len(permutation))
		copy(temp, permutation)
		*results = append(*results, temp)
		return
	}

	// 递归的拆解
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		// 添加当前数字到排列中，并标记为已访问
		permutation = append(permutation, nums[i])
		visited[i] = true
		// 递归继续生成排列
		dfs3(nums, visited, permutation, results)
		// 回溯，移除最后一个数字，并标记为未访问
		visited[i] = false
		permutation = permutation[:len(permutation)-1]
	}
}
