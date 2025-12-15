package main

import "github.com/fey/go-brain-games/games/progression"
import "github.com/fey/go-brain-games/internal/cli"

func main() {
	cli.Run(progression.Create())
}
