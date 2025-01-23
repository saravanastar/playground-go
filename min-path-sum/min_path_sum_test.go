package minpathsum

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	result := minPathSum(grid)
	if result != 7 {
		t.Errorf("Expected 7 but got %v", result)
	}
}

func TestMinPathSum2(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	result := minPathSum(grid)
	if result != 12 {
		t.Errorf("Expected 12 but got %v", result)
	}
}
