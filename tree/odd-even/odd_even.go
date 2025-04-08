package oddeven

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	// ToDo: Write Your Code Here.
	if root == nil {
		return false
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		isOdd := level%2 == 1
		var prevNode *TreeNode
		for index := 0; index < size; index++ {
			currentNode := queue[0]
			queue = queue[1:]
			if isOdd {
				if currentNode.Val%2 != 0 {
					return false
				}
				if prevNode != nil && prevNode.Val < currentNode.Val {
					return false
				}
			} else {
				if currentNode.Val%2 != 1 {
					return false
				}
				if prevNode != nil && prevNode.Val > currentNode.Val {
					return false
				}
			}
			prevNode = currentNode
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		level += 1
	}
	return true
}
