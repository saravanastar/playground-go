package min_depth_bt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct{}

// findDepth calculates the minimum depth of a binary tree
func (s *Solution) findDepth(root *TreeNode) int {
	minimumTreeDepth := 0
	//  TODO: Write your code here
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return s.findDepth(root.Right) + 1
	}
	if root.Right == nil {
		return s.findDepth(root.Left) + 1
	}
	leftDepth := s.findDepth(root.Left)
	rightDepth := s.findDepth(root.Right)
	if leftDepth < rightDepth {
		minimumTreeDepth = leftDepth + 1
	} else {
		minimumTreeDepth = rightDepth + 1
	}
	return minimumTreeDepth
}
