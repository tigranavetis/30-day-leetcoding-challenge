package main

import (
	"fmt"
	"sort"
)

// countElements counts element x such that x + 1 is also in arr.
func countElements(arr []int) int {
	count := 0
	m := make(map[int]int, 0)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	for _, v := range arr {
		if _, exist := m[v]; !exist {
			m[v] = 0
		}

		if _, exist := m[v+1]; exist {
			m[v+1] = m[v+1] + 1
		}
	}

	for _, v := range m {
		if v > 0 {
			count += v
		}
	}

	return count
}

func main() {
	ex1 := []int{1, 2, 3}
	ex2 := []int{1, 1, 3, 3, 5, 5, 7, 7}
	ex3 := []int{1, 3, 2, 3, 5, 0}
	ex4 := []int{1, 1, 2, 2}
	fmt.Println(countElements(ex1))
	fmt.Println(countElements(ex2))
	fmt.Println(countElements(ex3))
	fmt.Println(countElements(ex4))
}
