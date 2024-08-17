package dfs

import "strings"

/**
 * @param n: The number of queens
 * @return: All distinct solutions
 *          we will sort your return value in output
 */
func SolveNQueens(n int) [][]string {
	results := [][]string{}
	if n <= 0 {
		return results
	}

	search(&results, []int{}, n)
	return results
}

// search function handles the recursive backtracking to find all solutions
func search(results *[][]string, cols []int, n int) {
	if len(cols) == n {
		*results = append(*results, drawChessboard(cols))
		return
	}

	for colIndex := 0; colIndex < n; colIndex++ {
		if !isValid(cols, colIndex) {
			continue
		}
		cols = append(cols, colIndex)
		search(results, cols, n)
		cols = cols[:len(cols)-1]
	}
}

// drawChessboard generates the chessboard configuration based on the columns array
func drawChessboard(cols []int) []string {
	chessboard := []string{}
	for i := 0; i < len(cols); i++ {
		row := strings.Repeat(".", len(cols))
		row = row[:cols[i]] + "Q" + row[cols[i]+1:]
		chessboard = append(chessboard, row)
	}
	return chessboard
}

// isValid checks if the current column index is a valid position for the queen
func isValid(cols []int, column int) bool {
	row := len(cols)
	for rowIndex := 0; rowIndex < len(cols); rowIndex++ {
		if cols[rowIndex] == column ||
			rowIndex+cols[rowIndex] == row+column ||
			rowIndex-cols[rowIndex] == row-column {
			return false
		}
	}
	return true
}
