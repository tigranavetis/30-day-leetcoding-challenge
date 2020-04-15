package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Diameter int

// maxDepth computes the max height.
func maxDepth(root *TreeNode, diameter *Diameter) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left, diameter)
	right := maxDepth(root.Right, diameter)

	if int(*diameter) < left+right {
		*diameter = Diameter(left + right)
	}

	if left < right {
		return right + 1
	}

	return left + 1
}

// diameterOfBinaryTree computes the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path
// between any two nodes in a tree.
// This path may or may not pass through the root.
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var diameter Diameter = 0
	maxDepth(root, &diameter)
	return int(diameter)
}

func main() {
	root := TreeNode{1, nil, nil}
	left1 := TreeNode{2, nil, nil}
	right1 := TreeNode{3, nil, nil}
	left2 := TreeNode{4, nil, nil}
	right2 := TreeNode{5, nil, nil}
	root.Left = &left1
	root.Right = &right1
	left1.Left = &left2
	left1.Right = &right2

	fmt.Println(diameterOfBinaryTree(&root))
}
