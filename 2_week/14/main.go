package main

import "fmt"

/**
* You are given a string s containing lowercase English letters, and a matrix shift, where shift[i] = [direction, amount]:
*
* direction can be 0 (for left shift) or 1 (for right shift).
* amount is the amount by which string s is to be shifted.
* A left shift by 1 means remove the first character of s and append it to the end.
* Similarly, a right shift by 1 means remove the last character of s and add it to the beginning.
* Return the final string after all operations.
*
**/
func stringShift(s string, shift [][]int) string {
	steps := make(map[int]int)
	for _, v := range shift {
		val, exist := steps[v[0]]
		if !exist {
			steps[v[0]] = v[1]
		} else {
			steps[v[0]] = val + v[1]
		}
	}

	leftShifts, _ := steps[0]
	rigthShifts, _ := steps[1]
	shifts := 0

	if leftShifts > rigthShifts {
		shifts = leftShifts - rigthShifts
		shifts = shifts % len(s)
		shifts = len(s) - shifts
	} else {
		shifts = rigthShifts - leftShifts
		shifts = shifts % len(s)
	}

	s = s[len(s)-shifts:len(s)] + s[:len(s)-shifts]

	return s
}

func main() {
	fmt.Println(stringShift("abc", [][]int{[]int{0, 1}, []int{1, 2}}))
}
