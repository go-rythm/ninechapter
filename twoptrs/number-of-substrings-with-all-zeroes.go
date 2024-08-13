package twoptrs

/**
 * @param str: the string
 * @return: the number of substrings
 */
func StringCount(str string) int {
	if len(str) == 0 {
		return 0
	}
	j := 1
	ans := 0
	for i := 0; i < len(str); i++ {
		if str[i] != '0' {
			continue
		}
		j = max(i+1, j)
		for j < len(str) && str[j] == '0' {
			j++
		}
		ans += j - i
	}
	return ans
}
