package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func Divide(num int, div int) (int, error) {
	if div == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return num / div, nil
}

func main() {
	divideBy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, div := range divideBy {
		res, err := Divide(100, div)
		if err != nil {
			fmt.Printf("100 by %d error: %s\n", div, err)
			continue
		}
		fmt.Printf("100 divide by %d = %d\n", div, res)
	}
}
