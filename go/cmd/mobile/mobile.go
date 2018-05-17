package mobile

import (
	_ "github.com/explodes/tempura-template/go/cmd/games_registry"
	"github.com/explodes/tempura-template/go/core"
	"github.com/explodes/tempura-template/go/overworld"
	"github.com/hajimehoshi/ebiten/mobile"
)

var (
	running bool
	game    *overworld.Overworld
)

const (
	ScreenWidth  = core.ScreenWidth
	ScreenHeight = core.ScreenHeight
)

// IsRunning returns a boolean value indicating whether the game is running.
func IsRunning() bool {
	return running
}

// Start starts the game.
func Start(scale float64) error {
	running = true
	var err error
	game, err = overworld.NewOverworld("title")
	if err != nil {
		return err
	}
	if err := mobile.Start(game.Update, ScreenWidth, ScreenHeight, scale, core.Title); err != nil {
		return err
	}
	return nil
}

// Update proceeds the game.
func Update() error {
	return mobile.Update()
}

func Pause() {
	if game != nil {
		game.Pause()
	}
}

func Resume() {
	if game != nil {
		game.Resume()
	}
}
