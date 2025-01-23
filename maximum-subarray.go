package main

import (
	"math"
)

// func main() {
// 	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// 	fmt.Println(maxSubArray(input))
// }

/*
*
[-2,1,-3,4,-1,2,1,-5,4]
*
*/
func maxSubArray(nums []int) int {
	sum := math.MinInt
	n := len(nums)

	left := 0
	right := 0
	currentSum := 0

	for right < n {

		currentSum += nums[right]

		for currentSum < 0 {
			currentSum -= nums[left]
			left++
		}

		if left < right {
			sum = maxVal(currentSum, sum)
		}
		right++

	}

	return sum
}

func maxVal(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
