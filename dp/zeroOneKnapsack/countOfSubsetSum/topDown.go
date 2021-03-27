package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 6, 8, 10}
	sum := 10
	fmt.Println(topDown(a, sum))
}

func topDown(arr []int, sum int) int {
	// init
	n := len(arr)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
		dp[i][0] = 1
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}
