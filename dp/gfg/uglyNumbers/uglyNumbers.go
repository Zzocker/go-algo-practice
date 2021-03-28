package uglyNumbers

/**
	https://www.geeksforgeeks.org/ugly-numbers/
	Ugly numbers are numbers whose only prime factors are 2, 3 or 5.
	The sequence 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, … shows the first 11 ugly numbers. By convention, 1 is included.

	Given a number n, the task is to find n'th Ugly number.

		Examples:

			Input  : n = 7
			Output : 8

			Input  : n = 10
			Output : 12

			Input  : n = 15
			Output : 24

			Input  : n = 150
			Output : 5832
**/

// state : n
// relation : 1+state[n-1]; if n has only 2 , 3 or 5 prime factors
func simple(n int) int {
	var count int
	var i int
	for i = 1; count != n; i++ {
		if isUgly(i) {
			count++
		}
	}
	return i - 1
}

func tabulated(n int) int {
	state := make([]int, n+1)
	state[0] = 1
	var i2, i3, i5 int
	var next_2, next_3, next_5 int = 2, 3, 5
	for i := 1; i < n; i++ {
		state[i] = min(next_2, next_3, next_5)
		if state[i] == next_2 {
			i2++
			next_2 = state[i2] * 2
		}
		if state[i] == next_3 {
			i3++
			next_3 = state[i3] * 3
		}
		if state[i] == next_5 {
			i5++
			next_5 = state[i5] * 5
		}
	}
	return state[n-1]
}

func min(n1, n2, n3 int) int {
	if n2 < n3 {
		n3 = n2
	}
	if n3 < n1 {
		n1 = n3
	}
	return n1
}

// internal
func maxDivide(n1, n2 int) int {
	for n1%n2 == 0 {
		n1 /= n2
	}
	return n1
}

func isUgly(n int) bool {
	n = maxDivide(n, 2)
	n = maxDivide(n, 3)
	n = maxDivide(n, 5)
	if n != 1 {
		return false
	}
	return true
}
