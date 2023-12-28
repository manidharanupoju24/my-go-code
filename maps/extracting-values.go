package main

import "fmt"

func main() {
	modelToMake := map[string]string{
		"prius":    "toyota",
		"chevelle": "chevy",
		"outback":  "subaru",
	}

	for key, val := range modelToMake {
		fmt.Printf("car model %q has make %q\n", key, val)
	}
}
