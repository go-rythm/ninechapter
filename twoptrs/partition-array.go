package twoptrs

/**
 * @param nums: The integer array you should partition
 * @param k: An integer
 * @return: The index after partition
 */
func PartitionArray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		for left <= right && nums[left] < k {
			left++
		}
		for left <= right && nums[right] >= k {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return left
}
