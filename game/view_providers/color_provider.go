package view_providers

import (
	"github.com/Alecktos/clg/game/config"
	"image/color"
)

type ColorProvider struct {
	initialColor   config.ClgColor
	color          config.ClgColor
	onColorChanged func(config.ClgColor)
}

func NewColorProvider(c config.ClgColor, onColorChanged func(config.ClgColor)) ColorProvider {
	if onColorChanged == nil {
		onColorChanged = func(clgColor config.ClgColor) {}
	}
	return ColorProvider{
		initialColor:   c,
		color:          c,
		onColorChanged: onColorChanged,
	}
}

func (c *ColorProvider) ResetToInitialColor() {
	c.SetColor(c.initialColor)
}

func (c *ColorProvider) DarkenColor(factor float32) {
	r, g, b, a := c.color.RGBA()
	darkenColor := color.RGBA{
		R: uint8(float32(r>>8) * factor),
		G: uint8(float32(g>>8) * factor),
		B: uint8(float32(b>>8) * factor),
		A: uint8(a >> 8), // Keep alpha unchanged
	}
	c.SetColor(config.ClgColor{Color: darkenColor})
}

func (c *ColorProvider) GetColor() config.ClgColor {
	return c.color
}

func (c *ColorProvider) SetColor(newColor config.ClgColor) {
	if newColor != c.color {
		c.color = newColor
		c.onColorChanged(newColor)
	}

}
