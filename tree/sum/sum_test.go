package sum

import "testing"

func TestSumOfNodes(t *testing.T) {
	rootNode := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := sumOfNodes(rootNode)
	expected := 28 // 1 + 2 + 3 + 4 + 5 + 6 + 7
	if result != expected {
		t.Errorf("Expected sum %d, but got %d", expected, result)
	} else {
		t.Logf("The sum of nodes is correct: %d", result)
	}
}
