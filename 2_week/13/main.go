package main

import "fmt"

// Given a binary array, finds the maximum length of a contiguous subarray with equal number of 0 and 1.
func findMaxLength(nums []int) int {
	maxLength := 0
	count := 0
	m := make(map[int]int, 0)
	m[0] = -1

	for i, v := range nums {
		if v == 1 {
			count++
		} else {
			count--
		}

		if val, exist := m[count]; exist {
			if maxLength < i-val {
				maxLength = i - val
			}
		} else {
			m[count] = i
		}
	}

	return maxLength
}

func main() {
	in1 := []int{0, 1, 0, 1, 1, 1, 0}
	fmt.Println(findMaxLength(in1))
	in2 := []int{1, 0, 1, 0, 1, 1, 1}
	fmt.Println(findMaxLength(in2))
}
