package bisection

/**
 * @param nums: a mountain sequence which increase firstly and then decrease
 * @return: then mountain top
 */
func MountainSequence(nums []int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid+1] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[end]
	}
	return nums[start]
}
