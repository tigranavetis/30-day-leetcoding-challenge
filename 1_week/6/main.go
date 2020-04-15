package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// groupAnagrams groups anagrams together.
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, 0)

	for i, v := range strs {
		v = SortString(v)

		val, exist := m[v]
		if !exist {
			m[v] = append(m[v], strs[i])
		} else {
			val = append(val, strs[i])
			m[v] = val
		}
	}

	result := make([][]string, 0)
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func main() {
	ex := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(ex))
}
