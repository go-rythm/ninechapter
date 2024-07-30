package ds

/**
 * @param nums: a continuous stream of numbers
 * @param number: a number
 * @return: returns the first unique number
 */
func FirstUniqueNumber(nums []int, number int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
		if num == number {
			break
		}
	}
	if _, ok := counter[number]; !ok {
		return -1
	}
	for _, num := range nums {
		if counter[num] == 1 {
			return num
		}
		if num == number {
			break
		}
	}
	return -1
}
