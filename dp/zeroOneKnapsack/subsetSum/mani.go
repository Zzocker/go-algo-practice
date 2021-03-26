package main

import "fmt"

func main() {
	a := []int{2, 5, 7, 8, 10}
	sum := 11
	fmt.Println(recursion(a, sum))

	ans := topDown(a, sum)
	fmt.Println(ans)
}
