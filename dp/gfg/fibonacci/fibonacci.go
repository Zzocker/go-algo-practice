// Package fibonacci demonstrate fibonacci's using
// recursive solution
// memoized solution
// tabulated solution
package fibonacci

// overlaying subproblem
func recursive(n int) int {
	if n < 2 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}

// overlaying subprobelm is solved
// state - subproblem are filled on demand
// which means some state might remain empty
func memoized(n int) int {
	state := make([]int, n+1)
	for i := range state {
		state[i] = -1
	}
	return memoizedInternal(state, n)
}
func memoizedInternal(state []int, n int) int {
	if state[n] == -1 {
		if n < 2 {
			state[n] = n
		} else {
			state[n] = memoizedInternal(state, n-1) + memoizedInternal(state, n-2)
		}
	}
	return state[n]
}

// overlaying subproblem is solved
// state start with smallest subprobelm and gradually
// solves more bigger subproblem
func tabulated(n int) int {
	state := make([]int, n+1)
	state[0] = 0
	state[1] = 1
	for i := 2; i < n+1; i++ {
		state[i] = state[i-1] + state[i-2]
	}
	return state[n]
}
