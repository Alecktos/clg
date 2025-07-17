package view

import (
	"errors"
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view_providers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"strings"
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

type Text interface {
	SetText(text string, layout TextLayout)
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

func (t *clgText) SetText(content string, layout TextLayout) {
	t.f = &text.GoTextFace{
		Source: fonts.GameFont,
		Size:   layout.FontSize,
	}

	lineSpacing := layout.FontSize + 10

	t.op = &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   layout.HorizontalAlign,
			SecondaryAlign: layout.VerticalAlign,
			LineSpacing:    lineSpacing,
		},
	}
	t.op.LineSpacing = lineSpacing

	t.op.GeoM.Translate(layout.Position.X, layout.Position.Y)

	textColor := layout.Color
	t.op.ColorScale.ScaleWithColor(textColor)
	t.ColorProvider = view_providers.NewColorProvider(textColor, func(clgColor config.ClgColor) {
		t.op.ColorScale.Reset()
		t.op.ColorScale.ScaleWithColor(clgColor)
	})

	lines := t.splitTextIntoLines(content, layout.MaxWidth)
	t.text = &lines
}

func (t *clgText) Draw(screen *ebiten.Image) {
	if t.text == nil || t.f == nil || t.op == nil {
		return
	}

	text.Draw(screen, *t.text, t.f, t.op)
}

func (t *clgText) splitTextIntoLines(input string, maxWidth int) string {
	var lines []string
	var currentLine string
	var currentWidth int

	for _, word := range t.splitWords(input) {
		wordWidthFloat, _ := text.Measure(word, t.f, t.op.LineSpacing) //retunerar height också vilket är bra.
		wordWidth := int(wordWidthFloat)
		if currentWidth+wordWidth > maxWidth {
			lines = append(lines, currentLine)
			currentLine = word
			currentWidth = wordWidth
		} else {
			if currentLine != "" {
				currentLine += " "
				widthFloat, _ := text.Measure(word, t.f, t.op.LineSpacing)
				currentWidth = currentWidth + int(widthFloat)
			}
			currentLine += word
			currentWidth += wordWidth
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

func (t *clgText) splitWords(input string) []string {
	return strings.Fields(input)
}
