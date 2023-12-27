package main

import "fmt"

func main() {
	x := 1
	switch x {
	case 3:
		fmt.Println("x is 3")
	case 4, 5:
		fmt.Println("x is either 4 or 5")
	default:
		fmt.Println("x is unknown")
	}

	// switch can have init statement, similar to if:
	switch x := someFunc(); x {
	case 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("x is other than 3")
	}

	// we can eliminate match so that each case statement isn't an exact match, but a true or false evaluation
	y := -12
	switch {
	case y > 0:
		fmt.Println("y is greater than 0")
	case y < 0:
		fmt.Println("y is less than 0")
	default:
		fmt.Println("y must be 0")
	}
}

func someFunc() int {
	return 3
}
