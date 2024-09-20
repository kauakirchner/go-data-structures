package main

func InvertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.left, root.right = root.right, root.left
	InvertBinaryTree(root.left)
	InvertBinaryTree(root.right)

	return root
}
