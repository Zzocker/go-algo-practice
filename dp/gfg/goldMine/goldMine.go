package goldMine

// https://www.geeksforgeeks.org/gold-mine-problem/

func recursive(mine [][]int) int {
	var out int
	for i := len(mine) - 1; i >= 0; i-- {
		got := recursiveInternal(mine, i, len(mine[0])-1)
		if got > out {
			out = got
		}
	}
	return out
}

// has overlaying subproblem
func recursiveInternal(mine [][]int, n, m int) int {
	if n < 0 || m < 0 || n >= len(mine) {
		return 0
	}
	// arrival to n,m
	// from left
	fl := recursiveInternal(mine, n, m-1)
	// from left up
	flu := recursiveInternal(mine, n-1, m-1)
	// from left down
	fld := recursiveInternal(mine, n+1, m-1)
	return mine[n][m] + max(fl, flu, fld)

}

// time : O(n*m)
// space : O(n*m)
func dynamic(mine [][]int) int {
	n := len(mine)
	m := len(mine[0])
	state := make([][]int, n+2)
	for i := range state {
		state[i] = make([]int, m+2)
	}
	for j := 1; j < m+1; j++ {
		for i := 1; i < n+1; i++ {
			// came left
			r := state[i][j-1]
			// came from left down
			rd := state[i+1][j-1]
			// came from left up
			rup := state[i-1][j-1]
			state[i][j] = mine[i-1][j-1] + max(r, rd, rup)
		}
	}
	var out int
	for i := 1; i < n+1; i++ {
		if state[i][m] > out {
			out = state[i][m]
		}
	}

	return out
}

func max(n1, n2, n3 int) int {
	if n3 > n2 {
		n2 = n3
	}
	if n2 > n1 {
		n1 = n2
	}
	return n1
}
