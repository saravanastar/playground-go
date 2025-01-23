package minwindowstring

import "math"

func minWindow(s string, t string) string {

	// Create a map to store the frequency of each character in t
	var lookup = make(map[byte]int)

	// Count the frequency of each character in t
	for i := 0; i < len(t); i++ {
		lookup[t[i]]++
	}

	// Initialize the start and end pointers
	start, end := 0, 0

	// Initialize the counter to check whether the substring is valid
	size := len(s)
	counter := len(t)
	resultLen := math.MaxInt32
	resultHead, resultEnd := start, start

	// Loop through the string
	for end < size {
		// If the character is in the lookup map
		// Decrease the counter and the frequency of the character
		// Increase the end pointer
		if _, ok := lookup[s[end]]; ok {
			if lookup[s[end]] > 0 {
				counter--
			}
			lookup[s[end]]--
		}
		end++

		// If the counter is 0, it means the substring is valid
		// Check if the substring is shorter than the previous one
		// If it is, update the result
		// Move the start pointer to the right
		// Increase the counter and the frequency of the character
		// If the start pointer is out of the string, break the loop
		for counter == 0 {
			if end-start < resultLen {
				resultLen = end - start
				resultHead = start
				resultEnd = end
			}
			if start < size {
				if _, ok := lookup[s[start]]; ok {
					if lookup[s[start]] == 0 {
						counter++
					}
					lookup[s[start]]++
				}
				start++
			} else {
				break
			}
		}
	}
	// If the result length is the maximum integer, return an empty string
	// Otherwise, return the substring
	if resultLen == math.MaxInt32 {
		return ""
	}
	return s[resultHead:resultEnd]
}
