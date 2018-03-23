package game

type Input interface {
	Exit() bool
	ToggleFullscreen() bool
}

type inputImpl struct{}

func NewInput() Input {
	return &inputImpl{}
}
