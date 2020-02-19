package main

import "fmt"

func main() {
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Infinite loop
	// for {
	// 	fmt.Println("some stuff")
	// }

	// x := 5
	// for {
	// 	fmt.Println("Do stuff", x)
	// 	x += 3
	// 	if x > 25 {
	// 		break
	// 	}
	// }

	// a := 3
	// for x := 5; a < 25; x += 3 {
	// 	fmt.Println("do stuff", x)
	// 	a += 3
	// }
	//
	// for x := range []int{1, 2, 3, 4, 5} {
	// 	fmt.Print(x, ", ")
	// }
	// fmt.Println()

	for index, b := range []int{1, 2, 3, 4, 5} {
		fmt.Print(index, ":", b, ", ")
	}
	fmt.Println()

}
