package game

const ROUNDS_COUNT = 3

type Round struct {
	Question string
	Answer   string
}

type Game struct {
	Description string
	Rounds      []Round
}
