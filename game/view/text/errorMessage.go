package text

import (
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawErrorMessage(screen *ebiten.Image) {
	const errorMessage = `Failed to load game.
Please try and restart the app.`

	if fonts.GameFont == nil {
		ebitenutil.DebugPrint(screen, errorMessage)
		return
	}

	op := text.DrawOptions{}
	op.GeoM.Translate(20, 20)
	op.ColorScale.ScaleWithColor(config.ChampagneGold())
	op.LineSpacing = 25
	f := text.GoTextFace{
		Source: fonts.GameFont,
		Size:   20,
	}

	text.Draw(screen, errorMessage, &f, &op)
}
