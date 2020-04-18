package main

import "fmt"

/**
* Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right
* which minimizes the sum of all numbers along its path.
*
* Note: You can only move either down or right at any point in time.
*
* Example:
*
*	Input:
*	[
*		[1,3,1],
*		[1,5,1],
*		[4,2,1]
*	]
*
* Output: 7
* Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*
**/
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minPathSum(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	arr := make([][]int, rowLen)

	for i := range arr {
		arr[i] = make([]int, colLen)
	}

	arr[0][0] = grid[0][0]

	for i := 1; i < rowLen; i++ {
		arr[i][0] = arr[i-1][0] + grid[i][0]
	}

	for i := 1; i < colLen; i++ {
		arr[0][i] = arr[0][i-1] + grid[0][i]
	}

	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			arr[i][j] = grid[i][j] + min(arr[i][j-1], arr[i-1][j])
		}
	}

	return arr[len(arr)-1][len(arr[0])-1]
}

func main() {
	grid1 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid1))
}
