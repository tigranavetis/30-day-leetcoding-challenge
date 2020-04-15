package main

import "fmt"

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

// With extra space and O(2*n)
// func singleNumber(nums []int) int {
// 	m := make(map[int]int, 10)

// 	for _, n := range nums {
// 		m[n]++
// 	}

// 	for k, v := range m {
// 		if v == 1 {
// 			return k
// 		}
// 	}
// 	return 0
// }

// O(n)
func singleNumber(nums []int) int {
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}

	return res
}

func main() {
	a := []int{2, 2, 1}
	res := singleNumber(a)
	fmt.Println(res)
}
