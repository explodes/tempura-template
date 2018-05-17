package core

import (
	"errors"

	"github.com/explodes/tempura"
)

const (
	Title        = "Tanks"
	ScreenWidth  = 768
	ScreenHeight = 432

	AudioSampleRate = 44100
)

var (
	ScreenBounds = tempura.R(0, 0, ScreenWidth, ScreenHeight)

	RegularTermination = errors.New("goodbye")
	GameTermination    = errors.New("game over")
)
