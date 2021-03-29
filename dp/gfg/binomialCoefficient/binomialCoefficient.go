package binomialCoefficient

// https://www.geeksforgeeks.org/binomial-coefficient-dp-9/
// find coefficient of x^k in binomial expension of (1+x)^n
// => nCk = (n-1)Ck + (n-1)C(k-1)

// has overlaying subproblem
func recursive(n, k int) int { //
	if k == 0 || k == n {
		return 1
	}
	return recursive(n-1, k-1) + recursive(n-1, k)
}

func memoized(n, k int) int {
	state := make([][]int, n+1)
	for i := range state {
		state[i] = make([]int, k+1)
	}
	return memoizedInternal(state, n, k)
}

// top - down approach
// state : [n+1][k+1]
func memoizedInternal(state [][]int, n, k int) int { // time : O(n*k) , space : O(n*k)
	if state[n][k] == 0 { // no value will be zero
		if k == 0 || k == n {
			state[n][k] = 1
		} else {
			state[n][k] = memoizedInternal(state, n-1, k-1) + memoizedInternal(state, n-1, k)
		}
	}
	return state[n][k]
}

// bottom - up approach
// time : O(n*k)
// space : O(n*k)
func dynamic(n, k int) int {
	state := make([][]int, n+1)
	for i := range state {
		state[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			if j == 0 || j == i {
				state[i][j] = 1
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < k+1; j++ {
			state[i][j] = state[i-1][j] + state[i-1][j-1]
		}
	}
	return state[n][k]
}
