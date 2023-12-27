package main

import "fmt"

func main() {
	for {
		if err := doSomething(); err != nil {
			break //breaks out of the infinite loop if err is found
		}
	}
	fmt.Println("keep going")
}

func doSomething() int {
	return 0
}
