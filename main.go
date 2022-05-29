package main

import (
	"fmt"
	"snake/game"
)

func main() {
	err := game.Run(game.Config{
		Framerate: 2,
	})

	if err != nil {
		fmt.Printf("The game crashed with error: %v\n", err)
	}
}
