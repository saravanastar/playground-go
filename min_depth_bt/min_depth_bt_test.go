package min_depth_bt

import "testing"

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Test case 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 2,
		},
	}

	solution := Solution{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := solution.findDepth(test.root)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
