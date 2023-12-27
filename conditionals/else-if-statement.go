package main

import "fmt"

func main() {
	x := 0
	if x > 0 {
		fmt.Println("x is greater than 0")
	} else if x < 0 {
		fmt.Println("x is less than 0")
	} else {
		fmt.Println("x is equal to 0")
	}
}
