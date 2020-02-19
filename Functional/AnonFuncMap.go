package main 

import "fmt"

func main() {
	// message := "Hello from the other side"

	// var arr = []int {1,2,3,4}

	// fmt.Println(arr)

	// func(in int) {
	// 	fmt.Println(in)
	// }()

	mapper := func (i interface{}) interface{} {
		return strings.ToUpper(i.(string))
	}

	fmt.Println(Map(mapper, New("milu", "rantanplan")))
}