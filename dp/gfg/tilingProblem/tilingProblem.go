package tilingProblem

// https://www.geeksforgeeks.org/tiling-problem/
// if n < 2 return n
// return c(n) = c(n-1) + c(n-2)

// has overlaying subproblems
func recursive(n int) int {
	if n < 2 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}

func dynamic(n int) int {
	state := make([]int, n+1)
	state[0] = 0
	state[1] = 1
	for i := 2; i < n+1; i++ {
		state[i] = state[i-1] + state[i-2]
	}
	return state[n]
}
