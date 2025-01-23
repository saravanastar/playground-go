package main

import "github.com/saravanastar/playground-go/util"

func inorderTraversal(root *util.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	resultVal := []int{}
	resultVal = append(resultVal, inorderTraversal(root.Left)...)
	resultVal = append(resultVal, root.Val)
	resultVal = append(resultVal, inorderTraversal(root.Right)...)

	return resultVal
}
