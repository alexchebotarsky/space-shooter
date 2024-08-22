package main

import (
	"log"

	"github.com/goodleby/space-shooter/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.New()
	if err != nil {
		log.Fatalf("Error creating new game: %v", err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Space shooter")

	err = ebiten.RunGame(g)
	if err != nil {
		log.Fatalf("Error running the game: %v", err)
	}
}
