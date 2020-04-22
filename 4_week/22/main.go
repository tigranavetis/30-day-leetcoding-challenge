package main

import "fmt"

/**
* Given an array of integers and an integer k, you need to find the total number
* of continuous subarrays whose sum equals to k.
*
* Example 1:
*	Input:nums = [1,1,1], k = 2
*	Output: 2
* Note:
*	1. The length of the array is in range [1, 20,000].
*	2. The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].
**/
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	sums := make(map[int]int, len(nums))
	sums[0] = 1

	for _, v := range nums {
		sum += v
		count += sums[sum-k]
		sums[sum]++
	}
	return count
}

func main() {
	inp1 := []int{1, 1, 1, 4, 1, 0, 0, 1}
	k1 := 5
	fmt.Println(subarraySum(inp1, k1)) // 4
}
