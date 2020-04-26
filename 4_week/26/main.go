package main

import "fmt"

/**
*
* Given two strings text1 and text2, return the length of their longest common subsequence.
*
* A subsequence of a string is a new string generated from the original string with
* some characters(can be none) deleted without changing the relative order of the remaining
* characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.
*
* If there is no common subsequence, return 0.
*
* Example 1:
*	Input: text1 = "abcde", text2 = "ace"
*	Output: 3
*	Explanation: The longest common subsequence is "ace" and its length is 3.
*
* Example 2:
*	Input: text1 = "abc", text2 = "abc"
*	Output: 3
*	Explanation: The longest common subsequence is "abc" and its length is 3.
*
* Example 3:
*	Input: text1 = "abc", text2 = "def"
*	Output: 0
*	Explanation: There is no such common subsequence, so the result is 0.
*
* Constraints:
* 1 <= text1.length <= 1000
* 1 <= text2.length <= 1000
* The input strings consist of lowercase English characters only.
*
**/

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
