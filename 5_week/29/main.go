package main

import (
	"fmt"
	"math"
)

// For this problem, a path is defined as any sequence of nodes from some
// starting node to any node in the tree along the parent-child connections.
// The path must contain at least one node and does not need to go through the root.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Sum int

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

// maxSum computes the max path sum.
func maxSum(root *TreeNode, sum *Sum) int {
	if root == nil {
		return 0
	}

	left := max(maxSum(root.Left, sum), 0)
	right := max(maxSum(root.Right, sum), 0)

	maximSum := left + right + root.Val
	*sum = Sum(max(maximSum, int(*sum)))

	return root.Val + max(left, right)
}

// maxPathSum finds the maximum path sum.
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := Sum(math.Inf(-1))
	maxSum(root, &sum)

	return int(sum)
}

func main() {
	root := TreeNode{1, nil, nil}
	left1 := TreeNode{2, nil, nil}
	right1 := TreeNode{3, nil, nil}
	root.Left = &left1
	root.Right = &right1

	fmt.Println(maxPathSum(&root)) // 6
}
