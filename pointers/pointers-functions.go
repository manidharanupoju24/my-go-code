package main

import "fmt"

func changeValue(word *string) {
	*word += " world"
}

func main() {
	say := "hello"
	fmt.Println(say)
	changeValue(&say)
	fmt.Println(say)
}
