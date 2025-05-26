package view

import (
	"github.com/Alecktos/clg/assets/images"
	"github.com/Alecktos/clg/game/common"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameBrick struct {
	Img *ebiten.Image
	common.Rectangle
}

func (c *GameBrick) Update() {

}

func (c *GameBrick) Draw(screen *ebiten.Image) {
	//if input.IsPressed() && c.Contains(input.Position()) { does not work on phone
	op := &ebiten.DrawImageOptions{}
	imgSourceSize := c.Img.Bounds().Size()
	scaleX := float64(c.Width) / float64(imgSourceSize.X)
	scaleY := float64(c.Height) / float64(imgSourceSize.Y)
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(c.Position.X), float64(c.Position.Y))

	screen.DrawImage(c.Img, op)
	//}
}

func (c *GameBrick) loadImage() error {
	var err error
	image, err := images.LoadBrickImage()
	c.Img = image
	return err
}

func NewGameBrick() (*GameBrick, error) {
	clgImage := &GameBrick{
		Rectangle: common.Rectangle{
			Position: common.Position{X: 0, Y: 0},
			Width:    75,
			Height:   75,
		},
	}
	err := clgImage.loadImage()
	return clgImage, err
}
