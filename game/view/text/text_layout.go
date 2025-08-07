package text

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextLayout struct {
	FontSize        float64
	Position        common.Position
	HorizontalAlign text.Align
	VerticalAlign   text.Align
	MaxWidth        int
	Color           config.ClgColor
}

func NewTextLayout(position common.Position) TextLayout {
	return TextLayout{
		FontSize:        config.StandardFontSize,
		Position:        position,
		HorizontalAlign: text.AlignCenter,
		VerticalAlign:   text.AlignStart,
		MaxWidth:        config.WindowWidth,
		Color:           config.ChampagneGold(),
	}
}
