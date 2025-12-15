package main

import "github.com/fey/go-brain-games/games/even"
import "github.com/fey/go-brain-games/internal/cli"

func main() {
	cli.Run(even.Create())
}
