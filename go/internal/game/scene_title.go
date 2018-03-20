package game

import (
	"github.com/hajimehoshi/ebiten"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var _ Scene = (*titleScene)(nil)

type titleScene struct {
	g *Game
}

func NewTitleScene(g *Game) (Scene, error) {
	s := &titleScene{
		g: g,
	}
	return s, nil
}

func (s *titleScene) Update(dt float64) error {
	return nil
}

func (s *titleScene) Draw(image *ebiten.Image) {
	ebitenutil.DebugPrint(image, "Hello, world")
}
