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
}

func NewText(text string) *Text {
	return &Text{
		text: text,
	}
}

func (t *Text) Draw(screen *ebiten.Image) error {
	if fonts.GameFont == nil {
		return errors.New("font not loaded correctly")
	}

	op := text.DrawOptions{}
	op.GeoM.Translate(20, 20)
	op.ColorScale.ScaleWithColor(config.ChampagneGold())
	op.LineSpacing = 25
	f := text.GoTextFace{
		Source: fonts.GameFont,
		Size:   20,
	}

	text.Draw(screen, t.text, &f, &op)
	return nil
}