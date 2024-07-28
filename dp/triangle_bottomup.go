package dp

// https://www.lintcode.com/problem/triangle/

/**
 * @param triangle: a list of lists of integers
 * @return: An integer, minimum path sum
 */

func MinimumTotal2(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}

	for i := 0; i < n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
		}
	}

	return dp[0][0]
}
