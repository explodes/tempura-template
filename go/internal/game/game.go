package game

import (
	"math/rand"
	"time"

	"errors"

	"github.com/explodes/tempura-template/go/internal/res"
	"github.com/explodes/tempura"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"io"
)

var (
	ScreenBounds = tempura.R(0, 0, ScreenWidth, ScreenHeight)

	regularTermination = errors.New("goodbye!")
)

type Game struct {
	time         float64
	loader       tempura.Loader
	stopwatch    tempura.Stopwatch
	scene        Scene
	input        Input
	audioContext *audio.Context
	fullscreen   bool

	redScore  int
	blueScore int
}

type Scene interface {
	Update(dt float64) error
	Draw(image *ebiten.Image)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame() (*Game, error) {
	if debug {
		defer tempura.LogStart("Game init").End()
	}
	loader := tempura.NewCachedLoader(tempura.NewLoaderDebug(res.Asset, debug))
	audioContext, err := audio.NewContext(audioSampleRate)
	if err != nil {
		return nil, err
	}

	game := &Game{
		loader:       loader,
		stopwatch:    tempura.NewStopwatch(),
		input:        NewInput(),
		audioContext: audioContext,
	}

	if err := game.SetNewScene(NewTitleScene); err != nil {
		return nil, err
	}

	return game, nil
}

func (g *Game) SetNewScene(factory func(*Game) (scene Scene, err error)) error {
	if debug {
		defer tempura.LogStart("Set New Scene").End()
	}
	scene, err := factory(g)
	if err != nil {
		return err
	}
	return g.SetScene(scene)
}

func (g *Game) SetScene(scene Scene) error {
	DebugLog("new scene: %T", scene)

	if closeable, shouldClose := g.scene.(io.Closer); shouldClose {
		if err := closeable.Close(); err != nil {
			return err
		}
	}

	g.scene = scene
	return nil
}

func (g *Game) Update(image *ebiten.Image) error {
	dt := g.stopwatch.TimeDelta()
	g.time += dt

	if g.input.Exit() {
		return regularTermination
	}
	if g.input.ToggleFullscreen() {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if g.scene != nil {
		if err := g.scene.Update(dt); err != nil {
			return err
		}
		if !ebiten.IsRunningSlowly() {
			g.scene.Draw(image)
		}
	}

	return nil
}

func (g *Game) Pause() {
	g.stopwatch.Pause()
}

func (g *Game) Resume() {
	g.stopwatch.Resume()
}
