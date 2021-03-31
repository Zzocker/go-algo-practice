package friendsPairing

// https://www.geeksforgeeks.org/friends-pairing-problem/

// has overlaying problem
// f(n) = f(n-1) + (n-1)f(n-2)
func recursive(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}
	return recursive(n-1) + (n-1)*recursive(n-2)
}

func dynamic(n int) int {
	state := make([]int, n+1)
	state[0] = 1
	state[1] = 1
	for i := 2; i < n+1; i++ {
		state[i] = state[i-1] + (i-1)*state[i-2]
	}
	return state[n]
}
