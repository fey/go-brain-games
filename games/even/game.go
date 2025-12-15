package even

import (
	"math/rand"
	"strconv"

	"github.com/fey/go-brain-games/internal/cli"
	"github.com/fey/go-brain-games/internal/game"
)

const DESCRIPTION = "Answer \"yes\" if the number is even, otherwise answer \"no\"."

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
	question := rand.Intn(100)

	var answer string

	if question%2 == 0 {
		answer = "yes"
	} else {
		answer = "no"
	}

	return game.Round{
		Question: strconv.Itoa(question),
		Answer:   answer,
	}
}
