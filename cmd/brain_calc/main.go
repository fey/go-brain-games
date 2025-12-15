package main

import (
	"github.com/fey/go-brain-games/games/calc"
	"github.com/fey/go-brain-games/internal/cli"
)

func main() {
	cli.Run(calc.Create())
}
