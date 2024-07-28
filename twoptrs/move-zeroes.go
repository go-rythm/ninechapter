package twoptrs

/**
 * @param nums: an integer array
 * @return: nothing
 */
func MoveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	for ; left < len(nums); left++ {
		if nums[left] != 0 {
			nums[left] = 0
		}
	}
}
