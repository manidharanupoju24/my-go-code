package main

import "fmt"

// variadic argument is when we want to provide 0 to infinite arguments.
// notation is ...
func sum(numbers ...int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
