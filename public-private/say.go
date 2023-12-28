package say

import "fmt"

func PrintHello() {
	// this can be referred outside of the package
	// starts with capital letter
	fmt.Println("Hello")
}

func printWorld() {
	//this can't be referred outside of the package
	//starts with small letter
	fmt.Println("World")
}

func PrintHelloWorld() {
	PrintHello()
	printWorld()
}
