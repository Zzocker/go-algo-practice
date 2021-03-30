package minNumCoin

const (
	INT_MAX int = (1 << 32) - 1
)
// https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/

func recursive(c []int, n int, i int) int {
	if n == 0 {
		return 0
	}
	var res int = INT_MAX
	for j := 0; j < i; j++ {
		if c[j] <= n {
			subRes := recursive(c, n-c[j], i)
			if subRes != INT_MAX && subRes+1 < res {
				res = subRes + 1
			}
		}
	}
	return res
}

func dynamic(c []int, v int) int {
	state := make([]int, v+1)
	state[0] = 0
	for i := 1; i < v+1; i++ {
		state[i] = INT_MAX
	}
	for i := 1; i < v+1; i++ {
		for j := range c {
			if c[j] <= i {
				sub := state[i-c[j]]
				if sub != INT_MAX && sub+1 < state[i] {
					state[i] = sub + 1
				}
			}
		}
	}
	if state[v] == INT_MAX {
		return -1
	}
	return state[v]
}

func min(n1, n2 int) int {
	if n2 < n1 {
		n1 = n2
	}
	return n1
}
