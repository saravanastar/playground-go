package longestsubstring

func lengthOfLongestSubstring(s string) int {
	counter := 0

	lookup := make(map[byte]int)
	left := 0
	right := 0
	size := len(s)

	for right < size {
		lookup[s[right]]++
		for left < right && (right-left+1) != len(lookup) {
			lookup[s[left]]--
			if lookup[s[left]] == 0 {
				delete(lookup, s[left])
			}
			left++
		}
		counter = max(counter, right-left+1)
		right++
	}

	return counter

}
