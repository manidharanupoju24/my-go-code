package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //if even, continue with next iteration
		}
		fmt.Println("odd number: ", i)
	}
}
