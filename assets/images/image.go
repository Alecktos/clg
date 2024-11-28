package images

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png" // All the sudden I needed this. Might not work in iOS??
)

func LoadBrickColorsImage() (*ebiten.Image, error) {
	return loadSingleImage(brick_colors_png)
}

func LoadBrickImage() (*ebiten.Image, error) {
	return loadSingleImage(brick_png)
}

func loadSingleImage(b []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
