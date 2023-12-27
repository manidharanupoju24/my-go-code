package main

import "fmt"

func main() {
	var i int
	for ; i < 10; i++ { // without initialising i and declaring it before hand
		fmt.Println(i)
	}
	fmt.Println("i's final value: ", i)
}
