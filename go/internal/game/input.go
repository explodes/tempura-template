package game

type Input interface {
	Exit() bool
}

type inputImpl struct{}

func NewInput() Input {
	return &inputImpl{}
}
