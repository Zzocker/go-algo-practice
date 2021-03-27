package main

func recursion(arr []int) bool {
	var sum1 int
	for _, v := range arr {
		sum1 += v
	}
	if sum1&1 == 1 {
		return false
	}
	return recursionInternal(arr, len(arr), sum1, 0)
}

// sum1 : total sum from which a number will be subtracted
// sum2 : number will be added
func recursionInternal(a []int, n, sum1, sum2 int) bool {
	if sum1 == sum2 {
		return true
	} else if n == 0 {
		return false
	}

	if a[n-1] <= sum1 && (recursionInternal(a, n-1, sum1-a[n-1], sum2+a[n-1]) || recursionInternal(a, n-1, sum1, sum2)) {
		return true
	} else if a[n-1] > sum1 && recursionInternal(a, n-1, sum1, sum2) {
		return true
	}
	return false
}
