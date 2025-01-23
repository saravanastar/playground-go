package perfectsquare

import "testing"

func TestPerfectSquare(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{12, 3},
		{13, 2},
	}

	for _, test := range tests {
		ret := numSquares(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}

func TestPerfectSquareRecursive(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{12, 3},
		{13, 2},
	}
	for _, test := range tests {
		ret := numSquaresRecursive(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}
