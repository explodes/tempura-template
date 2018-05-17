package overworld

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/explodes/tempura"
	"github.com/explodes/tempura-template/go/core"
	"github.com/explodes/tempura-template/go/games"
	"github.com/explodes/tempura-template/go/res"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/pkg/errors"
)

var _ core.Context = (*Overworld)(nil)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Overworld struct {
	loader        tempura.Loader
	audioContext  *audio.Context
	muted         bool
	stopwatch     tempura.Stopwatch
	overworldGame string

	game core.Game
}

func NewOverworld(overworldGame string) (*Overworld, error) {
	loader := tempura.NewLoaderDebug(res.Asset, core.Debug)
	audioContext, err := audio.NewContext(core.AudioSampleRate)
	if err != nil {
		return nil, err
	}
	overworld := &Overworld{
		loader:        loader,
		audioContext:  audioContext,
		muted:         false,
		stopwatch:     tempura.NewStopwatch(),
		overworldGame: overworldGame,
	}
	if err := overworld.startOverworldGame(); err != nil {
		return nil, err
	}
	return overworld, nil
}

func (o *Overworld) Loader() tempura.Loader {
	return o.loader
}

func (o *Overworld) AudioContext() *audio.Context {
	return o.audioContext
}

func (o *Overworld) Muted() bool {
	return o.muted
}

func (o *Overworld) startOverworldGame() error {
	return o.LoadGame(o.overworldGame)
}

func (o *Overworld) LoadGame(name string) error {
	if core.Debug {
		defer tempura.LogStart(fmt.Sprintf("load game %s", name)).End()
	}
	factory := games.GetGameFactory(name)
	if factory == nil {
		return errors.Errorf("%s game does not exist. did you forget to register it?", name)
	}
	game, err := factory(o)
	if err != nil {
		return err
	}
	return o.setGame(game)
}
func (o *Overworld) setGame(game core.Game) error {
	core.DebugLog("new game: %T", game)
	o.game = game
	return nil
}

func (o *Overworld) Update(image *ebiten.Image) error {
	dt := o.stopwatch.TimeDelta()
	if Exit() {
		core.DebugLog("exit")
		return core.RegularTermination
	}
	if ToggleFullscreen() {
		if core.Debug {
			core.DebugLog("fullscreen: %v", !ebiten.IsFullscreen())
		}
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if ToggleMute() {
		o.muted = !o.muted
		core.DebugLog("mute: %v", o.muted)
		o.game.OnMuted(o.muted)
	}
	if o.game != nil {
		err := o.game.Update(dt)
		if err != nil {
			o.closeGame()
			if err == core.GameTermination {
				return o.startOverworldGame()
			} else if change, ok := err.(*core.ChangeGameError); ok {
				return o.LoadGame(change.Game)
			}
			return err
		}
		if !ebiten.IsRunningSlowly() {
			o.game.Draw(image)
		}
	}
	return nil
}

func (o *Overworld) closeGame() {
	if o.game == nil {
		return
	}
	if err := o.game.Close(); err != nil {
		core.Log("error closing game: %v", err)
	}
	o.game = nil
}

func (o *Overworld) Pause() {
	core.DebugLog("paused")
	o.stopwatch.Pause()
}

func (o *Overworld) Resume() {
	core.DebugLog("resumed")
	o.stopwatch.Resume()
}
