package main

import "fmt"

// getString works correct for english letters.
func getString(s string) string {
	sr := make([]rune, 0)

	for _, v := range s {
		if v != '#' {
			sr = append(sr, v)
			continue
		}

		if len(sr) != 0 {
			sr = sr[:len(sr)-1]
		}
	}

	return string(sr)
}

// Given two strings S and T, returns if they are equal
// when both are typed into empty text editors.
// # means a backspace character.
func backspaceCompare(S string, T string) bool {
	s := getString(S)
	t := getString(T)

	return s == t
}

func main() {
	s1 := "ab#c"
	t1 := "ad#c"
	s2 := "ab##"
	t2 := "c#d#"
	s3 := "a##c"
	t3 := "#a#c"
	s4 := "a#c"
	t4 := "b"

	fmt.Println(backspaceCompare(s1, t1))
	fmt.Println(backspaceCompare(s2, t2))
	fmt.Println(backspaceCompare(s3, t3))
	fmt.Println(backspaceCompare(s4, t4))
}
