package main

import "fmt"

/**
* Given a string containing only three types of characters: '(', ')' and '*', write
* a function to check whether this string is valid. We define the validity of a string by these rules:

* 1 .Any left parenthesis '(' must have a corresponding right parenthesis ')'.
* 2. Any right parenthesis ')' must have a corresponding left parenthesis '('.
* 3. Left parenthesis '(' must go before the corresponding right parenthesis ')'.
* 4. '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
* 5. An empty string is also valid.
* Example 1:
*	Input: "()"
*	Output: True
* Example 2:
*	Input: "(*)"
*	Output: True
* Example 3:
*	Input: "(*))"
*	Output: True
Note:
The string size will be in the range [1, 100].
**/

func checkValidString(s string) bool {
	sr := []rune(s)
	min := 0
	max := 0

	for _, v := range sr {
		ch := string(v)
		if ch == "(" {
			min++
		} else {
			if min > 0 {
				min--
			}
		}

		if ch != ")" {
			max++
		} else {
			max--
			if max < 0 {
				return false
			}
		}
	}

	return min == 0
}

func main() {
	fmt.Println(checkValidString("()(*))"))   // true
	fmt.Println(checkValidString("(())*)()")) // true
	fmt.Println(checkValidString("((())"))    // false
	fmt.Println(checkValidString("(***)"))    // true
}
