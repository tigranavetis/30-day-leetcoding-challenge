package main

import "fmt"

/**
*
* Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
*
* Example:
*
* Input:
*
*   1 0 1 0 0
*   1 0 1 1 1
*   1 1 1 1 1
*   1 0 0 1 0
*
* Output: 4
*
**/

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}

		return c
	}

	if b < c {
		return b
	}

	return c
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	rowLen := len(matrix)
	colLen := len(matrix[0])
	dp := make([][]int, rowLen+1)
	dp[0] = make([]int, colLen+1)
	maxLen := 0

	for i := 1; i <= rowLen; i++ {
		dp[i] = make([]int, colLen+1)
		for j := 1; j <= colLen; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				maxLen = max(maxLen, dp[i][j])
			}
		}
	}

	return maxLen * maxLen
}

func main() {
	m1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(m1)) // 4
}
