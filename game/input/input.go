package input

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
)

type inputDevice interface {
	update()
	isPressed() bool
	hasJustReleased() bool
	position() common.Position
}

var (
	device inputDevice
)

func Update() {
	device.update()
}

func IsPressed() bool {
	return device.isPressed()
}

func hasJustReleased() bool {
	return device.hasJustReleased()
}

func Position() common.Position {
	return device.position()
}

// Använd interface för att få det att enklare. Kräver struct på dem båda
func init() {
	if config.CurrentDevice == config.DeviceTypeDesktop {
		device = &mouse{}
	} else if config.CurrentDevice == config.DeviceTypeMobile {
		device = &touch{}
	} else {
		panic("Unknown device type")
	}
}
