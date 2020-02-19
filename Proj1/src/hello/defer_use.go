package main

import "fmt"

func foo2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func foo() {
	defer fmt.Println("Done1")
	defer fmt.Println("Done2")

	fmt.Println("Some stuff here")
}

func main() {
	// foo()
	foo2()
}
