package fallingpath

import "testing"

func TestFallingPathSum(t *testing.T) {
	tests := []struct {
		grid   [][]int
		output int
	}{
		{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
		{[][]int{{-19, 57}, {-40, -5}}, -59},
	}

	for _, test := range tests {
		ret := minFallingPathSum(test.grid)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.grid, test.output)
		}
	}
}
