package fallingpath

import "math"

var lookup map[Table]int

type Table struct {
	Row    int
	Column int
}

func minFallingPathSum(matrix [][]int) int {
	lookup = make(map[Table]int)
	result := math.MaxInt
	for index := 0; index < len(matrix); index++ {
		result = min(result, flowDown(matrix, 0, index))
	}
	return result
}

func flowDown(matrix [][]int, row int, column int) int {
	if row >= len(matrix) {
		return math.MaxInt
	}
	if !isBounded(matrix, row, column) {
		return math.MaxInt
	}
	if row == len(matrix)-1 {
		return matrix[row][column]
	}
	if val, ok := lookup[Table{Row: row, Column: column}]; ok {
		return val
	}

	leftDiagnol := flowDown(matrix, row+1, column-1)
	rightDiagnol := flowDown(matrix, row+1, column+1)
	down := flowDown(matrix, row+1, column)
	diagMin := min(leftDiagnol, rightDiagnol)
	result := min(diagMin, down) + matrix[row][column]
	lookup[Table{Row: row, Column: column}] = result
	return result

}

func isBounded(matrix [][]int, row int, column int) bool {
	return (row >= 0 && row < len(matrix)) && (column >= 0 && column < len(matrix[0]))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
