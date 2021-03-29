package permutationCoefficient

// https://www.geeksforgeeks.org/permutation-coefficient/
// nPk = (n-1)Pk + k*((n-1)P(k-1))

// has overlaying problem
func recursive(n, k int) int {
	if n < k {
		return 0
	} else if k == 0 {
		return 1
	}
	return recursive(n-1, k) + k*recursive(n-1, k-1)
}

// bottom-up approach
// space : O(n*k)
// time : O(n*k)
func dynamic(n, k int) int {
	state := make([][]int, n+1)
	for i := range state {
		state[i] = make([]int, k+1)
		state[i][0] = 1
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < min(i, k)+1; j++ {
			state[i][j] = state[i-1][j] + j*state[i-1][j-1]
		}
	}
	return state[n][k]
}

func min(n1, n2 int) int {
	if n2 < n1 {
		return n2
	}
	return n1
}
