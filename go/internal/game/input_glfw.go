// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (i *inputImpl) Exit() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}
