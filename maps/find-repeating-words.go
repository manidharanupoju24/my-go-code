package main

import (
	"fmt"
	"strings"
)

func main() {

	sentence := "The rain poured and poured, making the streets flood."

	inputString := strings.ToLower(sentence)
	fmt.Println(inputString)
	words := strings.Fields(inputString)
	wordCount := make(map[string]int)
	for _, word := range words {
		word = strings.ReplaceAll(word, ",", "")
		_, exists := wordCount[word]
		if exists {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}
	fmt.Println(wordCount)
	var wordsList string
	for word, _ := range wordCount {
		if wordCount[word] > 1 {
			wordsList += "," + word
		}
	}
	fmt.Println(wordsList[1:])
}
