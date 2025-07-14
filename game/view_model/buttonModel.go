package view_model

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/input"
)

type ButtonModel interface {
	Update(rectangle common.Rectangle)
	IsPressed() bool
	IsClicked() bool
}

type buttonModel struct {
	isPressed  bool
	hasPressed bool
}

func NewButtonModel() ButtonModel {
	return &buttonModel{}
}

func (b *buttonModel) Update(rectangle common.Rectangle) {
	if b.isPressed == true && input.IsPressed() == false && rectangle.Contains(*input.Position()) { //has released
		b.hasPressed = true
	} else {
		b.hasPressed = false
	}
	b.isPressed = input.IsPressed() && rectangle.Contains(*input.Position())
}

func (b *buttonModel) IsPressed() bool {
	return b.isPressed
}

func (b *buttonModel) IsClicked() bool {
	return b.hasPressed
}
