package input

import (
	"github.com/Alecktos/clg/misc"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	justPressedMouseButton  bool
	justReleasedMouseButton bool
)

type mouse struct{}

func (m *mouse) update() {
	justPressedMouseButton = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	justReleasedMouseButton = inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

func (m *mouse) hasJustPressed() bool {
	return justPressedMouseButton
}

func (m *mouse) hasJustReleased() bool {
	return justReleasedMouseButton
}

func (m *mouse) position() misc.Position {
	x, y := ebiten.CursorPosition()
	return misc.Position{
		X: x,
		Y: y,
	}
}

func init() {
	justPressedMouseButton = false
	justReleasedMouseButton = false
}
