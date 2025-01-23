package minpathsum

import "math"

type Table struct {
	Row    int
	Column int
}

var lookup map[Table]int

func minPathSum(grid [][]int) int {
	lookup = make(map[Table]int)
	return findMinSumRecursive(grid, len(grid)-1, len(grid[0])-1)
}

func findMinSumRecursive(grid [][]int, row int, column int) int {
	if row == 0 && column == 0 {
		return grid[0][0]
	}
	if (row < 0 || row >= len(grid)) || (column < 0 || column >= len(grid[0])) {
		return math.MaxInt
	}
	key := Table{Row: row, Column: column}
	if val, ok := lookup[key]; ok {
		return val
	}

	if row == 0 {
		result := grid[row][column] + findMinSumRecursive(grid, row, column-1)
		lookup[key] = result
		return result
	}
	if column == 0 {
		result := grid[row][column] + findMinSumRecursive(grid, row-1, column)
		lookup[key] = result
		return result
	}
	result := grid[row][column] + int(math.Min(float64(findMinSumRecursive(grid, row-1, column)), float64(findMinSumRecursive(grid, row, column-1))))
	lookup[key] = result
	return result
}

func minPathSumDP(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, columns)
	}

	dp[0][0] = grid[0][0]

	for row := 1; row < rows; row++ {
		dp[row][0] = dp[row-1][0] + grid[row][0]
	}

	for column := 1; column < columns; column++ {
		dp[0][column] = dp[0][column-1] + grid[0][column]
	}

	for row := 1; row < rows; row++ {
		for column := 1; column < columns; column++ {
			dp[row][column] = grid[row][column] + int(math.Min(float64(dp[row-1][column]), float64(dp[row][column-1])))
		}
	}

	return dp[rows-1][columns-1]
}
