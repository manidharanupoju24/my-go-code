package main

import "fmt"

func changeValueAtZeroIndex(array [2]int) {
	array[0] = 3
	fmt.Println("inside: ", array[0])    // will print 3
	fmt.Println("inside array: ", array) // will print array
}
func main() {
	x := [2]int{}
	changeValueAtZeroIndex(x)
	fmt.Println(x)
}
