package common

import (
	"github.com/Alecktos/clg/game/input"
)

type Button struct {
	Rectangle
	isPressed  bool
	hasPressed bool
}

func (b *Button) Update() {
	if b.isPressed == true && input.IsPressed() == false && b.Contains(*input.Position()) { //has released
		b.hasPressed = true
	} else {
		b.hasPressed = false
	}
	b.isPressed = input.IsPressed() && b.Contains(*input.Position())
}

func (b *Button) IsPressed() bool {
	return b.isPressed
}

func (b *Button) IsClicked() bool {
	return b.hasPressed
}
