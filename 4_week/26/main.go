package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	r1 := []rune(text1)
	r2 := []rune(text2)
	lcs := make([][]int, len(r1)+1)
	lcs[0] = make([]int, len(r2)+1)

	for i := 1; i <= len(r1); i++ {
		lcs[i] = make([]int, len(r2)+1)

		for j := 1; j <= len(r2); j++ {
			fmt.Println(string(r1[i-1]), string(r2[j-1]))
			if r1[i-1] == r2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	return lcs[len(r1)][len(r2)]
}

func main() {
	text1 := "abcde"
	text2 := "ace"

	fmt.Println(longestCommonSubsequence(text1, text2)) // 3
}
