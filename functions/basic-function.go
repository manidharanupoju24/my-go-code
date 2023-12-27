package main

import "fmt"

func add(x int, y int) int { // func add(x, y int) int {} is also valid
	return x + y
}

func main() {
	result := add(2, 2)
	fmt.Println(result)
}
