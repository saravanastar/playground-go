package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(input []*int) *TreeNode {

	if len(input) == 0 || input[0] == nil {
		return nil
	}
	root := BuildTreeWithIndex(input, 0)
	return root
}

func BuildTreeWithIndex(input []*int, index int) *TreeNode {

	if index >= len(input) || input[index] == nil {
		return nil
	}
	root := TreeNode{Val: *input[index]}
	leftIndex := getLeftChild(index)
	rightIndex := getRightChild(index)
	if leftIndex < len(input) {
		root.Left = BuildTreeWithIndex(input, leftIndex)
	}
	if rightIndex < len(input) {
		root.Right = BuildTreeWithIndex(input, rightIndex)
	}
	return &root
}

// func getParent(index int) int {
// 	return (index - 1) / 2
// }

func getLeftChild(index int) int {
	return (2 * index) + 1
}

func getRightChild(index int) int {
	return (2 * index) + 2
}
