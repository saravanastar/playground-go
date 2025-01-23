package perfectsquare

/**
* Explanation: https://www.youtube.com/watch?v=1xfx6M_GqFk&ab_channel=Techdose
**/
func numSquares(n int) int {

	dp := make([]int, n+1)
	dp[0] = 0
	for target := 1; target <= n; target++ {
		dp[target] = target
		for squareNumber := 1; squareNumber <= target; squareNumber++ {
			squareValue := squareNumber * squareNumber
			if squareValue > target {
				break
			}
			dp[target] = min(dp[target], dp[target-squareValue]+1)
		}
	}

	return dp[n]

}

var lookup = make(map[int]int)

func numSquaresRecursive(n int) int {
	if val, ok := lookup[n]; ok {
		return val
	}
	result := n
	if n <= 3 {
		return n
	}
	for i := 1; i*i <= n; i++ {
		result = min(result, numSquaresRecursive(n-i*i)+1)
	}
	lookup[n] = result
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
