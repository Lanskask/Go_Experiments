package main

import (
	"fmt" 
	"math"
	"math/rand"
)

func foo() {
	fmt.Println(rand.Int63())
	fmt.Println(rand.Intn(100))
}

func main() {
	fmt.Println("HEllo. math.Sqrt(4) is ", math.Sqrt(4))

	foo()
}