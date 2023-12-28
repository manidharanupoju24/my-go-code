package main

import "fmt"

func main() {
	someSlice := []int{1, 2, 3}

	for index, val := range someSlice {
		fmt.Printf("slice entry %d: %d\n", index, val)
	}

	fmt.Println("----")

	for _, val := range someSlice {
		fmt.Println(val)
	}
}
