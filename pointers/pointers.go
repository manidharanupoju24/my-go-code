package main

import "fmt"

func main() {
	var x int = 23
	fmt.Println(&x) // prints the address where x stores the value 23
	var intPtr *int
	intPtr = &x
	fmt.Println(intPtr)

	fmt.Println(x) // will print 23
	fmt.Println(*intPtr)
	*intPtr = 80
	fmt.Println(x)
}
