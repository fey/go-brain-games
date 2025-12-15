package gcd

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/fey/go-brain-games/internal/cli"
	"github.com/fey/go-brain-games/internal/game"
)

const DESCRIPTION = "What is the result of the expression?"

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
	first := 1 + rand.Intn(50)
	second := 1 + rand.Intn(50)

	question := fmt.Sprintf("%d %d", first, second)

	answer := gcd(first, second)

	return game.Round{
		Question: question,
		Answer:	strconv.Itoa(answer),
	}
}

func gcd(first, second int) (int) {
	max := max(first, second)
	min := min(first, second)

	rest := max % min

	if (rest == 0) {
		return min
	}

	return gcd(min, rest)
}
