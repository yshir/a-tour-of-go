package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("%d\n", add(2, 5))
}
