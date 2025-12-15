package game

type Round struct {
	Question string
	Answer   string
}

type Game struct {
	Description string
	BuildRound BuildRound
}

type BuildRound func() Round
