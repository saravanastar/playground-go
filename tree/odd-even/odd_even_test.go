package oddeven

import "testing"

func TestOddEvenTree(t *testing.T) {
	rootNode := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := isEvenOddTree(rootNode)
	if !result {
		t.Errorf("Expected the tree to be an odd-even tree, but got false")
	} else {
		t.Logf("The tree is an odd-even tree: %v", result)
	}
}

func TestOddEvenTreeFalse(t *testing.T) {
	rootNode := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 9, // Changed to an even value to make it not an odd-even tree
			Left: &TreeNode{
				Val: 12,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	result := isEvenOddTree(rootNode)
	if result {
		t.Errorf("Expected the tree to not be an odd-even tree, but got true")
	} else {
		t.Logf("The tree is correctly identified as not an odd-even tree: %v", result)
	}
}
