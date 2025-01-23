package trapwater

import "testing"

func TestCase1(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	result := trap(height)
	if result != 6 {
		t.Errorf("Invalid output %v", result)
	}
}
