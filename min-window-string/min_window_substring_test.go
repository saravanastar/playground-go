package minwindowstring

import "testing"

func TestMinWindow(t *testing.T) {
	// Test case 1
	var s = "ADOBECODEBANC"
	var t1 = "ABC"
	var result = minWindow(s, t1)
	var expected = "BANC"
	if result != expected {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected, result)
	}
}
