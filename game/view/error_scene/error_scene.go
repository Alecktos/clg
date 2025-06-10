package error_scene

import (
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ErrorScene struct {
	errorMessage *ErrorMessage
}

func NewErrorScene() *ErrorScene {
	return &ErrorScene{
		errorMessage: NewErrorMessage(),
	}
}

func (e *ErrorScene) Draw(screen *ebiten.Image, error string) {
	screen.Fill(config.MidnightBlue())
	e.errorMessage.Draw(screen)
	if config.DEV_MODE {
		ebitenutil.DebugPrint(screen, error)
	}
}