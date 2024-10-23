package view

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ClgSprite struct {
	Img *ebiten.Image
}

/*
	func (c *ClgSprite) loadImage() {
		var err error
		var image *ebiten.Image
		image, _, err = ebitenutil.NewImageFromFile("brick.png")
		c.Img = image
		if err != nil {
			log.Fatal(err)
		}
	}
*/
func NewClgSprite() *ClgSprite {
	clgImage := &ClgSprite{}
	//clgImage.loadImage()
	return clgImage
}
