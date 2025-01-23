package levelorder

import "github.com/saravanastar/playground-go/util"

func levelOrder(root *util.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*util.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		subResult := []int{}
		for index := 0; index < size; index++ {
			temp := queue[0]
			queue = queue[1:]
			subResult = append(subResult, temp.Val)
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		result = append(result, subResult)

	}
	return result
}
