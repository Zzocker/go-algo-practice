package coinChange

// https://www.geeksforgeeks.org/coin-change-dp-7/

// has overlaying problem
func recursive(s []int, n int, i int) int {
	if n == 0 {
		return 1
	} else if n < 0 || (n >= 0 && i <= 0) {
		return 0
	}

	// 1.not-included
	// 2.included
	return recursive(s, n, i-1) + recursive(s, n-s[i-1], i)
}

func dynamic(s []int, n int) int {
	k := len(s)
	state := make([][]int, k+1)
	for i := range state {
		state[i] = make([]int, n+1)
		state[i][0] = 1
	}
	for i := 1; i < k+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] <= j {
				state[i][j] = state[i-1][j] + state[i][j-s[i-1]]
			} else {
				state[i][j] = state[i-1][j]
			}
		}
	}
	return state[k][n]
}
