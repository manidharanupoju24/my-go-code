package main

import (
	"fmt"
	"math"
)

func isArmstrong(num int) string {
	// Complete the function signature and write your code here
	if num < 100 || num > 999 {
		return "InputError"
	}
	rem := 0
	sum := 0.0
	div := num
	for div > 0 {
		rem = div % 10
		sum += math.Pow(float64(rem), 3)
		div /= 10
	}
	fmt.Printf("sum is %d", int(sum))
	fmt.Println()
	fmt.Printf("input number is %d", num)
	fmt.Println()
	if int(sum) == num {
		return "Armstrong"
	} else {
		return "NotArmstrong"
	}

}

func main() {
	returnValue := isArmstrong(407)
	fmt.Println(returnValue)
}
