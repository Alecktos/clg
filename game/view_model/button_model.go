package view_model

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/dev_utils"
	"github.com/Alecktos/clg/game/input"
)

type ButtonModel interface {
	Update(rectangle common.Rectangle)
	IsPressed() bool
	IsClicked() bool
	FirstPressedTick() bool
}

type buttonModel struct {
	isPressed        bool
	pressedPosition  common.Position
	hasPressed       bool
	firstPressedTick bool
}

func NewButtonModel() ButtonModel {
	return &buttonModel{}
}

func (b *buttonModel) Update(rectangle common.Rectangle) {
	if b.isPressed == true && input.IsPressed() == false && rectangle.Contains(b.pressedPosition) { //has released
		b.isPressed = false
		b.hasPressed = true
		dev_utils.Log("Button clicked registered")
	} else {
		b.hasPressed = false
	}

	if b.isPressed {
		b.pressedPosition = *input.Position()
	}

	isPressed := input.IsPressed() && rectangle.Contains(*input.Position())
	if !b.isPressed && isPressed {
		b.firstPressedTick = true
	} else {
		b.firstPressedTick = false
	}

	b.isPressed = isPressed
}

func (b *buttonModel) IsPressed() bool {
	return b.isPressed
}

func (b *buttonModel) IsClicked() bool {
	return b.hasPressed
}

func (b *buttonModel) FirstPressedTick() bool {
	return b.firstPressedTick
}
