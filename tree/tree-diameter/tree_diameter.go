package treediameter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDiameter := 0

	var findDiameter func(node *TreeNode) int
	findDiameter = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := findDiameter(node.Left)
		rightDepth := findDiameter(node.Right)

		// Update the maximum diameter found so far
		if leftDepth+rightDepth+1 > maxDiameter {
			maxDiameter = leftDepth + rightDepth + 1
		}

		// Return the depth of the current node
		return 1 + max(leftDepth, rightDepth)
	}

	findDiameter(root)
	return maxDiameter
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
