package main

import "github.com/fey/go-brain-games/games/prime"
import "github.com/fey/go-brain-games/internal/cli"

func main() {
	cli.Run(prime.Create())
}
