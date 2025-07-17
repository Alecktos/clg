package error_scene

import (
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const errorMessageText = `Failed to load game.
Please try and restart the app.`

type ErrorMessage struct {
	errorText     view.Text
	textLoadError error
}

func NewErrorMessage() (*ErrorMessage, error) {
	errorText, loadError := view.NewCenterAlignedText()

	textLayout := view.NewTextLayout(common.Position{X: config.WindowWidth / 2, Y: config.WindowHeight / 2})
	textLayout.VerticalAlign = text.AlignCenter
	errorText.SetText(errorMessageText, textLayout)

	return &ErrorMessage{
		errorText:     errorText,
		textLoadError: loadError,
	}, nil
}

func (e *ErrorMessage) Draw(screen *ebiten.Image) {
	if fonts.GameFont == nil || e.textLoadError != nil {
		ebitenutil.DebugPrint(screen, errorMessageText)
		return
	}

	e.errorText.Draw(screen)
}
