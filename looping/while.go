package main

import "fmt"

func main() {
	var i int
	for i < 10 {
		i++
	}

	b := true

	for b {
		fmt.Println("hello") // infinite loop
	}
}
