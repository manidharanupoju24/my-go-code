package main

import "fmt"

func main() {
	modelToMake := map[string]string{
		"prius":    "toyota",
		"chevelle": "chevy",
	}

	carMake := modelToMake["chevelle"]
	fmt.Println(carMake)

	if carMake, ok := modelToMake["outback"]; ok {
		fmt.Printf("car model 'outback' has make %q", carMake)
		fmt.Println()
	} else {
		fmt.Printf("car model \"outback\" has an unknown make")
		fmt.Println()
	}

	// Creating a map with int keys and string values
	myMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	// Accessing an existing key
	value, present := myMap[2] //map by default returns the value for the key presented, and boolean if the value is present for that key
	fmt.Printf("Value: %s, Present: %v\n", value, present)

	// Accessing a non-existing key
	value, present = myMap[4]
	fmt.Printf("Value: %s, Present: %v\n", value, present)
}
