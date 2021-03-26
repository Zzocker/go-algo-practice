package main

//
// Base Case
// choice Diagram
//

// Base Case
// Smallest valid input
// --------------------
// 1. If There is no item , then profit will be zero
// 2. If capacity of beg is zero , then also there will be zero profit
//
// => if W==0 || n == 0 {return 0}

// choice diagram
// Item of Weight W1
// - W1 <= W
// 		- included
//		- not included
// - W1 > W
// 		- not included

// n : size of array
func recursion(weights, values []int, w, n int) int {
	if n == 0 || w == 0 {
		return 0
	}

	if weights[n-1] <= w {
		return max(values[n-1]+recursion(weights, values, w-weights[n-1], n-1), recursion(weights, values, w, n-1))
	}
	return recursion(weights, values, w, n-1)
}

func max(n1, n2 int) int {
	if n2 > n1 {
		return n2
	}
	return n1
}
