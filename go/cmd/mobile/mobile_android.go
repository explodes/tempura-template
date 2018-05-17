package mobile

import "github.com/hajimehoshi/ebiten/mobile"

// UpdateTouchesOnAndroid dispatches touch events on Android.
func UpdateTouchesOnAndroid(action int, id int, x, y int) {
	mobile.UpdateTouchesOnAndroid(action, id, x, y)
}
