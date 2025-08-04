package error_scene

import (
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ErrorScene struct {
	errorMessage *ErrorMessage
	err          error
}

func NewErrorScene() *ErrorScene {
	errorMessage, _ := NewErrorMessage()
	return &ErrorScene{
		errorMessage: errorMessage,
	}
}

func (e *ErrorScene) SetError(err error) {
	e.err = err
}

func (e *ErrorScene) Draw(screen *ebiten.Image) {
	screen.Fill(config.MidnightBlue())

	errorMessage := ""
	if e.err != nil {
		errorMessage = e.err.Error()
	}

	e.errorMessage.Draw(screen)
	if config.DevMode {
		ebitenutil.DebugPrint(screen, errorMessage)
	}
}
