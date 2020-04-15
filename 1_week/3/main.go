package main

import (
	"fmt"
	"math"
)

// maxSubArray finds the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum. Complexity: O(n)
func maxSubArray(nums []int) int {
	curr := 0
	max := int(math.Inf(-1))
	for _, v := range nums {
		curr += v
		if curr > max {
			max = curr
		}

		if curr < 0 {
			curr = 0
		}
	}

	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, -12, -5}))
}
