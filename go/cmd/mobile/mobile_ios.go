// +build ios

package mobile

import "github.com/hajimehoshi/ebiten/mobile"

// UpdateTouchesOnIOS dispatches touch events on iOS.
func UpdateTouchesOnIOS(phase int, ptr int64, x, y int) {
	// Prepare this function if you also want to make your game run on iOS.
	mobile.UpdateTouchesOnIOS(phase, ptr, x, y)
}
