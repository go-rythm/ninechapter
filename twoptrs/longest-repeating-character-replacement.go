package twoptrs

/**
 * @param s: a string
 * @param k: a integer
 * @return: return a integer
 */
func CharacterReplacement(s string, k int) int {
	counter := make(map[byte]int)
	ans := 0
	// j 指向满足条件的子字符串的后面一个位置
	j := 0
	maxFreq := 0
	for i := 0; i < len(s); i++ {
		// 找到 i ~ j-1 这一段是最长的以i开头的满足
		// 不超过 k 次替换的 substring
		for j < len(s) && j-i-maxFreq <= k {
			counter[s[j]]++
			maxFreq = max(maxFreq, counter[s[j]])
			j++
		}

		if j-i-maxFreq > k {
			// i ~ j-1 的 substring 需要 k+1
			// i ~ j-2 的 substring 只需要 k 次替换
			ans = max(ans, j-2-i+1)
		} else {
			ans = max(ans, j-i)
		}

		counter[s[i]]--
		maxFreq = 0
		for _, cnt := range counter {
			if cnt > maxFreq {
				maxFreq = cnt
			}
		}
	}
	return ans
}
