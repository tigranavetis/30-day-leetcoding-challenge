package main

import (
	"fmt"
	"math"
	"sort"
)

/**
* We have a collection of stones, each stone has a positive integer weight.
*
* Each turn, we choose the two heaviest stones and smash them together. Suppose the stones have weights x and y with x <= y.
* The result of this smash is:
*
* If x == y, both stones are totally destroyed;
* If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
* At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)
**/
func lastStoneWeight(stones []int) int {
	lenSt := len(stones)
	if lenSt == 0 {
		return 0
	}

	if lenSt == 1 {
		return stones[0]
	}

	sort.Ints(stones)
	el := smash(stones[lenSt-1], stones[lenSt-2])

	if el > 0 {
		stones = append(stones[:lenSt-2], el)
		return lastStoneWeight(stones)
	}

	stones = stones[:lenSt-2]
	return lastStoneWeight(stones)
}

func smash(st1 int, st2 int) int {
	diff := float64(st1 - st2)
	return int(math.Abs(diff))
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}
