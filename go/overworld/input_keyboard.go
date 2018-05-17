// +build darwin freebsd linux windows js
// +build !android
// +build !ios

package overworld

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func ToggleFullscreen() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyF)
}

func ToggleMute() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyM)
}
