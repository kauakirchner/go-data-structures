package main

import (
	"fmt"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func InOrder(root *TreeNode) *TreeNode {
	if root != nil {
		InOrder(root.left)
		fmt.Print(root.data, " ")
		InOrder(root.right)
	}
	return root
}

func main() {
	root := &TreeNode{data: 1}
	root.left = &TreeNode{data: 2}
	root.right = &TreeNode{data: 3}
	root.left.left = &TreeNode{data: 4}
	root.left.right = &TreeNode{data: 5}

	fmt.Println("Original tree (sorted):")
	InOrder(root)
	fmt.Println()

	InvertBinaryTree(root)

	fmt.Println("Inverted tree (sorted):")
	InOrder(root)
	fmt.Println()
}
