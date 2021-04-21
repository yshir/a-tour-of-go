package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func RunAdd() {
	fmt.Printf("%d\n", add(2, 5))
}
