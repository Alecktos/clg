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
	text          *view.Text
	textLoadError error
}

func NewErrorMessage() (*ErrorMessage, error) {
	text, loadError := view.NewCenterAlignedText(errorMessageText, 10)

	return &ErrorMessage{
		text:          text,
		textLoadError: loadError,
	}, nil
}

func (e *ErrorMessage) Draw(screen *ebiten.Image) {
	if fonts.GameFont == nil || e.textLoadError != nil {
		ebitenutil.DebugPrint(screen, errorMessageText)
		return
	}

	e.text.Draw(screen)
}
