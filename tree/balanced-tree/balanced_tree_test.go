package balancedtree

import (
	"testing"

	"github.com/saravanastar/playground-go/util"
)

func TestBalancedTree(t *testing.T) {
	rootNode := &util.TreeNode{
		Val: 1,
		Left: &util.TreeNode{
			Val: 2,
			Left: &util.TreeNode{
				Val: 4,
			},
			Right: &util.TreeNode{
				Val: 5,
			},
		},
		Right: &util.TreeNode{
			Val: 3,
			Left: &util.TreeNode{
				Val: 6,
			},
			Right: &util.TreeNode{
				Val: 7,
			},
		},
	}

	result := isBalanced(rootNode)
	if !result {
		t.Errorf("Expected the tree to be balanced, but got false")
	} else {
		t.Logf("The tree is balanced: %v", result)
	}
}

func TestUnbalancedTree(t *testing.T) {
	rootNode := &util.TreeNode{
		Val: 1,
		Left: &util.TreeNode{
			Val: 2,
			Left: &util.TreeNode{
				Val: 3,
				Left: &util.TreeNode{
					Val: 4,
				},
			},
			Right: &util.TreeNode{
				Val: 5,
			},
		},
		Right: &util.TreeNode{
			Val: 6,
		},
	}

	result := isBalanced(rootNode)
	if result {
		t.Errorf("Expected the tree to not be balanced, but got true")
	} else {
		t.Logf("The tree is correctly identified as unbalanced: %v", result)
	}
}
