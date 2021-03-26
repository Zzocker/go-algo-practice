package main

func topDown(arr []int, sum int) bool {
	// subprovlem initialization
	n := len(arr)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true // empty subset
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] <= j {
				dp[i][j] = (dp[i-1][j-arr[i-1]] || dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}
