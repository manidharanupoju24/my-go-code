//go:generate stringer -type=Pill
package main

import "fmt"

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func main() {
	fmt.Println(Aspirin)
	fmt.Println(Ibuprofen)
	fmt.Println(Paracetamol)
	fmt.Println(Acetaminophen)
	fmt.Println(Placebo)
}
