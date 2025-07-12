package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

/* Task: 2 description 1
Task:  Word Frequency Count ✅
Write a Go function that takes a string as input and returns
a dictionary containing the frequency of each word in the string.
 Treat words in a case-insensitive manner and ignore punctuation marks.
[Optional]: Write test for your function ✅
*/

// char frequency counter func
func Counter(word string) map[rune]int {
	charFrequency := map[rune]int{}
	for _, char := range word {
		if char == '\n' || char == '\r' {
			continue
		}
		if unicode.IsPunct(char) || unicode.IsSpace(char) {
			continue
		}
		charFrequency[char] += 1
	}
	return charFrequency
}

func main() {
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	result := Counter(input)
	for char, value := range result {
		fmt.Printf("'%c' : %d\n", char, value)	
	}
}