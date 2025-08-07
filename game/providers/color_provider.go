package providers

import (
	"image/color"

	"github.com/Alecktos/clg/game/config"
)

type ColorProvider interface {
	ResetToInitialColor()
	DarkenColor(factor float32)
	GetColor() config.ClgColor
	SetColor(newColor config.ClgColor)
}

type colorProvider struct {
	initialColor   config.ClgColor
	color          config.ClgColor
	onColorChanged func(config.ClgColor)
}

func NewColorProvider(c config.ClgColor, onColorChanged func(config.ClgColor)) ColorProvider {
	if onColorChanged == nil {
		onColorChanged = func(clgColor config.ClgColor) {}
	}
	return &colorProvider{
		initialColor:   c,
		color:          c,
		onColorChanged: onColorChanged,
	}
}

func (c *colorProvider) ResetToInitialColor() {
	c.SetColor(c.initialColor)
}

func (c *colorProvider) DarkenColor(factor float32) {
	r, g, b, a := c.color.RGBA()
	darkenColor := color.RGBA{
		R: uint8(float32(r>>8) * factor),
		G: uint8(float32(g>>8) * factor),
		B: uint8(float32(b>>8) * factor),
		A: uint8(a >> 8), // Keep alpha unchanged
	}
	c.SetColor(config.ClgColor{Color: darkenColor})
}

func (c *colorProvider) GetColor() config.ClgColor {
	return c.color
}

func (c *colorProvider) SetColor(newColor config.ClgColor) {
	if newColor != c.color {
		c.color = newColor
		c.onColorChanged(newColor)
	}

}
