package main

import "fmt"

func doAppend(sl []int) {
	sl = append(sl, 100)
	fmt.Println("inside: ", sl) // inside: [1 2 3 100]
}

func main() {
	x := []int{1, 2, 3}
	doAppend(x)                 // uses same underlying array, but sends a copy of x, as doAppend is a function call
	fmt.Println("outside: ", x) // outside: [1 2 3]
}
