package view

import (
	"errors"
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view_providers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	text string
	op   *text.DrawOptions
	f    *text.GoTextFace
	view_providers.ColorProvider
}

func NewCenterAlignedText(alignedText string, y float64, fontSize float64) (*Text, error) {
	if fonts.GameFont == nil {
		return nil, errors.New("font not loaded correctly")
	}

	f := &text.GoTextFace{
		Source: fonts.GameFont,
		Size:   fontSize,
	}

	lineSpacing := fontSize + 5
	width, _ := text.Measure(alignedText, f, lineSpacing)

	op := text.DrawOptions{}
	op.LineSpacing = lineSpacing

	op.GeoM.Translate(config.WindowWidth/2-width/2, y)

	textColor := config.ChampagneGold()
	op.ColorScale.ScaleWithColor(textColor)
	colorModel := view_providers.NewColorProvider(textColor, func(clgColor config.ClgColor) {
		op.ColorScale.Reset()
		op.ColorScale.ScaleWithColor(clgColor)
	})

	return &Text{
		text:          alignedText,
		op:            &op,
		f:             f,
		ColorProvider: colorModel,
	}, nil
}

func (t *Text) Draw(screen *ebiten.Image) {
	text.Draw(screen, t.text, t.f, t.op)
}
