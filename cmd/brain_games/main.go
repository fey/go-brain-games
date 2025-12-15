package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Brain Games!")
	name := prompt("May i have your name? ")

	fmt.Printf("Hello, %s!\n", name)
}

func prompt(query string) string {
	fmt.Print(query)
	var answer string
	if _, err := fmt.Scanln(&answer); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return ""
	}

	return answer
}
