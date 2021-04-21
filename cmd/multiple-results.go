package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func RunSwap() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
