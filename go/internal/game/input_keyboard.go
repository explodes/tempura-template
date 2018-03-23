// +build darwin freebsd linux windows js
// +build !android
// +build !ios

package game

import (
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten"
)

func (i *inputImpl) ToggleFullscreen() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyF)
}
