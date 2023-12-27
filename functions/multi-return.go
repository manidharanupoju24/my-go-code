package main

import "fmt"

func divide(num, div int) (res, rem int) {
	res = num / div
	rem = res % div
	return res, rem
}
func main() {
	result, remainder := divide(3, 2)
	fmt.Printf("Result: %d, Remainder: %d", result, remainder)
}
