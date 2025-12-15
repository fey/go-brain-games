package cli

import (
	"fmt"

	"github.com/fey/go-brain-games/internal/game"
)

func Run(game game.Game) {
	fmt.Println("Welcome to the Brain Games!")
	name := prompt("May i have your name? ")

	fmt.Printf("Hello, %s\n", name)
	fmt.Println(game.Description)

	for range game.RoundsCount {
		round := game.BuildRound()

		question := round.Question
		answer := round.Answer

		fmt.Println("Question:", question)
		userAnswer := prompt("Your answer: ")

		if userAnswer != answer {
			fmt.Printf("'%s' is wrong answer ;(. Correct answer was '%s'\n", userAnswer, answer)
			fmt.Printf("Let's try again, %s!\n", name)
			return
		}

		fmt.Println("Correct!")
	}

	fmt.Printf("Congratulations, %s!\n", name)
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
