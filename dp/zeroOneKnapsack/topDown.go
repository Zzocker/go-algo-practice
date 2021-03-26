package main

func topDown(weights, values []int, w, n int) int {
	// subproblem matrix internalization
	// dp[n+1][w+1]
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
		}
	}

	// choice diagram
	for i := 1; i < n+1; i++ {
		for j := 1; j < w+1; j++ {
			if weights[i-1] <= j {
				dp[i][j] = max(values[i-1]+dp[i-1][j-weights[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][w]
}
