package main

import "fmt"

func PrintForContinued() {
	sum := 10

	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}

	fmt.Println("done")
}
