package input

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	mouseButtonIsPressed    bool
	justReleasedMouseButton bool
)

type mouse struct{}

func (m *mouse) update() {
	justReleasedMouseButton = inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
	if mouseButtonIsPressed == true {
		mouseButtonIsPressed = !justReleasedMouseButton
	} else {
		mouseButtonIsPressed = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	}
}

func (m *mouse) isPressed() bool {
	return mouseButtonIsPressed
}

func (m *mouse) hasJustReleased() bool {
	return justReleasedMouseButton
}

func (m *mouse) position() *common.Position {
	x, y := ebiten.CursorPosition()
	return &common.Position{
		X: x,
		Y: y,
	}
}

func init() {
	mouseButtonIsPressed = false
	justReleasedMouseButton = false
}
