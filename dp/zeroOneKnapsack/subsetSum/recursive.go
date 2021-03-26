package main

func recursion(arr []int, sum int) bool {
	return recursionInternal(arr, sum, len(arr))
}

func recursionInternal(arr []int, sum, n int) bool {
	if sum == 0 { // sum == 0 return true empty subset
		return true
	} else if n == 0 { // return false
		return false
	}

	if arr[n-1] <= sum && (recursionInternal(arr, sum-arr[n-1], n-1) || recursionInternal(arr, sum, n-1)) {
		return true
	} else if arr[n-1] > sum && recursionInternal(arr, sum, n-1) {
		return true
	}
	return false
}
