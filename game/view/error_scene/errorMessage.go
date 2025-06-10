package error_scene

import (
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const errorMessageText = `Failed to load game.
Please try and restart the app.`

type ErrorMessage struct {
	text *view.Text
}

func NewErrorMessage() *ErrorMessage {
	return &ErrorMessage{
		text: view.NewText(errorMessageText),
	}
}

func (errorMessage *ErrorMessage) Draw(screen *ebiten.Image) {
	if fonts.GameFont == nil {
		ebitenutil.DebugPrint(screen, errorMessageText)
		return
	}

	errorMessage.text.Draw(screen)
}