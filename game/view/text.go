package view

import (
	"errors"
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	text string
	op   *text.DrawOptions
	f    *text.GoTextFace
}

func NewCenterAlignedText(alignedText string, y float64) (*Text, error) {
	if fonts.GameFont == nil {
		return nil, errors.New("font not loaded correctly")
	}

	f := &text.GoTextFace{
		Source: fonts.GameFont,
		Size:   20,
	}

	lineSpacing := 25.0
	width, _ := text.Measure(alignedText, f, lineSpacing)

	op := text.DrawOptions{}
	op.LineSpacing = lineSpacing

	op.GeoM.Translate(config.WindowWidth/2-width/2, y)
	op.ColorScale.ScaleWithColor(config.ChampagneGold())

	return &Text{
		text: alignedText,
		op:   &op,
		f:    f,
	}, nil
}

func (t *Text) Draw(screen *ebiten.Image) {
	text.Draw(screen, t.text, t.f, t.op)
}
