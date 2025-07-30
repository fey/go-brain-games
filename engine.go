package gobraingames

import (
	"bufio"
	"fmt"
	"os"
)

const ROUNDS_COUNT = 3

type Round struct {
	question string
	answer   string
}

type Game struct {
	Description string
	rounds      []Round
}

func Play(game Game) {
	fmt.Println("Welcome to the Brain Games!")
	name := prompt("May i have your name? ")

	fmt.Printf("Hello, %s\n", name)
	fmt.Println(Game.Description)

	for i := 0; i < ROUNDS_COUNT; i++ {
		round := game.BuildRound()
		answer := round.Answer()
		question := round.Question()

		userAnswer := prompt(question)

		if userAnswer != answer {
			fmt.Printf("'%s' is wrong answer ;(. Correct answer was '%s'.\n", userAnswer, answer)
			fmt.Printf("Let's try again, %s!\n", name)
			return
		}

		fmt.Println("Correct!")

	}

	fmt.Printf("Congratulations, %s!", name)
}

func prompt(query string) string {
	fmt.Print(query)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')

	return answer
}
