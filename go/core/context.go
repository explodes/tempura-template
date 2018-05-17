package core

import (
	"github.com/explodes/tempura"
	"github.com/hajimehoshi/ebiten/audio"
)

type Context interface {
	Loader() tempura.Loader
	AudioContext() *audio.Context
	Muted() bool
}
