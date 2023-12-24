package main

import "fmt"

func main() {
	var word = "hello"
	var word = "world" //redeclaring the variable in the same scope. Will give us compilation error.
	word = "world"     // can do assignment to already declared variable.
	fmt.Println(word)
}

/*
 go run 3.go
# command-line-arguments
./3.go:7:6: word redeclared in this block
        ./3.go:6:6: other declaration of word
*/
