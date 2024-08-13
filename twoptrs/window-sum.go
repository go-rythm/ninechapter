package twoptrs

func WinSum(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 || len(nums) < k {
		return res
	}

	j := 0
	windowSum := 0
	for i := 0; i < len(nums); i++ {
		for j < len(nums) && j-i < k {
			windowSum += nums[j]
			j++
		}
		if j-i == k {
			res = append(res, windowSum)
		}
		windowSum -= nums[i]
	}
	return res
}
