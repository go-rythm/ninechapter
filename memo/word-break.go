package memo

/**
 * @param s: A string
 * @param wordSet: A dictionary of words dict
 * @return: A boolean
 */

func WordBreak(s string, wordSet map[string]struct{}) bool {
	if s == "" {
		return true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	maxLength := 0
	for word := range wordSet {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	for i := 1; i <= n; i++ {
		for l := 1; l <= maxLength; l++ {
			if i < l {
				break
			}
			if !dp[i-l] {
				continue
			}
			word := s[i-l : i]
			if _, ok := wordSet[word]; ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
