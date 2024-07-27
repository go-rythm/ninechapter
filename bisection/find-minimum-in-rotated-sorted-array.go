package bisection

/**
 * @param nums: a rotated sorted array
 * @return: the minimum number in the array
 */
func FindMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid
		} else {
			end = mid
		}
	}
	min := nums[start]
	if nums[end] < min {
		min = nums[end]
	}
	return min
}
