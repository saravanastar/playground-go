package mincostchairs

import "math"

var lookup map[int]int

func minCostClimbingStairs(cost []int) int {
	lookup = make(map[int]int)
	lookup[0] = cost[0]
	lookup[1] = cost[1]
	return min(minCost(cost, len(cost)-1), minCost(cost, len(cost)-2))
}

func minCost(cost []int, n int) int {
	if n < 0 {
		return math.MaxInt
	}
	if val, ok := lookup[n]; ok {
		return val
	}
	result := min(minCost(cost, n-1), minCost(cost, n-2)) + cost[n]
	lookup[n] = result
	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
