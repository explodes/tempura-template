package title

import "github.com/explodes/tempura-template/go/games"

func init() {
	games.RegisterGameFactory("title", NewGame)
}
