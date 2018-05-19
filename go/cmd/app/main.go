package main

import (
	"log"

	_ "github.com/explodes/tempura-template/go/cmd/games_registry"
	"github.com/explodes/tempura-template/go/core"
	"github.com/explodes/tempura-template/go/overworld"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	game, err := overworld.NewOverworld("title")
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(game.Update, core.ScreenWidth, core.ScreenHeight, 1, core.Title); err != nil && err != core.RegularTermination {
		log.Fatal(err)
	}
}
