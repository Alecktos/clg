package view

import (
	"errors"
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view_providers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text interface {
	SetText(alignedText string, y float64, fontSize float64)
	Draw(screen *ebiten.Image)
	view_providers.ColorProviderInterface
}

type clgText struct {
	text *string
	op   *text.DrawOptions
	f    *text.GoTextFace
	view_providers.ColorProvider
}

func NewCenterAlignedText() (Text, error) {
	if fonts.GameFont == nil {
		return nil, errors.New("font not loaded correctly")
	}
	return &clgText{}, nil
}

func (t *clgText) SetText(alignedText string, y float64, fontSize float64) {
	t.f = &text.GoTextFace{
		Source: fonts.GameFont,
		Size:   fontSize,
	}

	lineSpacing := fontSize + 5
	width, _ := text.Measure(alignedText, t.f, lineSpacing)

	t.op = &text.DrawOptions{}
	t.op.LineSpacing = lineSpacing

	t.op.GeoM.Translate(config.WindowWidth/2-width/2, y)

	textColor := config.ChampagneGold()
	t.op.ColorScale.ScaleWithColor(textColor)
	t.ColorProvider = view_providers.NewColorProvider(textColor, func(clgColor config.ClgColor) {
		t.op.ColorScale.Reset()
		t.op.ColorScale.ScaleWithColor(clgColor)
	})
	t.text = &alignedText
}

func (t *clgText) Draw(screen *ebiten.Image) {
	if t.text == nil || t.f == nil || t.op == nil {
		return
	}
	text.Draw(screen, *t.text, t.f, t.op)
}
