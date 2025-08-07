package text

import (
	"errors"
	"strings"

	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/providers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text interface {
	SetText(text string, layout TextLayout)
	Draw(screen *ebiten.Image)
	providers.ColorProvider
	MeasureHeight() float64
}

type clgText struct {
	text *string
	op   *text.DrawOptions
	f    *text.GoTextFace
	providers.ColorProvider
}

func NewText() (Text, error) {
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
	t.ColorProvider = providers.NewColorProvider(textColor, func(clgColor config.ClgColor) {
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

func (t *clgText) MeasureHeight() float64 {
	if t.f == nil || t.op == nil {
		return 0
	}
	_, textHeight := text.Measure(*t.text, t.f, t.op.LineSpacing)
	return textHeight
}
