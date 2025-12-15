package calc

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/fey/go-brain-games/internal/cli"
	"github.com/fey/go-brain-games/internal/game"
)

const DESCRIPTION = "What is the result of the expression?"

var operations = [...]rune{'+', '-', '*'}

func Play() {
	game := game.Game{
		Description: DESCRIPTION,
		Rounds:      make([]game.Round, game.ROUNDS_COUNT),
	}

	for i := range game.Rounds {
		game.Rounds[i] = buildRound()
	}

	cli.Play(game)
}

func buildRound() game.Round {
	first := 1 + rand.Intn(10)
	second := 1 + rand.Intn(10)

	op := operations[rand.Intn(len(operations))]

	question := fmt.Sprintf("%d %s %d", first, string(op), second)

	answer := calc(op, first, second)

	return game.Round{
		Question: question,
		Answer:   strconv.Itoa(answer),
	}
}

func calc(op rune, first, second int) int {
	switch op {
	case '+':
		return first + second
	case '-':
		return first - second
	case '*':
		return first * second
	default:
		panic("unknown operator")
	}
}
