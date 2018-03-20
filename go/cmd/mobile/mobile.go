package mobile

import (
	"github.com/explodes/tempura-template/go/internal/game"
	"github.com/hajimehoshi/ebiten/mobile"
)

const (
	ScreenWidth  = game.ScreenWidth
	ScreenHeight = game.ScreenHeight
)

var (
	running bool
	g       *game.Game
)

// IsRunning returns a boolean value indicating whether the game is running.
func IsRunning() bool {
	return running
}

// Start starts the game.
func Start(scale float64) error {
	running = true
	var err error
	g, err = game.NewGame()
	if err != nil {
		return err
	}
	if err := mobile.Start(g.Update, game.ScreenWidth, game.ScreenHeight, scale, game.Title); err != nil {
		return err
	}
	return nil
}

// Update proceeds the game.
func Update() error {
	return mobile.Update()
}

func Pause() {
	if g != nil {
		g.Pause()
	}
}

func Resume() {
	if g != nil {
		g.Resume()
	}
}

// UpdateTouchesOnAndroid dispatches touch events on Android.
func UpdateTouchesOnAndroid(action int, id int, x, y int) {
	mobile.UpdateTouchesOnAndroid(action, id, x, y)
}

// UpdateTouchesOnIOS dispatches touch events on iOS.
func UpdateTouchesOnIOS(phase int, ptr int64, x, y int) {
	// Prepare this function if you also want to make your game run on iOS.
	mobile.UpdateTouchesOnIOS(phase, ptr, x, y)
}
