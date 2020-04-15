package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middle := head
	length := 1

	for head.Next != nil {
		length++
		head = head.Next
	}

	count := 0
	for count != length/2 {
		count++
		middle = middle.Next
	}

	return middle
}

func main() {
	n1 := ListNode{1, nil}
	n2 := ListNode{2, nil}
	n3 := ListNode{3, nil}
	n4 := ListNode{4, nil}
	n5 := ListNode{5, nil}
	n6 := ListNode{6, nil}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6

	val := middleNode(&n1)
	fmt.Println(val)
}
