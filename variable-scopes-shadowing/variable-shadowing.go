package main

import "fmt"

var word = "hello"

// Variable shadowing occurs when a variable that is within our variable scope, but not in our local scope, is redeclared. This causes the local scope to lose access to the outer scoped variable
func main() {
	var word = "world"
	fmt.Println("Inside main(): ", word)
	printOutter()
}

func printOutter() {
	fmt.Println("The package level word: ", word)
}
