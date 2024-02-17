package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	secretPhrase := "Thanks for playing Go Hangman"

	fmt.Println("Welcome to Go Hangman!")
	fmt.Println(secretPhrase, utf8.RuneCountInString(secretPhrase))

	for _, char := range secretPhrase {
		if char == 32 {
			fmt.Print(" ")
		} else {
			fmt.Print("_")
		}
	}
}
