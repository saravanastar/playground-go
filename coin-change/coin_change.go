package coinchange

import (
	"math"
)

type DP struct {
	Amount int
	Index  int
}

var lookup map[DP]int

func coinChange(coins []int, amount int) int {
	lookup = make(map[DP]int)
	result := coinChangeRecursive(coins, amount, 0)
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func coinChangeRecursive(coins []int, amount int, currentIndex int) int {
	if amount == 0 {
		return 0
	}

	if currentIndex >= len(coins) || amount < 0 {
		return math.MaxInt32
	}
	if val, ok := lookup[DP{Amount: amount, Index: currentIndex}]; ok {
		return val
	}
	includesSelf := 1 + coinChangeRecursive(coins, amount-coins[currentIndex], currentIndex)
	includesNext := 1 + coinChangeRecursive(coins, amount-coins[currentIndex], currentIndex+1)
	excludes := coinChangeRecursive(coins, amount, currentIndex+1)
	includes := min(includesSelf, includesNext)
	result := min(includes, excludes)
	lookup[DP{Amount: amount, Index: currentIndex}] = result
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
