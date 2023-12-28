package main

import "fmt"

func main() {
	var x = []int{}
	fmt.Println(x)
	fmt.Println(len(x))

	var y = []int{8, 4, 5, 6}
	fmt.Println(y)
	fmt.Println(len(y))

	y = append(y, 2, 3) // this grows the size of the slice
	fmt.Println("After append: ")
	fmt.Println(y)
	fmt.Println(len(y))

	//limited views of the slice
	z := y[1:3]
	fmt.Println("z:", z)
}
