package catalanNumber

// https://www.geeksforgeeks.org/program-nth-catalan-number/

// overlaying subproblem exists
func recursive(n int) int {
	if n == 0 {
		return 1
	}
	var out int
	for i := 0; i < n; i++ {
		out += recursive(i) * recursive(n-i-1)
	}
	return out
}

func dynamic(n int) int { // space O(n); and time : O(n*n)
	if n < 2 {
		return 1
	}
	state := make([]int, n+1)
	state[0] = 1
	state[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			state[i] += state[j] * state[i-1-j]
		}
	}
	return state[n]
}
