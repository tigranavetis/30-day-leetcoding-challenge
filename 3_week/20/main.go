package main

import "fmt"

/**
* Return the root node of a binary search tree that matches the given preorder traversal.
*
* (Recall that a binary search tree is a binary tree where for every node, any descendant of
* node.left has a value < node.val, and any descendant of node.right has a value > node.val.
* Also recall that a preorder traversal displays the value of the node first, then traverses node.left,
* then traverses node.right.)
*
* Example 1:
*	Input: [8,5,1,7,10,12]
*	Output: [8,5,10,1,7,null,12]
*
* Note:
* 	1. 1 <= preorder.length <= 100
* 	2. The values of preorder are distinct.
*
*
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, start int, par *TreeNode) (*TreeNode, int) {
	if len(preorder) <= start {
		return nil, start
	}

	if par != nil && par.Val < preorder[start] {
		return nil, start
	}

	root := &TreeNode{Val: preorder[start]}
	leftNode, index := buildTree(preorder, start+1, root)
	root.Left = leftNode
	rightNode, nextIndex := buildTree(preorder, index, par)
	root.Right = rightNode

	return root, nextIndex
}

func bstFromPreorder(preorder []int) *TreeNode {
	root, _ := buildTree(preorder, 0, nil)
	return root
}

func main() {
	inp1 := []int{8, 5, 1, 7, 10, 12}
	fmt.Println(bstFromPreorder(inp1))
}
