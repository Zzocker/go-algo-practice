package main

func memoization(weight, values []int, w, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return memoizationInternal(dp, weight, values, w, n)
}

// dp = [n+1][m+1]
// -1 : default
func memoizationInternal(dp [][]int, weights, values []int, w, n int) int {
	if n == 0 || w == 0 {
		return 0
	}
	if dp[n][w] != -1 {
		return dp[n][w]
	}

	if weights[n-1] <= w {
		dp[n][w] = max(values[n-1]+memoizationInternal(dp, weights, values, w-weights[n-1], n-1), memoizationInternal(dp, weights, values, w, n-1))
		return dp[n][w]
	}
	dp[n][w] = memoizationInternal(dp, weights, values, w, n-1)
	return dp[n][w]
}
