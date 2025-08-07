package view_models

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/dev_utils"
	"github.com/Alecktos/clg/game/input"
	"github.com/Alecktos/clg/game/providers"
)

type ButtonModel interface {
	providers.VisibilityProvider
	providers.DisableProvider
	Update(rectangle common.Rectangle)
	IsPressed() bool
	IsClicked() bool
	FirstPressedTick() bool
	Reset()
}

type buttonModel struct {
	providers.VisibilityProvider
	providers.DisableProvider
	isPressed        bool
	pressedPosition  common.Position
	hasPressed       bool
	firstPressedTick bool
}

func NewButtonModel() ButtonModel {
	return &buttonModel{
		VisibilityProvider: providers.NewVisibilityProvider(),
		DisableProvider:    providers.NewDisableProvider(),
	}
}

func (b *buttonModel) Update(rectangle common.Rectangle) {
	if !b.IsVisible() || b.IsDisabled() {
		b.isPressed = false
		b.hasPressed = false
		b.firstPressedTick = false
		return
	}

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

func (b *buttonModel) Reset() {
	b.isPressed = false
	b.hasPressed = false
	b.firstPressedTick = false
	b.pressedPosition = common.Position{}
	b.VisibilityProvider.Reset()
	b.DisableProvider.Reset()
}
