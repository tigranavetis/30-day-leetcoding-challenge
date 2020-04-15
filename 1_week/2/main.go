package main

import (
	"fmt"
	"time"
)

// isHappy determines if a number is "happy".
// A happy number is a number defined by the following process:
// Starting with any positive integer, replace the number by the sum
// of the squares of its digits, and repeat the process until the number
// equals 1 (where it will stay), or it loops endlessly in a cycle
// which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
func isHappy(n int) bool {
	now := time.Now()
	for {
		sum := 0
		for n > 0 {
			last := n % 10
			sum += last * last
			n /= 10
		}

		if sum == 1 {
			return true
		}

		if time.Since(now) > 1e+5 {
			return false
		}

		n = sum
	}

	return false
}

func main() {
	fmt.Println(isHappy(4343))
}
