package title

import (
	"image"

	"github.com/explodes/tempura"
	"github.com/explodes/tempura-template/go/core"
	"github.com/explodes/tempura/tinge"
	"github.com/explodes/tempura/tux"
	"github.com/hajimehoshi/ebiten"
)

var _ core.Scene = (*titleScene)(nil)

var (
	buttonColors = []tinge.Transform{
		tinge.Colorize(tinge.Color64(0xffff0000)), // TouchNone
		tinge.Colorize(tinge.Color64(0xff00ff00)), // TouchDown
		tinge.Colorize(tinge.Color64(0xff0000ff)), // TouchDrag
		tinge.Colorize(tinge.Color64(0xffff00ff)), // TouchUp
		tinge.Colorize(tinge.Color64(0xffffffff)), // Touch miss
	}
)

type titleScene struct {
	g       *Game
	touches *tux.TouchInput
	menu    *tempura.Objects
	err     error
}

type menuItem struct {
	name    string
	image   image.Image
	onClick func() error
}

func NewTitleScene(game *Game) (core.Scene, error) {
	loader := game.context.Loader()

	battleImage, err := loader.Image("images/battle.png")
	if err != nil {
		return nil, err
	}
	tankImage, err := loader.Image("images/tank.png")
	if err != nil {
		return nil, err
	}
	settingsImage, err := loader.Image("images/settings.png")
	if err != nil {
		return nil, err
	}
	exitImage, err := loader.Image("images/exit.png")
	if err != nil {
		return nil, err
	}

	touches := tux.NewTouchInput()

	titleScene := &titleScene{
		g:       game,
		touches: touches,
	}
	titleScene.menu = titleScene.fabricateMenu(touches, []menuItem{
		{"battle", battleImage, onBattleClick},
		{"tank", tankImage, onTankClick},
		{"settings", settingsImage, onSettingsClick},
		{"exit", exitImage, onClickExit},
	})
	return titleScene, nil
}

func (s *titleScene) fabricateMenu(touches *tux.TouchInput, menuItems []menuItem) *tempura.Objects {
	const (
		menuItemSize    = 48
		menuItemPadding = 24
	)

	menu := tempura.NewObjects()

	x := core.ScreenWidth*0.5 - ((float64(len(menuItems))*(menuItemSize+menuItemPadding))-menuItemPadding)*0.5
	y := core.ScreenHeight*0.5 - menuItemSize*0.5
	dx := float64(menuItemSize + menuItemPadding)

	for _, menuItem := range menuItems {
		menuItem := menuItem

		// build image states
		var images []*tempura.ImageDrawable
		for _, colorTransform := range buttonColors {
			images = append(images, colorizeImage(colorTransform(menuItem.image)))
		}

		// make our menu item
		menu.Add(&tempura.Object{
			Tag:      menuItem.name,
			Drawable: images[0],
			Pos:      tempura.V(x, y),
			Size:     tempura.V(menuItemSize, menuItemSize),
			Steps: tempura.MakeBehaviors(
				func(source *tempura.Object, dt float64) {
					// touch index 0: first pointer finger or left mouse button
					touch := touches.GetTouch(0)

					// pick our image
					switch source.HitTest(touch.Position) {
					case true:
						source.Drawable = images[int(touch.Event)]
					case false:
						source.Drawable = images[len(images)-1]
						// no-hit, do nothing
						return
					}
					if touch.Event == tux.TouchUp {
						if err := menuItem.onClick(); err != nil {
							s.err = err
						}
					}
				},
			),
		})
		x += dx
	}

	return menu
}

func colorizeImage(img image.Image, err error) *tempura.ImageDrawable {
	if err != nil {
		panic(err)
	}
	timg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	return tempura.NewImageDrawable(timg)
}

func (s *titleScene) Update(dt float64) error {
	s.touches.Update(nil)
	s.menu.Update(dt)
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s *titleScene) Draw(image *ebiten.Image) {
	s.menu.Draw(nil, image)
}

func onClickExit() error {
	return core.RegularTermination
}

func onBattleClick() error {
	core.Log("battle not implemented")
	return nil
}
func onTankClick() error {
	core.Log("tanks not implemented")
	return nil
}
func onSettingsClick() error {
	// TODO(explodes): return &core.ChangeGameError{Game: "settings"}
	// This requires that the "settings" Game is implemented and registered.
	core.Log("settings not implemented")
	return nil
}
