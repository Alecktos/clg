package input

import (
	"github.com/Alecktos/clg/misc"
)

type inputDevice interface {
	update()
	hasJustPressed() bool
	hasJustReleased() bool
	position() misc.Position
}

var (
	device inputDevice
)

func Update() {
	device.update()
}

func HasJustPressed() bool {
	return device.hasJustPressed()
}

// Använd interface för att få det att enklare. Kräver struct på dem båda
func init() {
	if misc.CurrentDevice == misc.DeviceTypeDesktop {
		device = &mouse{}
	} else if misc.CurrentDevice == misc.DeviceTypeMobile {
		device = &touch{}
	} else {
		panic("Unknown device type")
	}
}
