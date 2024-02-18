package main

import (
	"bufio"
	"fmt"
	"os"
)

func charInput() rune {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a Hangman character guess: ")
		char, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if !(char >= 'a' && char <= 'z') {
			fmt.Printf("That's not a char bro - try again\n\n")
			reader.ReadString('\n') // clears input buffer
		} else {
			return char
		}
	}
}

func main() {
	secretPhrase := "Thanks for playing Go Hangman"

	fmt.Println("Welcome to Go Hangman!")
	// fmt.Println(secretPhrase, utf8.RuneCountInString(secretPhrase))

	for _, char := range secretPhrase {
		if char == 32 {
			fmt.Print(" ")
		} else {
			fmt.Print("_")
		}
	}
	fmt.Printf("\n\n")

	charInput()
}
