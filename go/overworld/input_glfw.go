// +build darwin freebsd linux windows
// +build !js
// +build !android
// +build !ios

package overworld

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func Exit() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}
