package main

import (
	"fmt"
	"math/rand"
)

func PrintRandomInt() {
	randInt := rand.Int()
	fmt.Printf("My favorite number is %d\n", randInt)
}
