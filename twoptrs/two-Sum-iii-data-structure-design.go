package twoptrs

var counter = make(map[int]int)

/**
 * @param number: An integer
 * @return: nothing
 */
func add(number int) {
	counter[number]++
}

/**
 * @param value: An integer
 * @return: Find if there exists any pair of numbers which sum is equal to the value.
 */
func find(value int) bool {
	for num1 := range counter {
		num2 := value - num1
		wantCnt := 1
		if num1 == num2 {
			wantCnt = 2
		}
		if counter[num2] >= wantCnt {
			return true
		}
	}
	return false
}
