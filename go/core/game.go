package core

import (
	"github.com/hajimehoshi/ebiten"
)

type Game interface {
	Scene
	OnMuted(muted bool)
	Close() error
}

var _ Game = (*GameSceneLoop)(nil)

type GameSceneLoop struct {
	scene Scene
}

func (g *GameSceneLoop) SetScene(scene Scene) error {
	DebugLog("new scene: %T", scene)
	g.scene = scene
	return nil
}

func (g *GameSceneLoop) Update(dt float64) error {
	if g.scene != nil {
		if err := g.scene.Update(dt); err != nil {
			return err
		}
	}
	return nil
}

func (g *GameSceneLoop) Draw(image *ebiten.Image) {
	if g.scene != nil {
		g.scene.Draw(image)
	}
}

func (g *GameSceneLoop) OnMuted(muted bool) {}
func (g *GameSceneLoop) Close() error       { return nil }
