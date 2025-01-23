package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		output int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1}, 1, 1},
		{[]int{1}, 2, 2},
	}

	for _, test := range tests {
		ret := coinChange(test.coins, test.amount)
		if ret != test.output {
			t.Errorf("Got %d for input %v, %d; expected %d", ret, test.coins, test.amount, test.output)
		}
	}
}
