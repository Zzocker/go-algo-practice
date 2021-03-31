package subsetSum

// https://www.geeksforgeeks.org/subset-sum-problem-dp-25/

// has overlaying problem
func recursive(set []int, v int, i int) bool {
	if v == 0 { // sum become zero
		return true
	} else if i == 0 { // v!=0 and i == 0 , meas set has been executed
		return false
	}
	// can be included
	if set[i-1] <= v {
		return recursive(set, v-set[i-1], i-1) || recursive(set, v, i-1)
	}
	return recursive(set, v, i-1)
}

//
func dynamic(set []int, v int) bool {
	n := len(set)
	state := make([][]bool, n+1)
	for i := range state {
		state[i] = make([]bool, v+1)
		state[i][0] = true
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < v+1; j++ {
			if set[i-1] <= j {
				state[i][j] = state[i-1][j-set[i-1]] || state[i-1][j]
			} else {
				state[i][j] = state[i-1][j]
			}
		}
	}
	return state[n][v]
}
