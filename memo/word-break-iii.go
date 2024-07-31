package memo

import (
	"strings"
)

func WordBreak3(s string, dict map[string]struct{}) int {
	n := len(s)
	lowerS := strings.ToLower(s)

	lowerDict := make(map[string]struct{})
	for str := range dict {
		lowerDict[strings.ToLower(str)] = struct{}{}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sub := lowerS[i : j+1]
			if _, exists := lowerDict[sub]; exists {
				dp[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := i; k < j; k++ {
				dp[i][j] += dp[i][k] * dp[k+1][j]
			}
		}
	}

	return dp[0][n-1]
}
