package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Brain Games!")
	name := prompt("May i have your name? ")

	fmt.Printf("Hello, %s\n!", name)
}

func prompt(query string) string {
	fmt.Print(query)
	var answer string
	fmt.Scanln(&answer)

	return answer
}
