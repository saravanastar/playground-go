package trapwater

func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	for index := 1; index < len(height); index++ {
		leftMax[index] = Max(leftMax[index-1], height[index])
	}
	size := len(height)
	rightMax[size-1] = height[size-1]
	for index := len(height) - 2; index >= 0; index-- {
		rightMax[index] = Max(rightMax[index+1], height[index])
	}
	total := 0
	for index := 0; index < len(height); index++ {
		min := Min(leftMax[index], rightMax[index])
		total = total + (min - height[index])
	}
	return total
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
