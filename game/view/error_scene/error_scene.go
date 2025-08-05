package error_scene

import (
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/dev_utils"
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

	errorDescription := ""
	if e.err != nil {
		errorDescription = e.err.Error()
	}

	e.errorMessage.Draw(screen)
	if config.DevMode {
		ebitenutil.DebugPrint(screen, errorDescription)
		dev_utils.Log(errorDescription)
	}
}
