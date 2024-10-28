package view

import (
	"github.com/Alecktos/clg/assets/images"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type ClgSprite struct {
	Img *ebiten.Image
}

func (c *ClgSprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(c.Img, op)
}

func (c *ClgSprite) loadImage() {
	var err error
	image, err := images.LoadBrickImage()
	c.Img = image
	if err != nil {
		log.Fatal(err)
	}
}

func NewClgSprite() *ClgSprite {
	clgImage := &ClgSprite{}
	clgImage.loadImage()
	return clgImage
}
