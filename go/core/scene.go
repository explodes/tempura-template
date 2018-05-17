package core

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type Scene interface {
	Update(dt float64) error
	Draw(image *ebiten.Image)
}

type ChangeGameError struct {
	Game string
}

func (c *ChangeGameError) Error() string {
	return fmt.Sprintf("change game: %s", c.Game)
}
