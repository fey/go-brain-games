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
		BuildRound: buildRound,
	}

	cli.Run(game)
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
