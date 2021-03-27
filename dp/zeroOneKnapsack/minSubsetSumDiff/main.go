package main

import "fmt"

func main() {
	a := []int{1, 6, 11, 9}
	fmt.Println(recursion(a))
	fmt.Println(memoization(a))
	fmt.Println(lessSpace(a))
}

func lessSpace(arr []int) int { // time - n*sum*sum
	var sum int					// space - n*sum
	for _, v := range arr {
		sum += v
	}
	// valid sum candidates
	min := sum
	for i := 0; i <= sum/2; i++ {
		check := subsetSum(arr, i)
		if check && abs(sum-2*i) < min {
			min = abs(sum - 2*i)
		}
	}
	return min
}

// check if subset of given sum exists or not
func subsetSum(arr []int, sum int) bool {
	dp := make([][]bool, len(arr)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true // sum =0 => we can select empty subset
	}

	for i := 1; i < len(arr)+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j-arr[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(arr)][sum]
}

////

func memoization(a []int) int { // O(n*sum*sum) space
	var sum1 int // O(n*sum*sum) time
	for _, v := range a {
		sum1 += v
	}
	n := len(a)
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = matrix(sum1 + 1)
	}
	return memoizationInternal(dp, a, n, sum1, 0)
}

func matrix(size int) [][]int {
	out := make([][]int, size)
	for i := range out {
		out[i] = make([]int, size)
		for j := 0; j < size; j++ {
			out[i][j] = -1
		}
	}
	return out
}

func memoizationInternal(dp [][][]int, arr []int, n, sum1, sum2 int) int {
	if n == 0 {
		return abs(sum1 - sum2)
	}
	if dp[n][sum1][sum2] != -1 {
		return dp[n][sum1][sum2]
	}
	if arr[n-1] <= sum1 {
		dp[n][sum1][sum2] = min(memoizationInternal(dp, arr, n-1, sum1-arr[n-1], sum2+arr[n-1]), memoizationInternal(dp, arr, n-1, sum1, sum2))
	} else {
		dp[n][sum1][sum2] = memoizationInternal(dp, arr, n-1, sum1, sum2)
	}
	return dp[n][sum1][sum2]
}

func recursion(a []int) int {
	var sum1 int
	for _, v := range a {
		sum1 += v
	}
	return recursionInternal(a, len(a), sum1, 0)
}

// sum1 : removal from this
// sum2 : add to this
func recursionInternal(arr []int, n, sum1, sum2 int) int {
	if n == 0 {
		return abs(sum1 - sum2)
	}

	if arr[n-1] <= sum1 {
		return min(recursionInternal(arr, n-1, sum1-arr[n-1], sum2+arr[n-1]), recursionInternal(arr, n-1, sum1, sum2))
	}
	return recursionInternal(arr, n-1, sum1, sum2)
}

func abs(n1 int) int {
	if n1 < 0 {
		n1 *= -1
	}
	return n1
}

func min(n1, n2 int) int {
	if n2 < n1 {
		return n2
	}
	return n1
}
