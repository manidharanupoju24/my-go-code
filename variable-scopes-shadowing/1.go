package main

import "fmt"

var word = "hello" // this variable is declared at package level

func main() {
	var newWord string = "hello world" // defined at function level (main in this case)
	fmt.Println(word)
	fmt.Println(newWord)
}
