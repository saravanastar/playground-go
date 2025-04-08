package treediameter

import "testing"

func TestTreeDiameter(t *testing.T) {
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

	result := findDiameter(rootNode)
	expected := 4 // The diameter is the path 4 -> 2 -> 1 -> 3 or 5 -> 2 -> 1 -> 3
	if result != expected {
		t.Errorf("Expected diameter %d, but got %d", expected, result)
	} else {
		t.Logf("Diameter of the binary tree is: %d, test passed", result)
	}
	// Additional test case for a single node tree
	singleNode := &TreeNode{
		Val: 1,
	}
	singleResult := findDiameter(singleNode)
	if singleResult != 0 {
		t.Errorf("Expected diameter for single node tree to be 0, but got %d", singleResult)
	} else {
		t.Logf("Diameter of the single node tree is: %d, test passed", singleResult)
	}
	// Additional test case for an unbalanced tree
	unbalancedTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}
	unbalancedResult := findDiameter(unbalancedTree)
	if unbalancedResult != 5 { // The diameter is the path 5 -> 4 -> 3 -> 2 -> 1 -> 7
		t.Errorf("Expected diameter for unbalanced tree to be 5, but got %d", unbalancedResult)
	} else {
		t.Logf("Diameter of the unbalanced tree is: %d, test passed", unbalancedResult)
	}
}

func TestDiameter(t *testing.T) {
	// This is the entry point for the test suite.
	// If you need to set up any global state or configurations, do it here.
	// For now, we just run the tests.
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	val := findDiameter(root)
	if val != 5 {
		t.Errorf("Diameter of the binary tree should be 5, but got %d", val)
	} else {
		t.Logf("Diameter of the binary tree is: %d, test passed", val)
	}
}
