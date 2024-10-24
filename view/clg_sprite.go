package view

import (
	"github.com/Alecktos/clg/assets/images"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type ClgSprite struct {
	Img *ebiten.Image
}

//TODO: implement draw function and make image private

func (c *ClgSprite) loadImage() {
	var err error
	//var image *ebiten.Image
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
