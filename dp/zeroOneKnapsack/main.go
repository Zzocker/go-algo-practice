package main

import "fmt"

func main() {

	values := []int{60, 100, 120}
	weights := []int{10, 20, 30}
	w := 50

	recu := recursion(weights, values, w, len(weights))
	fmt.Println(recu)

	mem := memoization(weights, values, w, len(weights))
	fmt.Println(mem)

	tDown := topDown(weights, values, w, len(weights))
	fmt.Println(tDown)
}
