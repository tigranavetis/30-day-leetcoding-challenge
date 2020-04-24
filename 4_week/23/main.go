package main

import "fmt"

/**
* Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
*
* Example 1:
*	Input: [5,7]
*	Output: 4
* Example 2:
*	Input: [0,1]
*	Output: 0
*
**/

// O(n) solution
// func rangeBitwiseAnd(m int, n int) int {
// 	res := m
// 	for i := m + 1; i <= n; i++ {
// 		res &= i
// 	}
//
// 	return res
// }

func rangeBitwiseAnd(m int, n int) int {
	res := 0

	for m >= 1 && n >= 1 {
		if m == n {
			return m << res
		}

		m >>= 1
		n >>= 1
		res++
	}

	return 0
}

func main() {
	m1, n1 := 10, 12
	m2, n2 := 0, 1
	fmt.Println(rangeBitwiseAnd(m1, n1))
	fmt.Println(rangeBitwiseAnd(m2, n2))
}
