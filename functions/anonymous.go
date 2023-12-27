package main

import "fmt"

// anonymous functions are functions without any name
func main() {
	concat := func(word1, word2 string) string {
		return word1 + word2
	}("hello", " world")

	sum := func(add1, add2 int) int {
		return add1 + add2
	}(2, 3)

	fmt.Println(concat)
	fmt.Println(sum)
}
