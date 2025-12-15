package game

const roundsCount = 3

type Round struct {
	Question string
	Answer   string
}

type Game struct {
	Description string
	BuildRound  func() Round
	RoundsCount int
}

type BuildRound func() Round

func New(description string, buildRound BuildRound) Game {
	return Game{
		Description: description,
		BuildRound:  buildRound,
		RoundsCount: roundsCount,
	}
}
