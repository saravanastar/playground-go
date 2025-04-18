package mincostchairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost   []int
		output int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	for _, test := range tests {
		ret := minCostClimbingStairs(test.cost)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.cost, test.output)
		}
	}
}
