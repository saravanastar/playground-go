package balancedtree

import "github.com/saravanastar/playground-go/util"

func isBalanced(root *util.TreeNode) bool {
	return height(root) != -1
}

func height(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	leftNodeHeight := height(root.Left)
	if leftNodeHeight == -1 {
		return -1 // if left subtree is not balanced
	}
	rightNodeHeight := height(root.Right)
	if rightNodeHeight == -1 {
		return -1 // if right subtree is not balanced
	}
	if leftNodeHeight-rightNodeHeight > 1 || rightNodeHeight-leftNodeHeight > 1 {
		return -1 // not balanced
	}
	if leftNodeHeight > rightNodeHeight {
		return leftNodeHeight + 1
	}
	return rightNodeHeight + 1
}
