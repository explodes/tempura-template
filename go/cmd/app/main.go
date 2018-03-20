package main

import (
	"log"

	"github.com/explodes/tempura-template/go/internal/game"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(g.Update, game.ScreenWidth, game.ScreenHeight, 1, game.Title); err != nil {
		log.Fatal(err)
	}
}
