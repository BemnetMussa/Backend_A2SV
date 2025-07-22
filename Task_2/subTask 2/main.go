package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
	"strings"
)

/*
Task : Palindrome Check ✅ak
Write a Go function that takes a string as input and checks
whether it is a palindrome or not.
A palindrome is a word, phrase, number, or other sequence of characters that reads
the same forward and backward (ignoring spaces, punctuation,
and capitalization).
[Optional]: Write test for your function ✅

*/

func PalindromeCheck(word string) bool {
	word = strings.ToLower(word)
	var cleaned []rune

	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

	left := 0
	right := len(cleaned) - 1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}
	return true
}


func main() {
	fmt.Println("Check if palindrome: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.ToLower(input) // to lower case

	if PalindromeCheck(input) {
		fmt.Printf("True, '%s' is a Palindrome\n", input)
	} else {
		fmt.Printf("False, '%s' is not a Palindrome\n", input)
	}
}