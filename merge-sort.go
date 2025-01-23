package main

import (
	"fmt"
)

// func main() {
// 	input := []int{9, 8, 7, 6, 5}
// 	mergeSort(input[:], 0, len(input)-1)
// 	fmt.Println(input)
// }

func mergeSort(input []int, left int, right int) {
	fmt.Println(left, right)
	if left < right {
		mid := (right + left) / 2

		mergeSort(input, left, mid)
		mergeSort(input, mid+1, right)
		merge(input, left, right, mid)
	}

}

func merge(input []int, left int, right int, mid int) {
	result := make([]int, right-left+1)

	currentLeft := left
	currentRight := mid + 1
	currentIndex := 0

	for currentLeft <= mid && currentRight <= right {
		if input[currentLeft] < input[currentRight] {
			result[currentIndex] = input[currentLeft]
			currentIndex++
			currentLeft++
		} else if input[currentLeft] == input[currentRight] {
			result[currentIndex] = input[currentLeft]
			currentIndex++
			currentLeft++
			result[currentIndex] = input[currentRight]
			currentIndex++
			currentRight++
		} else {
			result[currentIndex] = input[currentRight]
			currentIndex++
			currentRight++
		}
	}

	if currentLeft > mid && currentRight <= right {
		for index := currentRight; index <= right; index++ {
			result[currentIndex] = input[index]
			currentIndex++
		}

	}

	if currentLeft <= mid && currentRight > right {
		for index := currentLeft; index <= mid; index++ {
			result[currentIndex] = input[index]
			currentIndex++
		}

	}
	fmt.Println(result)
	for index := 0; index < len(result); index++ {
		input[index+left] = result[index]
	}

}
