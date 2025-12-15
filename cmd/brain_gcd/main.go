package main

import "github.com/fey/go-brain-games/games/gcd"
import "github.com/fey/go-brain-games/internal/cli"

func main() {
	cli.Run(gcd.Create())
}
