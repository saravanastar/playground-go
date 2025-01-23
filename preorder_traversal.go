package main

import (
	"github.com/saravanastar/playground-go/util"
)

func preorderTraversal(root *util.TreeNode) []int {

	if root == nil {
		return []int{}
	}
	resultNode := []int{}
	resultNode = append(resultNode, root.Val)

	if root.Left != nil {
		leftResult := preorderTraversal(root.Left)
		resultNode = append(resultNode, leftResult...)
	}

	if root.Right != nil {
		rightResult := preorderTraversal(root.Right)
		resultNode = append(resultNode, rightResult...)
	}

	return resultNode
}

// func main() {
// 	root := util.TreeNode{Val: 1}
// 	root.Right = &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 3}}
// 	for _, val := range preorderTraversal(&root) {
// 		fmt.Println(val)
// 	}
// }
