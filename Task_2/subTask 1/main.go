package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

/* Task: 2 description 1
Task:  Word Frequency Count ✅
Write a Go function that takes a string as input and returns
a dictionary containing the frequency of each word in the string.
 Treat words in a case-insensitive manner and ignore punctuation marks.
[Optional]: Write test for your function ✅
*/

// WordFrequency counts the frequency of each word in the input string
func WordFrequency(input string) map[string]int {
	// Remove punctuation and normalize case
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return unicode.ToLower(r) // case-insensitive
	}, input)

	words := strings.Fields(clean)

	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	result := WordFrequency(input)
	for word, count := range result {
		fmt.Printf("'%s': %d\n", word, count)
	}
}
