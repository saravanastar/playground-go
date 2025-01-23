package longestsubstring

import "testing"

func TestLongestSubString(t *testing.T) {
	// Test case 1
	var s = "abcabcbb"
	var result = lengthOfLongestSubstring(s)
	var expected = 3
	if result != expected {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected, result)
	}
}
