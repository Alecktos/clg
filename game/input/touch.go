package input

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"slices"
)

var (
	justPressedTouchIds []ebiten.TouchID
	firstPressedTouchId ebiten.TouchID
)

type touch struct{}

func (t *touch) update() {
	justPressedTouchIds = inpututil.AppendJustPressedTouchIDs(justPressedTouchIds)
	justReleasedTouchIds := []ebiten.TouchID{}
	justReleasedTouchIds = inpututil.AppendJustReleasedTouchIDs(justReleasedTouchIds)

	for _, releasedTouchId := range justReleasedTouchIds {
		justPressedTouchIds = slices.DeleteFunc(justPressedTouchIds, func(c ebiten.TouchID) bool { return releasedTouchId == c })
	}

	if len(justPressedTouchIds) > 0 {
		firstPressedTouchId = justPressedTouchIds[0]
	}
}

func (t *touch) hasJustReleased() bool {
	return inpututil.IsTouchJustReleased(firstPressedTouchId)
}

func (t *touch) isPressed() bool {
	return len(justPressedTouchIds) > 0
}

func (t *touch) position() *common.Position {
	x, y := ebiten.TouchPosition(firstPressedTouchId)
	return &common.Position{
		X: float64(x),
		Y: float64(y),
	}
}
