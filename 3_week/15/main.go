package main

import "fmt"

/**
* Given an array nums of n integers where n > 1, return an array output such that
* output[i] is equal to the product of all the elements of nums except nums[i].
*
* Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array)
* fits in a 32 bit integer.
*
* Note: Please solve it without division and in O(n).
*
* Follow up:
* Could you solve it with constant space complexity?
* (The output array does not count as extra space for the purpose of space complexity analysis.)
**/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	product := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * product
		product *= nums[i]
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums1))
}
