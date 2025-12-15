package prime

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/fey/go-brain-games/internal/cli"
	"github.com/fey/go-brain-games/internal/game"
)

const description = "Answer \"yes\" if given number is prime. Otherwise answer \"no\"."

func Play() {
	game := game.Game{
		Description: description,
		BuildRound:  buildRound,
	}

	cli.Run(game)
}

func buildRound() game.Round {
	var answer string
	number := 1 + rand.Intn(30)

	if isPrime(number) {
		answer = "yes"
	} else {
		answer = "no"
	}

	return game.Round{
		Question: strconv.Itoa(number),
		Answer:   answer,
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
