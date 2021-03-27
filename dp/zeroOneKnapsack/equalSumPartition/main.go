package main

import "fmt"

func main() {
	a := []int{1, 5, 11, 5}
	fmt.Println(recursion(a))
	fmt.Println(topDown(a))
}
