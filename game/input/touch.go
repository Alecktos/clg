package input

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	justPressedTouchIds  []ebiten.TouchID
	justReleasedTouchIds []ebiten.TouchID
	firstPressedTouchId  ebiten.TouchID
)

type touch struct{}

func (t *touch) update() {
	//TODO: Behöver jag rensa innan. printa ut resultatet här
	justPressedTouchIds = inpututil.AppendJustPressedTouchIDs(justPressedTouchIds)
	justReleasedTouchIds = inpututil.AppendJustReleasedTouchIDs(justReleasedTouchIds)
	firstPressedTouchId = justPressedTouchIds[0]
}

func (t *touch) hasJustReleased() bool {
	return inpututil.IsTouchJustReleased(firstPressedTouchId)
}

func (t *touch) isPressed() bool {
	return len(justPressedTouchIds) > 0
}

func (t *touch) position() common.Position {
	x, y := ebiten.TouchPosition(firstPressedTouchId)
	return common.Position{
		X: x,
		Y: y,
	}
}

func init() {
	// justPressedTouchIds = []
	// justReleasedTouchIds = []
}

/*
var (
	allTouchIDs          []ebiten.TouchID
	currentTouchIDs      map[ebiten.TouchID]bool
	justPressedTouchIDs  map[ebiten.TouchID]bool
	justReleasedTouchIDs map[ebiten.TouchID]bool
)

func UpdateTouchIDs() {
	newPressedTouchIDs := []ebiten.TouchID{}
	newPressedTouchIDs = inpututil.AppendJustPressedTouchIDs(newPressedTouchIDs)
	justPressedTouchIDs = map[ebiten.TouchID]bool{}
	for i := 0; i < len(newPressedTouchIDs); i++ {
		justPressedTouchIDs[newPressedTouchIDs[i]] = true
		currentTouchIDs[newPressedTouchIDs[i]] = true
	}
	justReleasedTouchIDs = map[ebiten.TouchID]bool{}
	allTouchIDs = maps.Keys(currentTouchIDs)
	newReleasedTouchIDs := []ebiten.TouchID{}
	newReleasedTouchIDs = inpututil.AppendJustReleasedTouchIDs(newReleasedTouchIDs)
	for i := 0; i < len(newReleasedTouchIDs); i++ {
		justReleasedTouchIDs[newReleasedTouchIDs[i]] = true
		delete(currentTouchIDs, newReleasedTouchIDs[i])
	}
}

func GetTouchIDs() []ebiten.TouchID {
	return allTouchIDs
}

func IsTouchJustPressed(touchID ebiten.TouchID) bool {
	return justPressedTouchIDs[touchID]
}

func IsTouchJustReleased(touchID ebiten.TouchID) bool {
	return justReleasedTouchIDs[touchID]
}

func TouchPosition(id ebiten.TouchID) (x, y int) {
	return ebiten.TouchPosition(id)
}

func init() {
	allTouchIDs = []ebiten.TouchID{}
	currentTouchIDs = map[ebiten.TouchID]bool{}
	justPressedTouchIDs = map[ebiten.TouchID]bool{}
	justReleasedTouchIDs = map[ebiten.TouchID]bool{}
}
*/
