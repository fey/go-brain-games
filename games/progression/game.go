package progression

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/fey/go-brain-games/internal/cli"
	"github.com/fey/go-brain-games/internal/game"
)

const DESCRIPTION = "What number is missing in the progression?"

func Play() {
	game := game.Game{
		Description: DESCRIPTION,
		BuildRound: buildRound,
	}

	cli.Run(game)
}

func buildRound() game.Round {
	first := 1 + rand.Intn(10)
	step := 1 + rand.Intn(10)
	count := 10

	hiddenIndex := rand.Intn(count)
	progression := progression(first, step, count)
	answer := progression[hiddenIndex]

	question := buildQuestion(progression, hiddenIndex)

	return game.Round{
		Question: question,
		Answer:   strconv.Itoa(answer),
	}
}

func progression(first, step, length int) []int {
	s := make([]int, length)

	for i := range s {
		s[i] = first + step*i
	}

	return s
}

func buildQuestion(progression []int, hiddenIndex int) string {
	if hiddenIndex < 0 || hiddenIndex >= len(progression) {
		panic("hiddenIndex out of range")
	}

	result := make([]string, len(progression))

	for i, v := range progression {
		if i == hiddenIndex {
			result[i] = ".."
		} else {
			result[i] = strconv.Itoa(v)
		}
	}

	return strings.Join(result, " ")
}
